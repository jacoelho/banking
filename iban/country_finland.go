// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// ValidateFinlandIBAN validates Finland IBAN
func ValidateFinlandIBAN(iban string) error {
	if len(iban) != 18 {
		return fmt.Errorf("unexpected length, want: 18: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "FI" {
		return fmt.Errorf("static value rule, pos: 0, expected value: FI, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:18]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 16, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateFinlandIBAN generates Finland IBAN
func GenerateFinlandIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("FI")
	generator.Digits(sb, 16)

	return ReplaceChecksum(sb.String())
}

// GetFinlandBBAN retrieves BBAN structure from Finland IBAN
func GetFinlandBBAN(iban string) (BBAN, error) {
	if len(iban) != 18 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 18: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:],
		BankCode:         iban[4:7],
		BranchCode:       "",
		NationalChecksum: iban[17:18],
		AccountNumber:    iban[7:17],
	}, nil
}
