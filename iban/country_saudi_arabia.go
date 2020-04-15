// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateSaudiArabiaIBAN validates Saudi Arabia IBAN
func validateSaudiArabiaIBAN(iban string) error {
	if len(iban) != 24 {
		return fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "SA" {
		return fmt.Errorf("static value rule, pos: 0, expected value: SA, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:6]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 4, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[6:24]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 6, length: 18, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateSaudiArabiaIBAN generates Saudi Arabia IBAN
func generateSaudiArabiaIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("SA")
	generator.Digits(sb, 4)
	generator.AlphaNumeric(sb, 18)

	return replaceChecksum(sb.String())
}

// getSaudiArabiaBBAN retrieves BBAN structure from Saudi Arabia IBAN
func getSaudiArabiaBBAN(iban string) (BBAN, error) {
	if len(iban) != 24 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:24],
		BankCode:         iban[4:6],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[6:24],
	}, nil
}
