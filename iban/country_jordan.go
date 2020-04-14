// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// ValidateJordanIBAN validates Jordan IBAN
func ValidateJordanIBAN(iban string) error {
	if len(iban) != 30 {
		return fmt.Errorf("unexpected length, want: 30: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "JO" {
		return fmt.Errorf("static value rule, pos: 0, expected value: JO, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:8]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
		return fmt.Errorf("range rule, start pos: 4, length: 4, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[8:12]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 8, length: 4, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[12:30]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 12, length: 18, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateJordanIBAN generates Jordan IBAN
func GenerateJordanIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("JO")
	generator.Digits(sb, 2)
	generator.UpperCaseLetters(sb, 4)
	generator.Digits(sb, 4)
	generator.AlphaNumeric(sb, 18)

	return ReplaceChecksum(sb.String())
}

// GetJordanBBAN retrieves BBAN structure from Jordan IBAN
func GetJordanBBAN(iban string) (BBAN, error) {
	if len(iban) != 30 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 30: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:30],
		BankCode:         iban[4:8],
		BranchCode:       iban[8:12],
		NationalChecksum: "",
		AccountNumber:    iban[12:30],
	}, nil
}
