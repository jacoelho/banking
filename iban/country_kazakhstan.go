// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateKazakhstanIBAN validates Kazakhstan IBAN
func validateKazakhstanIBAN(iban string) error {
	if len(iban) != 20 {
		return fmt.Errorf("unexpected length, want: 20: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "KZ" {
		return fmt.Errorf("static value rule, pos: 0, expected value: KZ, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:7]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 5, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[7:20]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 7, length: 13, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateKazakhstanIBAN generates Kazakhstan IBAN
func generateKazakhstanIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("KZ")
	generator.Digits(sb, 5)
	generator.AlphaNumeric(sb, 13)

	return replaceChecksum(sb.String())
}

// getKazakhstanBBAN retrieves BBAN structure from Kazakhstan IBAN
func getKazakhstanBBAN(iban string) (BBAN, error) {
	if len(iban) != 20 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 20: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:20],
		BankCode:         iban[4:7],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[7:20],
	}, nil
}
