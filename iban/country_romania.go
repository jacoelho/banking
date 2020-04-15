// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateRomaniaIBAN validates Romania IBAN
func validateRomaniaIBAN(iban string) error {
	if len(iban) != 24 {
		return fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "RO" {
		return fmt.Errorf("static value rule, pos: 0, expected value: RO, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:8]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
		return fmt.Errorf("range rule, start pos: 4, length: 4, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[8:24]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 8, length: 16, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateRomaniaIBAN generates Romania IBAN
func generateRomaniaIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("RO")
	generator.Digits(sb, 2)
	generator.UpperCaseLetters(sb, 4)
	generator.AlphaNumeric(sb, 16)

	return ReplaceChecksum(sb.String())
}

// getRomaniaBBAN retrieves BBAN structure from Romania IBAN
func getRomaniaBBAN(iban string) (BBAN, error) {
	if len(iban) != 24 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:24],
		BankCode:         iban[4:8],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[8:24],
	}, nil
}
