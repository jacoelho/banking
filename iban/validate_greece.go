// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateGreeceIBAN(iban string) error {
    if len(iban) != 27 {
        return fmt.Errorf("unexpected length, want: 27: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "GR" {
        return fmt.Errorf("static value rule, pos: 0, expected value: GR, found %s: %w", subject, ErrValidation)
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
    
    if subject := iban[11:27]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
        return fmt.Errorf("range rule, start pos: 11, length: 16, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
    }
    
	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

    return nil
}