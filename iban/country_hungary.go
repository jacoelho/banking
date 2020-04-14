// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// ValidateHungaryIBAN validates Hungary IBAN
func ValidateHungaryIBAN(iban string) error {
	if len(iban) != 28 {
		return fmt.Errorf("unexpected length, want: 28: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "HU" {
		return fmt.Errorf("static value rule, pos: 0, expected value: HU, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:28]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 26, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateHungaryIBAN generates Hungary IBAN
func GenerateHungaryIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("HU")
	generator.Digits(sb, 26)

	return ReplaceChecksum(sb.String())
}

// GetHungaryBBAN retrieves BBAN structure from Hungary IBAN
func GetHungaryBBAN(iban string) (BBAN, error) {
	if len(iban) != 28 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 28: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:28],
		BankCode:         iban[4:7],
		BranchCode:       iban[7:12],
		NationalChecksum: iban[27:28],
		AccountNumber:    iban[12:27],
	}, nil
}
