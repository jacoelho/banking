// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// ValidateGreeceIBAN validates Greece IBAN
func ValidateGreeceIBAN(iban string) error {
	if len(iban) != 27 {
		return fmt.Errorf("unexpected length, want: 27: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "GR" {
		return fmt.Errorf("static value rule, pos: 0, expected value: GR, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:11]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 9, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[11:27]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 11, length: 16, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateGreeceIBAN generates Greece IBAN
func GenerateGreeceIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("GR")
	generator.Digits(sb, 9)
	generator.AlphaNumeric(sb, 16)

	return ReplaceChecksum(sb.String())
}

// GetGreeceBBAN retrieves BBAN structure from Greece IBAN
func GetGreeceBBAN(iban string) (BBAN, error) {
	if len(iban) != 27 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 27: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:27],
		BankCode:         iban[4:7],
		BranchCode:       iban[7:11],
		NationalChecksum: "",
		AccountNumber:    iban[11:27],
	}, nil
}
