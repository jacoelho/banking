// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/pool"
)

// ValidateWallisAndFutunaIslandsIBAN validates Wallis And Futuna Islands IBAN
func ValidateWallisAndFutunaIslandsIBAN(iban string) error {
	if len(iban) != 27 {
		return fmt.Errorf("unexpected length, want: 27: %w", ErrValidation)
	}

	if subject := iban[0:2]; subject != "WF" {
		return fmt.Errorf("static value rule, pos: 0, expected value: WF, found %s: %w", subject, ErrValidation)
	}

	if subject := iban[2:14]; !ascii.Every(subject, ascii.IsDigit) {
		return fmt.Errorf("range rule, start pos: 2, length: 12, expected type Digit, found %s: %w", subject, ErrValidation)
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

// GenerateWallisAndFutunaIslandsIBAN generates Wallis And Futuna Islands IBAN
func GenerateWallisAndFutunaIslandsIBAN() string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.WriteString("WF")
	generator.Digits(sb, 12)
	generator.AlphaNumeric(sb, 11)
	generator.Digits(sb, 2)

	return ReplaceChecksum(sb.String())
}
