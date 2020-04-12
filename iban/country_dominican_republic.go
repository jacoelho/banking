// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"
	"strings"

	"github.com/jacoelho/banking/ascii"
)

// ValidateDominicanRepublicIBAN validates Dominican Republic IBAN
func ValidateDominicanRepublicIBAN(iban string) error {
	if len(iban) != 28 {
		return fmt.Errorf("unexpected length, want: 28: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "DO" {
		return fmt.Errorf("static value rule, pos: 0, expected value: DO, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:8]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 4, length: 4, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[8:28]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 8, length: 20, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateDominicanRepublicIBAN generates Dominican Republic IBAN
func GenerateDominicanRepublicIBAN() string {
	var sb = new(strings.Builder)

	sb.WriteString("DO")
	generator.Digits(sb, 2)
	generator.AlphaNumeric(sb, 4)
	generator.Digits(sb, 20)

	return ReplaceChecksum(sb.String())
}
