// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateSaintLuciaIBAN(iban string) error {
    if len(iban) != 32 {
        return fmt.Errorf("unexpected length, want: 32: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "LC" {
        return fmt.Errorf("static value rule, pos: 0, expected value: LC, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:8]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
        return fmt.Errorf("range rule, start pos: 4, length: 4, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[8:32]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
        return fmt.Errorf("range rule, start pos: 8, length: 24, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
    }
    
	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

    return nil
}
