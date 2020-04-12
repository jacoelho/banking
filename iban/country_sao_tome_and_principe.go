// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"
	"strings"

	"github.com/jacoelho/banking/ascii"
)

// ValidateSaoTomeAndPrincipeIBAN validates Sao Tome And Principe IBAN
func ValidateSaoTomeAndPrincipeIBAN(iban string) error {
	if len(iban) != 25 {
		return fmt.Errorf("unexpected length, want: 25: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "ST" {
		return fmt.Errorf("static value rule, pos: 0, expected value: ST, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:25]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 23, expected type Digit, found %s: %w", subject, ErrValidation)
	}

	if c := Checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

	return nil
}

// GenerateSaoTomeAndPrincipeIBAN generates Sao Tome And Principe IBAN
func GenerateSaoTomeAndPrincipeIBAN() string {
	var sb = new(strings.Builder)

	sb.WriteString("ST")
	generator.Digits(sb, 23)

	return ReplaceChecksum(sb.String())
}
