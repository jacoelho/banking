// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// ValidateKuwaitIBAN validates Kuwait IBAN
func ValidateKuwaitIBAN(iban string) error {
	if len(iban) != 30 {
		return fmt.Errorf("unexpected length, want: 30: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "KW" {
		return fmt.Errorf("static value rule, pos: 0, expected value: KW, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:8]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
		return fmt.Errorf("range rule, start pos: 4, length: 4, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[8:30]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 8, length: 22, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateKuwaitIBAN generates Kuwait IBAN
func GenerateKuwaitIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("KW")
	generator.Digits(sb, 2)
	generator.UpperCaseLetters(sb, 4)
	generator.AlphaNumeric(sb, 22)

	return ReplaceChecksum(sb.String())
}

// GetKuwaitBBAN retrieves BBAN structure from Kuwait IBAN
func GetKuwaitBBAN(iban string) (BBAN, error) {
	if len(iban) != 30 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 30: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:30],
		BankCode:         iban[4:8],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[8:30],
	}, nil
}
