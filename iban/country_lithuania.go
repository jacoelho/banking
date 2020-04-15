// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateLithuaniaIBAN validates Lithuania IBAN
func validateLithuaniaIBAN(iban string) error {
	if len(iban) != 20 {
		return fmt.Errorf("unexpected length, want: 20: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "LT" {
		return fmt.Errorf("static value rule, pos: 0, expected value: LT, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:20]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 18, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateLithuaniaIBAN generates Lithuania IBAN
func generateLithuaniaIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("LT")
	generator.Digits(sb, 18)

	return ReplaceChecksum(sb.String())
}

// getLithuaniaBBAN retrieves BBAN structure from Lithuania IBAN
func getLithuaniaBBAN(iban string) (BBAN, error) {
	if len(iban) != 20 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 20: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:20],
		BankCode:         iban[4:9],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[9:20],
	}, nil
}
