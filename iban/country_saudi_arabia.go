// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"
	"strings"

	"github.com/jacoelho/banking/ascii"
)

func ValidateSaudiArabiaIBAN(iban string) error {
	if len(iban) != 24 {
		return fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "SA" {
		return fmt.Errorf("static value rule, pos: 0, expected value: SA, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:6]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 4, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[6:24]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 6, length: 18, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

func GenerateSaudiArabiaIBAN() string {
	var sb = new(strings.Builder)
	sb.WriteString("SA")
	generator.Digits(sb, 2)
	generator.Digits(sb, 2)
	generator.AlphaNumeric(sb, 18)

	return ReplaceChecksum(sb.String())
}
