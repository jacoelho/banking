// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"
	"strings"

	"github.com/jacoelho/banking/ascii"
)

// ValidateSwedenIBAN validates Sweden IBAN
func ValidateSwedenIBAN(iban string) error {
	if len(iban) != 24 {
		return fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "SE" {
		return fmt.Errorf("static value rule, pos: 0, expected value: SE, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:24]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 22, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateSwedenIBAN generates Sweden IBAN
func GenerateSwedenIBAN() string {
	var sb = new(strings.Builder)

	sb.WriteString("SE")
	generator.Digits(sb, 22)

	return ReplaceChecksum(sb.String())
}
