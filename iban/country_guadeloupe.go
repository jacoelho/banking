// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateGuadeloupeIBAN validates Guadeloupe IBAN
func validateGuadeloupeIBAN(iban string) error {
	if len(iban) != 27 {
		return fmt.Errorf("unexpected length, want: 27: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "GP" {
		return fmt.Errorf("static value rule, pos: 0, expected value: GP, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:14]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 12, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[14:25]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 14, length: 11, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[25:27]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 25, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateGuadeloupeIBAN generates Guadeloupe IBAN
func generateGuadeloupeIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("GP")
	generator.Digits(sb, 12)
	generator.AlphaNumeric(sb, 11)
	generator.Digits(sb, 2)

	return ReplaceChecksum(sb.String())
}

// getGuadeloupeBBAN retrieves BBAN structure from Guadeloupe IBAN
func getGuadeloupeBBAN(iban string) (BBAN, error) {
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
