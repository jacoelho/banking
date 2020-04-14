// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// ValidateTurkeyIBAN validates Turkey IBAN
func ValidateTurkeyIBAN(iban string) error {
	if len(iban) != 26 {
		return fmt.Errorf("unexpected length, want: 26: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "TR" {
		return fmt.Errorf("static value rule, pos: 0, expected value: TR, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:10]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 8, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[10:26]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 10, length: 16, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateTurkeyIBAN generates Turkey IBAN
func GenerateTurkeyIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("TR")
	generator.Digits(sb, 8)
	generator.AlphaNumeric(sb, 16)

	return ReplaceChecksum(sb.String())
}

// GetTurkeyBBAN retrieves BBAN structure from Turkey IBAN
func GetTurkeyBBAN(iban string) (BBAN, error) {
	if len(iban) != 26 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 26: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:26],
		BankCode:         iban[4:9],
		BranchCode:       "",
		NationalChecksum: iban[9:10],
		AccountNumber:    iban[10:26],
	}, nil
}
