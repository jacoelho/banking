// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateBurundiIBAN validates Burundi IBAN
func validateBurundiIBAN(iban string) error {
	if len(iban) != 27 {
		return fmt.Errorf("unexpected length, want: 27: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "BI" {
		return fmt.Errorf("static value rule, pos: 0, expected value: BI, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:27]; !ascii.IsDigit(subject) {
		return fmt.Errorf("range rule, start pos: 2, length: 25, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateBurundiIBAN generates Burundi IBAN
func generateBurundiIBAN() (string, error) {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("BI")
	ascii.Digits(sb, 25)

	return ReplaceChecksum(sb.String())
}

// getBurundiBBAN retrieves BBAN structure from Burundi IBAN
func getBurundiBBAN(iban string) (BBAN, error) {
	if len(iban) != 27 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 27: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:27],
		BankCode:         iban[4:9],
		BranchCode:       iban[9:14],
		NationalChecksum: "",
		AccountNumber:    iban[14:27],
	}, nil
}
