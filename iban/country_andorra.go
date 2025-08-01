// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateAndorraIBAN validates Andorra IBAN
func validateAndorraIBAN(iban string) error {
	if len(iban) != 24 {
		return fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "AD" {
		return fmt.Errorf("static value rule, pos: 0, expected value: AD, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:12]; !ascii.IsDigit(subject) {
		return fmt.Errorf("range rule, start pos: 2, length: 10, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[12:24]; !ascii.IsAlphaNumeric(subject) {
		return fmt.Errorf("range rule, start pos: 12, length: 12, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateAndorraIBAN generates Andorra IBAN
func generateAndorraIBAN() (string, error) {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("AD")
	ascii.Digits(sb, 10)
	ascii.AlphaNumeric(sb, 12)

	return ReplaceChecksum(sb.String())
}

// getAndorraBBAN retrieves BBAN structure from Andorra IBAN
func getAndorraBBAN(iban string) (BBAN, error) {
	if len(iban) != 24 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:24],
		BankCode:         iban[4:8],
		BranchCode:       iban[8:12],
		NationalChecksum: "",
		AccountNumber:    iban[12:24],
	}, nil
}
