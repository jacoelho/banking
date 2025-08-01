// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateSeychellesIBAN validates Seychelles IBAN
func validateSeychellesIBAN(iban string) error {
	if len(iban) != 31 {
		return fmt.Errorf("unexpected length, want: 31: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "SC" {
		return fmt.Errorf("static value rule, pos: 0, expected value: SC, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.IsDigit(subject) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:8]; !ascii.IsUpperCase(subject) {
		return fmt.Errorf("range rule, start pos: 4, length: 4, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[8:28]; !ascii.IsDigit(subject) {
		return fmt.Errorf("range rule, start pos: 8, length: 20, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[28:31]; !ascii.IsUpperCase(subject) {
		return fmt.Errorf("range rule, start pos: 28, length: 3, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateSeychellesIBAN generates Seychelles IBAN
func generateSeychellesIBAN() (string, error) {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("SC")
	ascii.Digits(sb, 2)
	ascii.UpperCaseLetters(sb, 4)
	ascii.Digits(sb, 20)
	ascii.UpperCaseLetters(sb, 3)

	return ReplaceChecksum(sb.String())
}

// getSeychellesBBAN retrieves BBAN structure from Seychelles IBAN
func getSeychellesBBAN(iban string) (BBAN, error) {
	if len(iban) != 31 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 31: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:31],
		BankCode:         iban[4:8],
		BranchCode:       iban[8:12],
		NationalChecksum: "",
		AccountNumber:    iban[12:31],
	}, nil
}
