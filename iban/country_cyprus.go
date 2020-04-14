// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// ValidateCyprusIBAN validates Cyprus IBAN
func ValidateCyprusIBAN(iban string) error {
	if len(iban) != 28 {
		return fmt.Errorf("unexpected length, want: 28: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "CY" {
		return fmt.Errorf("static value rule, pos: 0, expected value: CY, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:12]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 10, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[12:28]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 12, length: 16, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateCyprusIBAN generates Cyprus IBAN
func GenerateCyprusIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("CY")
	generator.Digits(sb, 10)
	generator.AlphaNumeric(sb, 16)

	return ReplaceChecksum(sb.String())
}

// GetCyprusBBAN retrieves BBAN structure from Cyprus IBAN
func GetCyprusBBAN(iban string) (BBAN, error) {
	if len(iban) != 28 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 28: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:28],
		BankCode:         iban[4:7],
		BranchCode:       iban[7:12],
		NationalChecksum: "",
		AccountNumber:    iban[12:28],
	}, nil
}
