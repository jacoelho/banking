// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateNewCaledoniaIBAN(iban string) error {
    if len(iban) != 27 {
        return fmt.Errorf("unexpected length, want: 27: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "NC" {
        return fmt.Errorf("static value rule, pos: 0, expected value: NC, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:9]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 4, length: 5, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[9:14]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 9, length: 5, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[14:25]; !ascii.Every(subject, ascii.IsAlphaNumeric) {
        return fmt.Errorf("range rule, start pos: 14, length: 11, expected type AlphaNumeric, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[25:27]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 25, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

    return nil
}