// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"
	"strings"

	"github.com/jacoelho/banking/ascii"
)

// ValidateRepublicOfAzerbaijanIBAN validates Republic Of Azerbaijan IBAN
func ValidateRepublicOfAzerbaijanIBAN(iban string) error {
	if len(iban) != 28 {
		return fmt.Errorf("unexpected length, want: 28: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "AZ" {
		return fmt.Errorf("static value rule, pos: 0, expected value: AZ, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:8]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
		return fmt.Errorf("range rule, start pos: 4, length: 4, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[8:28]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 8, length: 20, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateRepublicOfAzerbaijanIBAN generates Republic Of Azerbaijan IBAN
func GenerateRepublicOfAzerbaijanIBAN() string {
	var sb = new(strings.Builder)

	sb.WriteString("AZ")
	generator.Digits(sb, 2)
	generator.UpperCaseLetters(sb, 4)
	generator.AlphaNumeric(sb, 20)

	return ReplaceChecksum(sb.String())
}