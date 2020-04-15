// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateSpainIBAN validates Spain IBAN
func validateSpainIBAN(iban string) error {
	if len(iban) != 24 {
		return fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "ES" {
		return fmt.Errorf("static value rule, pos: 0, expected value: ES, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:24]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 22, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateSpainIBAN generates Spain IBAN
func generateSpainIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("ES")
	generator.Digits(sb, 22)

	return ReplaceChecksum(sb.String())
}

// getSpainBBAN retrieves BBAN structure from Spain IBAN
func getSpainBBAN(iban string) (BBAN, error) {
	if len(iban) != 24 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:24],
		BankCode:         iban[4:8],
		BranchCode:       iban[8:12],
		NationalChecksum: iban[12:14],
		AccountNumber:    iban[14:24],
	}, nil
}
