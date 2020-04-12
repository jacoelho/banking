// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateBulgariaIBAN(iban string) error {
    if len(iban) != 22 {
        return fmt.Errorf("unexpected length, want: 22: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "BG" {
        return fmt.Errorf("static value rule, pos: 0, expected value: BG, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:8]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
        return fmt.Errorf("range rule, start pos: 4, length: 4, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[8:12]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 8, length: 4, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[12:14]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 12, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[14:22]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
        return fmt.Errorf("range rule, start pos: 14, length: 8, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
    }
    
	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

    return nil
}