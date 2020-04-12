// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateFaroeIslandsIBAN(iban string) error {
    if len(iban) != 18 {
        return fmt.Errorf("unexpected length, want: 18: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "FO" {
        return fmt.Errorf("static value rule, pos: 0, expected value: FO, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:8]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 4, length: 4, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[8:17]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 8, length: 9, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[17:18]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 17, length: 1, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    return nil
}
