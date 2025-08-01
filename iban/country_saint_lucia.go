// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateSaintLuciaIBAN validates Saint Lucia IBAN
func validateSaintLuciaIBAN(iban string) error {
	if len(iban) != 32 {
		return fmt.Errorf("unexpected length, want: 32: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "LC" {
		return fmt.Errorf("static value rule, pos: 0, expected value: LC, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.IsDigit(subject) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:8]; !ascii.IsUpperCase(subject) {
		return fmt.Errorf("range rule, start pos: 4, length: 4, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[8:32]; !ascii.IsAlphaNumeric(subject) {
		return fmt.Errorf("range rule, start pos: 8, length: 24, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateSaintLuciaIBAN generates Saint Lucia IBAN
func generateSaintLuciaIBAN() (string, error) {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("LC")
	ascii.Digits(sb, 2)
	ascii.UpperCaseLetters(sb, 4)
	ascii.AlphaNumeric(sb, 24)

	return ReplaceChecksum(sb.String())
}

// getSaintLuciaBBAN retrieves BBAN structure from Saint Lucia IBAN
func getSaintLuciaBBAN(iban string) (BBAN, error) {
	if len(iban) != 32 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 32: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:32],
		BankCode:         iban[4:8],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[8:32],
	}, nil
}
