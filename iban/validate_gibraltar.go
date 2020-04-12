// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateGibraltarIBAN(iban string) error {
    if len(iban) != 23 {
        return fmt.Errorf("unexpected length, want: 23: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "GI" {
        return fmt.Errorf("static value rule, pos: 0, expected value: GI, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:8]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
        return fmt.Errorf("range rule, start pos: 4, length: 4, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[8:23]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
        return fmt.Errorf("range rule, start pos: 8, length: 15, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
    }
    
	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

    return nil
}
