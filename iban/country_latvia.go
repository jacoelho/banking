// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// ValidateLatviaIBAN validates Latvia IBAN
func ValidateLatviaIBAN(iban string) error {
	if len(iban) != 21 {
		return fmt.Errorf("unexpected length, want: 21: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "LV" {
		return fmt.Errorf("static value rule, pos: 0, expected value: LV, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:8]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
		return fmt.Errorf("range rule, start pos: 4, length: 4, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[8:21]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 8, length: 13, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateLatviaIBAN generates Latvia IBAN
func GenerateLatviaIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("LV")
	generator.Digits(sb, 2)
	generator.UpperCaseLetters(sb, 4)
	generator.AlphaNumeric(sb, 13)

	return ReplaceChecksum(sb.String())
}
