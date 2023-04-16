// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// validateMongoliaIBAN validates Mongolia IBAN
func validateMongoliaIBAN(iban string) error {
	if len(iban) != 20 {
		return fmt.Errorf("unexpected length, want: 20: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "MN" {
		return fmt.Errorf("static value rule, pos: 0, expected value: MN, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:20]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 18, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// generateMongoliaIBAN generates Mongolia IBAN
func generateMongoliaIBAN() (string, error) {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("MN")
	ascii.Digits(sb, 18)

	return ReplaceChecksum(sb.String())
}

// getMongoliaBBAN retrieves BBAN structure from Mongolia IBAN
func getMongoliaBBAN(iban string) (BBAN, error) {
	if len(iban) != 20 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 20: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:20],
		BankCode:         iban[4:8],
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    iban[8:20],
	}, nil
}
