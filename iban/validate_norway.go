// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateNorwayIBAN(iban string) error {
    if len(iban) != 15 {
        return fmt.Errorf("unexpected length, want: 15: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "NO" {
        return fmt.Errorf("static value rule, pos: 0, expected value: NO, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:8]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 4, length: 4, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[8:14]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 8, length: 6, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[14:15]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 14, length: 1, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    return nil
}
