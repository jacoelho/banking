// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateMoldovaIBAN validates Moldova IBAN
func validateMoldovaIBAN(iban string) error {
	if len(iban) != 24 {
		return fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "MD" {
		return fmt.Errorf("static value rule, pos: 0, expected value: MD, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.IsDigit(subject) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:24]; !ascii.IsAlphaNumeric(subject) {
		return fmt.Errorf("range rule, start pos: 4, length: 20, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateMoldovaIBAN generates Moldova IBAN
func generateMoldovaIBAN() (string, error) {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("MD")
	ascii.Digits(sb, 2)
	ascii.AlphaNumeric(sb, 20)

	return ReplaceChecksum(sb.String())
}

// getMoldovaBBAN retrieves BBAN structure from Moldova IBAN
func getMoldovaBBAN(iban string) (BBAN, error) {
	if len(iban) != 24 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:24],
		BankCode:         iban[4:6],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[6:24],
	}, nil
}
