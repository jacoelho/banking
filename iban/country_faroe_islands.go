// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// ValidateFaroeIslandsIBAN validates Faroe Islands IBAN
func ValidateFaroeIslandsIBAN(iban string) error {
	if len(iban) != 18 {
		return fmt.Errorf("unexpected length, want: 18: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "FO" {
		return fmt.Errorf("static value rule, pos: 0, expected value: FO, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:18]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 16, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateFaroeIslandsIBAN generates Faroe Islands IBAN
func GenerateFaroeIslandsIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("FO")
	generator.Digits(sb, 16)

	return ReplaceChecksum(sb.String())
}

// GetFaroeIslandsBBAN retrieves BBAN structure from Faroe Islands IBAN
func GetFaroeIslandsBBAN(iban string) (BBAN, error) {
	if len(iban) != 18 {
		return BBAN{}, fmt.Errorf("unexpected length, want: 18: %w", ErrValidation)
	}

	return BBAN{
		BBAN:             iban[4:18],
		BankCode:         iban[4:8],
		BranchCode:       "",
		NationalChecksum: iban[17:18],
		AccountNumber:    iban[8:17],
	}, nil
}
