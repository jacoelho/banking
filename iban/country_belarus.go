// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateBelarusIBAN validates Belarus IBAN
func validateBelarusIBAN(iban string) error {
	if len(iban) != 28 {
		return fmt.Errorf("unexpected length, want: 28: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "BY" {
		return fmt.Errorf("static value rule, pos: 0, expected value: BY, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:8]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 4, length: 4, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[8:12]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 8, length: 4, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[12:28]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 12, length: 16, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateBelarusIBAN generates Belarus IBAN
func generateBelarusIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("BY")
	ascii.Digits(sb, 2)
	ascii.AlphaNumeric(sb, 4)
	ascii.Digits(sb, 4)
	ascii.AlphaNumeric(sb, 16)

	return replaceChecksum(sb.String())
}

// getBelarusBBAN retrieves BBAN structure from Belarus IBAN
func getBelarusBBAN(iban string) (BBAN, error) {
	if len(iban) != 28 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 28: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:28],
		BankCode:         iban[4:8],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[8:28],
	}, nil
}
