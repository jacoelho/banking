// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateCostaRicaIBAN validates Costa Rica IBAN
func validateCostaRicaIBAN(iban string) error {
	if len(iban) != 22 {
		return fmt.Errorf("unexpected length, want: 22: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "CR" {
		return fmt.Errorf("static value rule, pos: 0, expected value: CR, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:22]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 20, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateCostaRicaIBAN generates Costa Rica IBAN
func generateCostaRicaIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("CR")
	generator.Digits(sb, 20)

	return replaceChecksum(sb.String())
}

// getCostaRicaBBAN retrieves BBAN structure from Costa Rica IBAN
func getCostaRicaBBAN(iban string) (BBAN, error) {
	if len(iban) != 22 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 22: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:22],
		BankCode:         iban[4:8],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[8:22],
	}, nil
}
