// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateSomaliaIBAN validates Somalia IBAN
func validateSomaliaIBAN(iban string) error {
	if len(iban) != 23 {
		return fmt.Errorf("unexpected length, want: 23: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "SO" {
		return fmt.Errorf("static value rule, pos: 0, expected value: SO, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:23]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 21, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateSomaliaIBAN generates Somalia IBAN
func generateSomaliaIBAN() (string, error) {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("SO")
	ascii.Digits(sb, 21)

	return ReplaceChecksum(sb.String())
}

// getSomaliaBBAN retrieves BBAN structure from Somalia IBAN
func getSomaliaBBAN(iban string) (BBAN, error) {
	if len(iban) != 23 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 23: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:23],
		BankCode:         iban[4:8],
		BranchCode:       iban[8:11],
		NationalChecksum: "",
		AccountNumber:    iban[11:23],
	}, nil
}
