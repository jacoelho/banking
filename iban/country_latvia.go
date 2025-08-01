// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateLatviaIBAN validates Latvia IBAN
func validateLatviaIBAN(iban string) error {
	if len(iban) != 21 {
		return fmt.Errorf("unexpected length, want: 21: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "LV" {
		return fmt.Errorf("static value rule, pos: 0, expected value: LV, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.IsDigit(subject) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:8]; !ascii.IsUpperCase(subject) {
		return fmt.Errorf("range rule, start pos: 4, length: 4, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[8:21]; !ascii.IsAlphaNumeric(subject) {
		return fmt.Errorf("range rule, start pos: 8, length: 13, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateLatviaIBAN generates Latvia IBAN
func generateLatviaIBAN() (string, error) {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("LV")
	ascii.Digits(sb, 2)
	ascii.UpperCaseLetters(sb, 4)
	ascii.AlphaNumeric(sb, 13)

	return ReplaceChecksum(sb.String())
}

// getLatviaBBAN retrieves BBAN structure from Latvia IBAN
func getLatviaBBAN(iban string) (BBAN, error) {
	if len(iban) != 21 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 21: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:21],
		BankCode:         iban[4:8],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[8:21],
	}, nil
}
