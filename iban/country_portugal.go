// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validatePortugalIBAN validates Portugal IBAN
func validatePortugalIBAN(iban string) error {
	if len(iban) != 25 {
		return fmt.Errorf("unexpected length, want: 25: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "PT" {
		return fmt.Errorf("static value rule, pos: 0, expected value: PT, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:25]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 23, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generatePortugalIBAN generates Portugal IBAN
func generatePortugalIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("PT")
	ascii.Digits(sb, 23)

	return replaceChecksum(sb.String())
}

// getPortugalBBAN retrieves BBAN structure from Portugal IBAN
func getPortugalBBAN(iban string) (BBAN, error) {
	if len(iban) != 25 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 25: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:25],
		BankCode:         iban[4:8],
		BranchCode:       iban[8:12],
		NationalChecksum: iban[23:25],
		AccountNumber:    iban[12:23],
	}, nil
}
