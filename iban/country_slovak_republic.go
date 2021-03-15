// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateSlovakRepublicIBAN validates Slovak Republic IBAN
func validateSlovakRepublicIBAN(iban string) error {
	if len(iban) != 24 {
		return fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "SK" {
		return fmt.Errorf("static value rule, pos: 0, expected value: SK, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:24]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 22, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateSlovakRepublicIBAN generates Slovak Republic IBAN
func generateSlovakRepublicIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("SK")
	ascii.Digits(sb, 22)

	return replaceChecksum(sb.String())
}

// getSlovakRepublicBBAN retrieves BBAN structure from Slovak Republic IBAN
func getSlovakRepublicBBAN(iban string) (BBAN, error) {
	if len(iban) != 24 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:24],
		BankCode:         iban[4:8],
		BranchCode:       iban[8:14],
		NationalChecksum: iban[21:24],
		AccountNumber:    iban[14:21],
	}, nil
}
