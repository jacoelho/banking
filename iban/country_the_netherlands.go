// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"
	"strings"

	"github.com/jacoelho/banking/ascii"
)

func ValidateTheNetherlandsIBAN(iban string) error {
	if len(iban) != 18 {
		return fmt.Errorf("unexpected length, want: 18: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "NL" {
		return fmt.Errorf("static value rule, pos: 0, expected value: NL, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:8]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
		return fmt.Errorf("range rule, start pos: 4, length: 4, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[8:18]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 8, length: 10, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

func GenerateTheNetherlandsIBAN() string {
	var sb = new(strings.Builder)
	sb.WriteString("NL")
	generator.Digits(sb, 2)
	generator.UpperCaseLetters(sb, 4)
	generator.Digits(sb, 10)

	return ReplaceChecksum(sb.String())
}
