// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// ValidateMauritaniaIBAN validates Mauritania IBAN
func ValidateMauritaniaIBAN(iban string) error {
	if len(iban) != 27 {
		return fmt.Errorf("unexpected length, want: 27: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "MR" {
		return fmt.Errorf("static value rule, pos: 0, expected value: MR, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:27]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 25, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateMauritaniaIBAN generates Mauritania IBAN
func GenerateMauritaniaIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("MR")
	generator.Digits(sb, 25)

	return ReplaceChecksum(sb.String())
}

// GetMauritaniaBBAN retrieves BBAN structure from Mauritania IBAN
func GetMauritaniaBBAN(iban string) (BBAN, error) {
	if len(iban) != 27 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 27: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:27],
		BankCode:         iban[4:9],
		BranchCode:       iban[9:14],
		NationalChecksum: iban[25:27],
		AccountNumber:    iban[14:25],
	}, nil
}
