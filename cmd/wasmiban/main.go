//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"syscall/js"

	"github.com/jacoelho/banking/iban"
	"github.com/jacoelho/banking/iso3166"
)

type Response struct {
	Error string          `json:"error,omitempty"`
	Data  *ValidationData `json:"data"`
}

type ValidationData struct {
	CountryCode      string `json:"countryCode,omitempty"`
	FormattedIBAN    string `json:"formattedIban,omitempty"`
	CorrectedIBAN    string `json:"correctedIban,omitempty"`
	BankCode         string `json:"bankCode,omitempty"`
	BranchCode       string `json:"branchCode,omitempty"`
	AccountNumber    string `json:"accountNumber,omitempty"`
	NationalChecksum string `json:"nationalChecksum,omitempty"`
}

var (
	lengthRegex       = regexp.MustCompile(`want: (\d+)`)
	countryCodeRegex  = regexp.MustCompile(`expected value: ([A-Z]{2})`)
	notSupportedRegex = regexp.MustCompile(`^([A-Z0-9]+) is not supported`)
	rangeRuleRegex    = regexp.MustCompile(`range rule, start pos: (\d+), length: (\d+), expected type (\w+)`)
)

func humanizeError(err error) string {
	if err == nil {
		return ""
	}

	errMsg := err.Error()

	switch {
	case strings.Contains(errMsg, "incorrect checksum"):
		return "Invalid checksum"

	case strings.Contains(errMsg, "unexpected length"):
		if matches := lengthRegex.FindStringSubmatch(errMsg); len(matches) > 1 {
			return fmt.Sprintf("Invalid length (expected %s characters)", matches[1])
		}
		return "Invalid IBAN length"

	case strings.Contains(errMsg, "range rule"):
		if matches := rangeRuleRegex.FindStringSubmatch(errMsg); len(matches) > 3 {
			startPos := matches[1]
			length := matches[2]
			expectedType := matches[3]

			start, _ := strconv.Atoi(startPos)
			len, _ := strconv.Atoi(length)
			endPos := start + len - 1

			var typeDesc string
			switch expectedType {
			case "Digit":
				typeDesc = "only digits"
			case "UpperCaseLetters":
				typeDesc = "only letters"
			case "AlphaNumeric":
				typeDesc = "letters and digits"
			default:
				typeDesc = expectedType
			}

			return fmt.Sprintf("Invalid characters between position %s and %d - %s allowed", startPos, endPos, typeDesc)
		}
		return "Invalid format"

	case strings.Contains(errMsg, "static value rule"):
		if strings.Contains(errMsg, "pos: 0") {
			if matches := countryCodeRegex.FindStringSubmatch(errMsg); len(matches) > 1 {
				return fmt.Sprintf("Invalid country code (expected %s)", matches[1])
			}
		}
		return "Invalid format - expected specific value"

	case strings.Contains(errMsg, "is not supported"):
		if matches := notSupportedRegex.FindStringSubmatch(errMsg); len(matches) > 1 {
			return fmt.Sprintf("Country code '%s' is not supported", matches[1])
		}
		return "Country code not supported"

	default:
		return strings.TrimSuffix(strings.TrimSpace(errMsg), ": validation error")
	}
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

		if strings.Contains(err.Error(), "incorrect checksum") {
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
		data.NationalChecksum = bban.NationalChecksum
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
