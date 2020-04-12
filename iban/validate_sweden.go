// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateSwedenIBAN(iban string) error {
    if len(iban) != 24 {
        return fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "SE" {
        return fmt.Errorf("static value rule, pos: 0, expected value: SE, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:7]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 4, length: 3, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[7:23]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 7, length: 16, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[23:24]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 23, length: 1, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    return nil
}
