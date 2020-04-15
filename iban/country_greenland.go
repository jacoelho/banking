// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateGreenlandIBAN validates Greenland IBAN
func validateGreenlandIBAN(iban string) error {
	if len(iban) != 18 {
		return fmt.Errorf("unexpected length, want: 18: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "GL" {
		return fmt.Errorf("static value rule, pos: 0, expected value: GL, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:18]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 16, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateGreenlandIBAN generates Greenland IBAN
func generateGreenlandIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("GL")
	generator.Digits(sb, 16)

	return replaceChecksum(sb.String())
}

// getGreenlandBBAN retrieves BBAN structure from Greenland IBAN
func getGreenlandBBAN(iban string) (BBAN, error) {
	if len(iban) != 18 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 18: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:18],
		BankCode:         iban[4:8],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[8:18],
	}, nil
}
