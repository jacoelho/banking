package iban_test

import (
	"fmt"
	"math/rand"

	"github.com/jacoelho/banking/iban"
)

func ExampleValidate() {
	// Validate an iban
	fmt.Println(iban.Validate("VG96VPVG0000012345678901") == nil)
	// Output: true
}

func ExampleValidateUnitedKingdomIBAN() {
	// Validate an iban using country specific validate function
	err := iban.ValidateUnitedKingdomIBAN("GB29NWBK60161331926819")

	fmt.Println(err == nil)
	// Output: true
}

func ExampleGenerate() {
	// force results - not needed during regular usage
	iban.SeedGenerator(rand.New(rand.NewSource(1)))

	// Validate an iban using country specific validate function
	result, _ := iban.Generate("GB")

	fmt.Println(result)
	// Output: GB21LBZG50604129841576
}

func ExampleGetBBAN() {
	// Get BBAN from IBAN
	result, _ := iban.GetBBAN("GB29NWBK60161331926819")

	fmt.Println(result.BBAN, result.BankCode, result.BranchCode, result.AccountNumber)
	// Output: NWBK60161331926819 NWBK 601613 31926819
}

func ExamplePaperFormat() {
	// Pretty print an iban
	fmt.Println(iban.PaperFormat("GB29NWBK60161331926819"))
	// Output: GB29 NWBK 6016 1331 9268 19
}
