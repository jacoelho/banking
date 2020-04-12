// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateBrazilIBAN(iban string) error {
    if len(iban) != 29 {
        return fmt.Errorf("unexpected length, want: 29: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "BR" {
        return fmt.Errorf("static value rule, pos: 0, expected value: BR, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:12]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 4, length: 8, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[12:17]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 12, length: 5, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[17:27]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 17, length: 10, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[27:28]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
        return fmt.Errorf("range rule, start pos: 27, length: 1, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[28:29]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
        return fmt.Errorf("range rule, start pos: 28, length: 1, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
    }
    
	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

    return nil
}