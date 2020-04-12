// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"
	"strings"

	"github.com/jacoelho/banking/ascii"
)

// ValidateMacedoniaIBAN validates Macedonia IBAN
func ValidateMacedoniaIBAN(iban string) error {
	if len(iban) != 19 {
		return fmt.Errorf("unexpected length, want: 19: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "MK" {
		return fmt.Errorf("static value rule, pos: 0, expected value: MK, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:7]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 5, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[7:17]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 7, length: 10, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[17:19]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 17, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateMacedoniaIBAN generates Macedonia IBAN
func GenerateMacedoniaIBAN() string {
	var sb = new(strings.Builder)

	sb.WriteString("MK")
	generator.Digits(sb, 5)
	generator.AlphaNumeric(sb, 10)
	generator.Digits(sb, 2)

	return ReplaceChecksum(sb.String())
}
