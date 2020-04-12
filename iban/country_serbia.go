// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"
	"strings"

	"github.com/jacoelho/banking/ascii"
)

func ValidateSerbiaIBAN(iban string) error {
	if len(iban) != 22 {
		return fmt.Errorf("unexpected length, want: 22: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "RS" {
		return fmt.Errorf("static value rule, pos: 0, expected value: RS, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:7]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 4, length: 3, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[7:20]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 7, length: 13, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[20:22]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 20, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

func GenerateSerbiaIBAN() string {
	var sb = new(strings.Builder)
	sb.WriteString("RS")
	generator.Digits(sb, 2)
	generator.Digits(sb, 3)
	generator.Digits(sb, 13)
	generator.Digits(sb, 2)

	return ReplaceChecksum(sb.String())
}
