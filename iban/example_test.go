package iban_test

import (
	"fmt"

	"github.com/jacoelho/banking/iban"
)

func ExampleValidate() {
	// Validate an iban
	fmt.Println(iban.Validate("VG96VPVG0000012345678901") == nil)
	// Output: true
}

func ExampleReplaceChecksum() {
	// Replace an iban checksum
	result, err := iban.ReplaceChecksum("GB99NWBK60161331926819")
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	// Output: GB29NWBK60161331926819
}

func ExampleGenerate() {
	// Validate an iban using country specific validate function
	result, err := iban.Generate("GB")
	if err != nil {
		panic(err)
	}

	fmt.Println(len(result), result[:2])
	// Output: 22 GB
}

func ExampleGenerateWithBBAN() {
	result, err := iban.GenerateWithBBAN("GB", iban.BBANParts{
		BankCode:      "NWBK",
		BranchCode:    "601613",
		AccountNumber: "31926819",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	// Output: GB29NWBK60161331926819
}

func ExampleGetBBAN() {
	// Get BBAN from IBAN
	result, err := iban.GetBBAN("GB29NWBK60161331926819")
	if err != nil {
		panic(err)
	}

	fmt.Println(result.BBAN, result.BankCode, result.BranchCode, result.AccountNumber)
	// Output: NWBK60161331926819 NWBK 601613 31926819
}

func ExamplePaperFormat() {
	// Pretty print an iban
	fmt.Println(iban.PaperFormat("GB29NWBK60161331926819"))
	// Output: GB29 NWBK 6016 1331 9268 19
}

func ExampleIsSEPA() {
	// IsSEPA returns if an iban country is a SEPA member
	result, err := iban.IsSEPA("GB29NWBK60161331926819")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	// Output: true
}

func ExampleIsSEPACountryCode() {
	// IsSEPACountry returns if a country is a SEPA member
	result, err := iban.IsSEPACountryCode("AE")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	// Output: false
}
