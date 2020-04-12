// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateCroatiaIBAN(iban string) error {
    if len(iban) != 21 {
        return fmt.Errorf("unexpected length, want: 21: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "HR" {
        return fmt.Errorf("static value rule, pos: 0, expected value: HR, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:11]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 4, length: 7, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[11:21]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 11, length: 10, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

    return nil
}