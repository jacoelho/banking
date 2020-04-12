// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateTimorLesteIBAN(iban string) error {
    if len(iban) != 23 {
        return fmt.Errorf("unexpected length, want: 23: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "TL" {
        return fmt.Errorf("static value rule, pos: 0, expected value: TL, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:7]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 4, length: 3, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[7:21]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 7, length: 14, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[21:23]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 21, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

    return nil
}
