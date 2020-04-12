// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"
	"strings"

	"github.com/jacoelho/banking/ascii"
)

// ValidateAndorraIBAN validates Andorra IBAN
func ValidateAndorraIBAN(iban string) error {
	if len(iban) != 24 {
		return fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "AD" {
		return fmt.Errorf("static value rule, pos: 0, expected value: AD, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:12]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 10, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[12:24]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 12, length: 12, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateAndorraIBAN generates Andorra IBAN
func GenerateAndorraIBAN() string {
	var sb = new(strings.Builder)

	sb.WriteString("AD")
	generator.Digits(sb, 10)
	generator.AlphaNumeric(sb, 12)

	return ReplaceChecksum(sb.String())
}
