// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// ValidateUnitedArabEmiratesIBAN validates United Arab Emirates IBAN
func ValidateUnitedArabEmiratesIBAN(iban string) error {
	if len(iban) != 23 {
		return fmt.Errorf("unexpected length, want: 23: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "AE" {
		return fmt.Errorf("static value rule, pos: 0, expected value: AE, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:23]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 21, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateUnitedArabEmiratesIBAN generates United Arab Emirates IBAN
func GenerateUnitedArabEmiratesIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("AE")
	generator.Digits(sb, 21)

	return ReplaceChecksum(sb.String())
}

// GetUnitedArabEmiratesBBAN retrieves BBAN structure from United Arab Emirates IBAN
func GetUnitedArabEmiratesBBAN(iban string) (BBAN, error) {
	if len(iban) != 23 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 23: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:23],
		BankCode:         iban[4:7],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[7:23],
	}, nil
}
