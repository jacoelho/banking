// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateQatarIBAN validates Qatar IBAN
func validateQatarIBAN(iban string) error {
	if len(iban) != 29 {
		return fmt.Errorf("unexpected length, want: 29: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "QA" {
		return fmt.Errorf("static value rule, pos: 0, expected value: QA, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:8]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
		return fmt.Errorf("range rule, start pos: 4, length: 4, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[8:29]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 8, length: 21, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateQatarIBAN generates Qatar IBAN
func generateQatarIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("QA")
	ascii.Digits(sb, 2)
	ascii.UpperCaseLetters(sb, 4)
	ascii.AlphaNumeric(sb, 21)

	return replaceChecksum(sb.String())
}

// getQatarBBAN retrieves BBAN structure from Qatar IBAN
func getQatarBBAN(iban string) (BBAN, error) {
	if len(iban) != 29 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 29: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:29],
		BankCode:         iban[4:8],
		BranchCode:       iban[8:12],
		NationalChecksum: "",
		AccountNumber:    iban[12:29],
	}, nil
}
