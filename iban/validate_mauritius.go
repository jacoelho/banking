// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateMauritiusIBAN(iban string) error {
    if len(iban) != 30 {
        return fmt.Errorf("unexpected length, want: 30: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "MU" {
        return fmt.Errorf("static value rule, pos: 0, expected value: MU, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:8]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
        return fmt.Errorf("range rule, start pos: 4, length: 4, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[8:10]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 8, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[10:12]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 10, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[12:24]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 12, length: 12, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[24:27]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 24, length: 3, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[27:30]; !ascii.Every(subject, ascii.IsUpperCaseLetter) {
        return fmt.Errorf("range rule, start pos: 27, length: 3, expected type UpperCaseLetters, found %s: %w", subject, ErrValidation)
    }
    
	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

    return nil
}