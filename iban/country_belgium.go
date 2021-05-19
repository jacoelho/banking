// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateBelgiumIBAN validates Belgium IBAN
func validateBelgiumIBAN(iban string) error {
	if len(iban) != 16 {
		return fmt.Errorf("unexpected length, want: 16: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "BE" {
		return fmt.Errorf("static value rule, pos: 0, expected value: BE, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:16]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 14, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateBelgiumIBAN generates Belgium IBAN
func generateBelgiumIBAN() (string, error) {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("BE")
	ascii.Digits(sb, 14)

	return ReplaceChecksum(sb.String())
}

// getBelgiumBBAN retrieves BBAN structure from Belgium IBAN
func getBelgiumBBAN(iban string) (BBAN, error) {
	if len(iban) != 16 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 16: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:16],
		BankCode:         iban[4:7],
		BranchCode:       "",
		NationalChecksum: iban[14:16],
		AccountNumber:    iban[7:14],
	}, nil
}
