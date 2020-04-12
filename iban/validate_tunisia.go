// Code generated DO NOT EDIT.

package iban

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func ValidateTunisiaIBAN(iban string) error {
    if len(iban) != 24 {
        return fmt.Errorf("unexpected length, want: 24: %w", ErrValidation)
    }
    
    if subject := iban[0:2]; subject != "TN" {
        return fmt.Errorf("static value rule, pos: 0, expected value: TN, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[2:4]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 2, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[4:6]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 4, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[6:9]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 6, length: 3, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[9:22]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 9, length: 13, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
    if subject := iban[22:24]; !ascii.Every(subject, ascii.IsDigit) {
        return fmt.Errorf("range rule, start pos: 22, length: 2, expected type Digit, found %s: %w", subject, ErrValidation)
    }
    
	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

    return nil
}