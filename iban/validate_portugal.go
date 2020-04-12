// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidatePortugalIBAN(iban string) error {
    if len(iban) != 25 {
        return fmt.Errorf("unexpected length, want: 25: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "PT" {
        return fmt.Errorf("static value rule, pos: 0, expected value: PT, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:8]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 4, length: 4, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[8:12]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 8, length: 4, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[12:23]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 12, length: 11, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[23:25]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 23, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    return nil
}
