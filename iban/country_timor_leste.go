// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateTimorLesteIBAN validates Timor Leste IBAN
func validateTimorLesteIBAN(iban string) error {
	if len(iban) != 23 {
		return fmt.Errorf("unexpected length, want: 23: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "TL" {
		return fmt.Errorf("static value rule, pos: 0, expected value: TL, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:23]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 21, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateTimorLesteIBAN generates Timor Leste IBAN
func generateTimorLesteIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("TL")
	ascii.Digits(sb, 21)

	return replaceChecksum(sb.String())
}

// getTimorLesteBBAN retrieves BBAN structure from Timor Leste IBAN
func getTimorLesteBBAN(iban string) (BBAN, error) {
	if len(iban) != 23 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 23: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:23],
		BankCode:         iban[4:8],
		BranchCode:       iban[8:11],
		NationalChecksum: iban[21:23],
		AccountNumber:    iban[11:21],
	}, nil
}
