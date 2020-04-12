// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"
	"strings"

	"github.com/jacoelho/banking/ascii"
)

func ValidateKazakhstanIBAN(iban string) error {
	if len(iban) != 20 {
		return fmt.Errorf("unexpected length, want: 20: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "KZ" {
		return fmt.Errorf("static value rule, pos: 0, expected value: KZ, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:7]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 4, length: 3, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[7:20]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 7, length: 13, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

func GenerateKazakhstanIBAN() string {
	var sb = new(strings.Builder)
	sb.WriteString("KZ")
	generator.Digits(sb, 2)
	generator.Digits(sb, 3)
	generator.AlphaNumeric(sb, 13)

	return ReplaceChecksum(sb.String())
}
