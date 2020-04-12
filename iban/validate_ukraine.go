// Code generated DO NOT EDIT.

package iban

import (
	"fmt"
	"github.com/jacoelho/banking/ascii"
)

func ValidateUkraineIBAN(iban string) error {
	if len(iban) != 29 {
		return fmt.Errorf("unexpected length, want: 29: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "UA" {
		return fmt.Errorf("static value rule, pos: 0, expected value: UA, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:10]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 4, length: 6, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[10:29]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
		return fmt.Errorf("range rule, start pos: 10, length: 19, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}
