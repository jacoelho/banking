// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateSerbiaIBAN validates Serbia IBAN
func validateSerbiaIBAN(iban string) error {
	if len(iban) != 22 {
		return fmt.Errorf("unexpected length, want: 22: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "RS" {
		return fmt.Errorf("static value rule, pos: 0, expected value: RS, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:22]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 20, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateSerbiaIBAN generates Serbia IBAN
func generateSerbiaIBAN() (string, error) {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("RS")
	ascii.Digits(sb, 20)

	return ReplaceChecksum(sb.String())
}

// getSerbiaBBAN retrieves BBAN structure from Serbia IBAN
func getSerbiaBBAN(iban string) (BBAN, error) {
	if len(iban) != 22 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 22: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:22],
		BankCode:         iban[4:7],
		BranchCode:       "",
		NationalChecksum: iban[20:22],
		AccountNumber:    iban[7:20],
	}, nil
}
