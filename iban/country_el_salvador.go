// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateElSalvadorIBAN validates El Salvador IBAN
func validateElSalvadorIBAN(iban string) error {
	if len(iban) != 28 {
		return fmt.Errorf("unexpected length, want: 28: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "SV" {
		return fmt.Errorf("static value rule, pos: 0, expected value: SV, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:8]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
		return fmt.Errorf("range rule, start pos: 4, length: 4, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[8:28]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 8, length: 20, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateElSalvadorIBAN generates El Salvador IBAN
func generateElSalvadorIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("SV")
	generator.Digits(sb, 2)
	generator.UpperCaseLetters(sb, 4)
	generator.Digits(sb, 20)

	return replaceChecksum(sb.String())
}

// getElSalvadorBBAN retrieves BBAN structure from El Salvador IBAN
func getElSalvadorBBAN(iban string) (BBAN, error) {
	if len(iban) != 28 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 28: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:28],
		BankCode:         iban[4:8],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[8:28],
	}, nil
}
