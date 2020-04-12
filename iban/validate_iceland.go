// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateIcelandIBAN(iban string) error {
    if len(iban) != 26 {
        return fmt.Errorf("unexpected length, want: 26: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "IS" {
        return fmt.Errorf("static value rule, pos: 0, expected value: IS, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:8]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 4, length: 4, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[8:10]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 8, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[10:16]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 10, length: 6, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[16:26]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 16, length: 10, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    return nil
}
