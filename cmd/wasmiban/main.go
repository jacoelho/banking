//go:build js && wasm

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"syscall/js"

	"github.com/jacoelho/banking/iban"
	"github.com/jacoelho/banking/iso3166"
)

type Response struct {
	Error string          `json:"error,omitempty"`
	Data  *ValidationData `json:"data"`
}

type ValidationData struct {
	CountryCode   string `json:"countryCode,omitempty"`
	FormattedIBAN string `json:"formattedIban,omitempty"`
	CorrectedIBAN string `json:"correctedIban,omitempty"`
	BankCode      string `json:"bankCode,omitempty"`
	BranchCode    string `json:"branchCode,omitempty"`
	AccountNumber string `json:"accountNumber,omitempty"`
}

func humanizeError(err error) string {
	if err == nil {
		return ""
	}

	var lengthErr *iban.ErrValidationLength
	if errors.As(err, &lengthErr) {
		return fmt.Sprintf("Invalid length (expected %d characters)", lengthErr.Expected)
	}

	var checksumErr *iban.ErrValidationChecksum
	if errors.As(err, &checksumErr) {
		return "Invalid checksum"
	}

	var rangeErr *iban.ErrValidationRange
	if errors.As(err, &rangeErr) {
		endPos := rangeErr.Position + rangeErr.Length - 1

		var typeDesc string
		switch rangeErr.Expected {
		case iban.CharacterTypeDigit:
			typeDesc = "only digits"
		case iban.CharacterTypeUpperCase:
			typeDesc = "only letters"
		case iban.CharacterTypeAlphaNumeric:
			typeDesc = "letters and digits"
		default:
			typeDesc = rangeErr.Expected.String()
		}

		return fmt.Sprintf("Invalid characters between position %d and %d - %s allowed",
			rangeErr.Position, endPos, typeDesc)
	}

	var staticErr *iban.ErrValidationStaticValue
	if errors.As(err, &staticErr) {
		if staticErr.Position == 0 {
			return fmt.Sprintf("Invalid country code (expected %s)", staticErr.Expected)
		}
		return "Invalid format - expected specific value"
	}

	var unsupportedErr *iban.ErrUnsupportedCountry
	if errors.As(err, &unsupportedErr) {
		return fmt.Sprintf("Country code '%s' is not supported", unsupportedErr.CountryCode)
	}

	return err.Error()
}

func marshalResponse(error string, data *ValidationData) string {
	resp := Response{
		Error: error,
		Data:  data,
	}

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		return `{"error":"internal serialization error","data":null}`
	}

	return string(jsonBytes)
}

func validateIBAN(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return marshalResponse("Invalid number of arguments", nil)
	}

	ibanStr := args[0].String()
	if ibanStr == "" {
		return marshalResponse("IBAN cannot be empty", nil)
	}

	data := &ValidationData{}

	if len(ibanStr) >= 2 && iso3166.IsCountryCode(ibanStr[0:2]) {
		data.CountryCode = ibanStr[0:2]
	}

	err := iban.Validate(ibanStr)
	if err != nil {
		data.FormattedIBAN = iban.PaperFormat(ibanStr)

		var checksumErr *iban.ErrValidationChecksum
		if errors.As(err, &checksumErr) {
			if corrected, corrErr := iban.ReplaceChecksum(ibanStr); corrErr == nil {
				data.CorrectedIBAN = iban.PaperFormat(corrected)
			}
		}

		return marshalResponse(humanizeError(err), data)
	}

	data.FormattedIBAN = iban.PaperFormat(ibanStr)

	if bban, bbanErr := iban.GetBBAN(ibanStr); bbanErr == nil {
		data.BankCode = bban.BankCode
		data.BranchCode = bban.BranchCode
		data.AccountNumber = bban.AccountNumber
	}

	return marshalResponse("", data)
}

func paperFormat(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return ""
	}
	return iban.PaperFormat(args[0].String())
}

func replaceChecksum(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return map[string]any{
			"success": false,
			"error":   "invalid number of arguments",
		}
	}

	corrected, err := iban.ReplaceChecksum(args[0].String())
	if err != nil {
		return map[string]any{
			"success": false,
			"error":   err.Error(),
		}
	}

	return map[string]any{
		"success":   true,
		"corrected": corrected,
		"formatted": iban.PaperFormat(corrected),
	}
}

func main() {
	c := make(chan struct{})

	js.Global().Set("validateIBAN", js.FuncOf(validateIBAN))
	js.Global().Set("paperFormat", js.FuncOf(paperFormat))
	js.Global().Set("replaceChecksum", js.FuncOf(replaceChecksum))

	js.Global().Call("postMessage", "WASM Ready")

	<-c
}
