// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateTurkeyIBAN(iban string) error {
    if len(iban) != 26 {
        return fmt.Errorf("unexpected length, want: 26: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "TR" {
        return fmt.Errorf("static value rule, pos: 0, expected value: TR, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:9]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 4, length: 5, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[9:10]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 9, length: 1, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[10:26]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
        return fmt.Errorf("range rule, start pos: 10, length: 16, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
    }
    
    return nil
}
