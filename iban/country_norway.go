// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateNorwayIBAN validates Norway IBAN
func validateNorwayIBAN(iban string) error {
	if len(iban) != 15 {
		return fmt.Errorf("unexpected length, want: 15: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "NO" {
		return fmt.Errorf("static value rule, pos: 0, expected value: NO, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:15]; !ascii.IsDigit(subject) {
		return fmt.Errorf("range rule, start pos: 2, length: 13, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateNorwayIBAN generates Norway IBAN
func generateNorwayIBAN() (string, error) {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("NO")
	ascii.Digits(sb, 13)

	return ReplaceChecksum(sb.String())
}

// getNorwayBBAN retrieves BBAN structure from Norway IBAN
func getNorwayBBAN(iban string) (BBAN, error) {
	if len(iban) != 15 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 15: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:15],
		BankCode:         iban[4:8],
		BranchCode:       "",
		NationalChecksum: iban[14:15],
		AccountNumber:    iban[8:14],
	}, nil
}
