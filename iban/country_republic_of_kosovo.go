// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateRepublicOfKosovoIBAN validates Republic Of Kosovo IBAN
func validateRepublicOfKosovoIBAN(iban string) error {
	if len(iban) != 20 {
		return fmt.Errorf("unexpected length, want: 20: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "XK" {
		return fmt.Errorf("static value rule, pos: 0, expected value: XK, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:20]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 18, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateRepublicOfKosovoIBAN generates Republic Of Kosovo IBAN
func generateRepublicOfKosovoIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("XK")
	generator.Digits(sb, 18)

	return replaceChecksum(sb.String())
}

// getRepublicOfKosovoBBAN retrieves BBAN structure from Republic Of Kosovo IBAN
func getRepublicOfKosovoBBAN(iban string) (BBAN, error) {
	if len(iban) != 20 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 20: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:20],
		BankCode:         iban[4:6],
		BranchCode:       iban[6:8],
		NationalChecksum: "",
		AccountNumber:    iban[8:20],
	}, nil
}
