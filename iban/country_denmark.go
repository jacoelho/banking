// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// ValidateDenmarkIBAN validates Denmark IBAN
func ValidateDenmarkIBAN(iban string) error {
	if len(iban) != 18 {
		return fmt.Errorf("unexpected length, want: 18: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "DK" {
		return fmt.Errorf("static value rule, pos: 0, expected value: DK, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:18]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 16, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateDenmarkIBAN generates Denmark IBAN
func GenerateDenmarkIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("DK")
	generator.Digits(sb, 16)

	return ReplaceChecksum(sb.String())
}

// GetDenmarkBBAN retrieves BBAN structure from Denmark IBAN
func GetDenmarkBBAN(iban string) (BBAN, error) {
	if len(iban) != 18 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 18: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:],
		BankCode:         iban[4:8],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[8:18],
	}, nil
}
