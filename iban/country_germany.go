// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateGermanyIBAN validates Germany IBAN
func validateGermanyIBAN(iban string) error {
	if len(iban) != 22 {
		return fmt.Errorf("unexpected length, want: 22: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "DE" {
		return fmt.Errorf("static value rule, pos: 0, expected value: DE, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:22]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 20, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateGermanyIBAN generates Germany IBAN
func generateGermanyIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("DE")
	ascii.Digits(sb, 20)

	return replaceChecksum(sb.String())
}

// getGermanyBBAN retrieves BBAN structure from Germany IBAN
func getGermanyBBAN(iban string) (BBAN, error) {
	if len(iban) != 22 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 22: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:22],
		BankCode:         iban[4:12],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[12:22],
	}, nil
}
