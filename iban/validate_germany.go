// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateGermanyIBAN(iban string) error {
    if len(iban) != 22 {
        return fmt.Errorf("unexpected length, want: 22: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "DE" {
        return fmt.Errorf("static value rule, pos: 0, expected value: DE, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:12]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 4, length: 8, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[12:22]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 12, length: 10, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    return nil
}
