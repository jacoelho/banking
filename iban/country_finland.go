// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateFinlandIBAN validates Finland IBAN
func validateFinlandIBAN(iban string) error {
	if len(iban) != 18 {
		return fmt.Errorf("unexpected length, want: 18: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "FI" {
		return fmt.Errorf("static value rule, pos: 0, expected value: FI, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:18]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 16, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateFinlandIBAN generates Finland IBAN
func generateFinlandIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("FI")
	ascii.Digits(sb, 16)

	return replaceChecksum(sb.String())
}

// getFinlandBBAN retrieves BBAN structure from Finland IBAN
func getFinlandBBAN(iban string) (BBAN, error) {
	if len(iban) != 18 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 18: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:18],
		BankCode:         iban[4:7],
		BranchCode:       "",
		NationalChecksum: iban[17:18],
		AccountNumber:    iban[7:17],
	}, nil
}
