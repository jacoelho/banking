// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// ValidateItalyIBAN validates Italy IBAN
func ValidateItalyIBAN(iban string) error {
	if len(iban) != 27 {
		return fmt.Errorf("unexpected length, want: 27: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "IT" {
		return fmt.Errorf("static value rule, pos: 0, expected value: IT, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:5]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
		return fmt.Errorf("range rule, start pos: 4, length: 1, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[5:15]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 5, length: 10, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[15:27]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 15, length: 12, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateItalyIBAN generates Italy IBAN
func GenerateItalyIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("IT")
	generator.Digits(sb, 2)
	generator.UpperCaseLetters(sb, 1)
	generator.Digits(sb, 10)
	generator.AlphaNumeric(sb, 12)

	return ReplaceChecksum(sb.String())
}

// GetItalyBBAN retrieves BBAN structure from Italy IBAN
func GetItalyBBAN(iban string) (BBAN, error) {
	if len(iban) != 27 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 27: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:],
		BankCode:         iban[5:10],
		BranchCode:       iban[10:15],
		NationalChecksum: iban[4:5],
		AccountNumber:    iban[15:27],
	}, nil
}
