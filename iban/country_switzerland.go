// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateSwitzerlandIBAN validates Switzerland IBAN
func validateSwitzerlandIBAN(iban string) error {
	if len(iban) != 21 {
		return fmt.Errorf("unexpected length, want: 21: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "CH" {
		return fmt.Errorf("static value rule, pos: 0, expected value: CH, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:9]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 7, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[9:21]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 9, length: 12, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateSwitzerlandIBAN generates Switzerland IBAN
func generateSwitzerlandIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("CH")
	generator.Digits(sb, 7)
	generator.AlphaNumeric(sb, 12)

	return ReplaceChecksum(sb.String())
}

// getSwitzerlandBBAN retrieves BBAN structure from Switzerland IBAN
func getSwitzerlandBBAN(iban string) (BBAN, error) {
	if len(iban) != 21 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 21: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:21],
		BankCode:         iban[4:9],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[9:21],
	}, nil
}
