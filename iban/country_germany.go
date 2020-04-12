// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"
	"strings"

	"github.com/jacoelho/banking/ascii"
)

// ValidateGermanyIBAN validates Germany IBAN
func ValidateGermanyIBAN(iban string) error {
	if len(iban) != 22 {
		return fmt.Errorf("unexpected length, want: 22: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "DE" {
		return fmt.Errorf("static value rule, pos: 0, expected value: DE, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:22]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 20, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateGermanyIBAN generates Germany IBAN
func GenerateGermanyIBAN() string {
	var sb = new(strings.Builder)

	sb.WriteString("DE")
	generator.Digits(sb, 20)

	return ReplaceChecksum(sb.String())
}