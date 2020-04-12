// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateSanMarinoIBAN(iban string) error {
    if len(iban) != 27 {
        return fmt.Errorf("unexpected length, want: 27: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "SM" {
        return fmt.Errorf("static value rule, pos: 0, expected value: SM, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:5]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
        return fmt.Errorf("range rule, start pos: 4, length: 1, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[5:10]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 5, length: 5, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[10:15]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 10, length: 5, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[15:27]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
        return fmt.Errorf("range rule, start pos: 15, length: 12, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
    }
    
	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

    return nil
}