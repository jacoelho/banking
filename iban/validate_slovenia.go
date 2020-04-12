// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateSloveniaIBAN(iban string) error {
    if len(iban) != 19 {
        return fmt.Errorf("unexpected length, want: 19: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "SI" {
        return fmt.Errorf("static value rule, pos: 0, expected value: SI, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:9]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 4, length: 5, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[9:17]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 9, length: 8, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[17:19]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 17, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

    return nil
}