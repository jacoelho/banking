// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateFalklandIslandsIBAN validates Falkland Islands IBAN
func validateFalklandIslandsIBAN(iban string) error {
	if len(iban) != 18 {
		return fmt.Errorf("unexpected length, want: 18: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "FK" {
		return fmt.Errorf("static value rule, pos: 0, expected value: FK, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.IsDigit(subject) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:6]; !ascii.IsUpperCase(subject) {
		return fmt.Errorf("range rule, start pos: 4, length: 2, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[6:18]; !ascii.IsDigit(subject) {
		return fmt.Errorf("range rule, start pos: 6, length: 12, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateFalklandIslandsIBAN generates Falkland Islands IBAN
func generateFalklandIslandsIBAN() (string, error) {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("FK")
	ascii.Digits(sb, 2)
	ascii.UpperCaseLetters(sb, 2)
	ascii.Digits(sb, 12)

	return ReplaceChecksum(sb.String())
}

// getFalklandIslandsBBAN retrieves BBAN structure from Falkland Islands IBAN
func getFalklandIslandsBBAN(iban string) (BBAN, error) {
	if len(iban) != 18 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 18: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:18],
		BankCode:         iban[4:6],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[6:18],
	}, nil
}
