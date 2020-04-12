// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"
	"github.com/jacoelho/banking/ascii"
)

func ValidateHungaryIBAN(iban string) error {
	if len(iban) != 28 {
		return fmt.Errorf("unexpected length, want: 28: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "HU" {
		return fmt.Errorf("static value rule, pos: 0, expected value: HU, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[4:7]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 4, length: 3, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[7:11]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 7, length: 4, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[11:12]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 11, length: 1, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[12:27]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 12, length: 15, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[27:28]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 27, length: 1, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}
