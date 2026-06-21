package iban_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/jacoelho/banking/iban"
)

func TestValidationErrorMatchesInvalidIBAN(t *testing.T) {
	err := &iban.ValidationError{
		Reason:   iban.ReasonInvalidCharacters,
		Position: 4,
		Length:   3,
		Expected: iban.CharClassUpperAlpha,
		Actual:   "123",
	}

	if !errors.Is(err, iban.ErrInvalidIBAN) {
		t.Fatalf("errors.Is(err, ErrInvalidIBAN) = false, want true")
	}
	if errors.Is(err, iban.ErrUnsupportedCountry) {
		t.Fatalf("errors.Is(err, ErrUnsupportedCountry) = true, want false")
	}

	wrapped := fmt.Errorf("wrapped: %w", err)
	var got *iban.ValidationError
	if !errors.As(wrapped, &got) {
		t.Fatalf("errors.As(wrapped, *ValidationError) = false, want true")
	}
	if got.Reason != iban.ReasonInvalidCharacters ||
		got.Position != 4 ||
		got.Length != 3 ||
		got.Expected != iban.CharClassUpperAlpha ||
		got.Actual != "123" {
		t.Fatalf("ValidationError = %+v", got)
	}
}

func TestValidationErrorMatchesUnsupportedCountry(t *testing.T) {
	err := &iban.ValidationError{
		Reason:   iban.ReasonUnsupportedCountry,
		Position: 0,
		Length:   2,
		Expected: iban.CharClassUpperAlpha,
		Actual:   "ZZ",
	}

	if !errors.Is(err, iban.ErrInvalidIBAN) {
		t.Fatalf("errors.Is(err, ErrInvalidIBAN) = false, want true")
	}
	if !errors.Is(err, iban.ErrUnsupportedCountry) {
		t.Fatalf("errors.Is(err, ErrUnsupportedCountry) = false, want true")
	}
}

func TestCountryCodeErrorMatching(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want error
	}{
		{
			name: "invalid country code",
			err: &iban.CountryCodeError{
				CountryCode: "G",
				Err:         iban.ErrInvalidCountryCode,
			},
			want: iban.ErrInvalidCountryCode,
		},
		{
			name: "unsupported country",
			err: &iban.CountryCodeError{
				CountryCode: "ZZ",
				Err:         iban.ErrUnsupportedCountry,
			},
			want: iban.ErrUnsupportedCountry,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wrapped := fmt.Errorf("wrapped: %w", tt.err)
			if !errors.Is(wrapped, tt.want) {
				t.Fatalf("errors.Is(wrapped, %v) = false, want true", tt.want)
			}

			var got *iban.CountryCodeError
			if !errors.As(wrapped, &got) {
				t.Fatalf("errors.As(wrapped, *CountryCodeError) = false, want true")
			}
		})
	}
}

func TestValidateReturnsStructuredCharacterError(t *testing.T) {
	err := iban.Validate("GB29NWBK6016133192681X")
	if err == nil {
		t.Fatalf("Validate() error = nil, want error")
	}

	var got *iban.ValidationError
	if !errors.As(err, &got) {
		t.Fatalf("errors.As(err, *ValidationError) = false, want true")
	}
	if got.Reason != iban.ReasonInvalidCharacters ||
		got.Position != 8 ||
		got.Length != 14 ||
		got.Expected != iban.CharClassDigit ||
		got.Actual != "6016133192681X" {
		t.Fatalf("ValidationError = %+v", got)
	}
}

func TestValidateRejectsLowercaseAlphaNumericSpan(t *testing.T) {
	err := iban.Validate("FR1420041010050500013m02606")
	if err == nil {
		t.Fatalf("Validate() error = nil, want error")
	}

	var got *iban.ValidationError
	if !errors.As(err, &got) {
		t.Fatalf("errors.As(err, *ValidationError) = false, want true")
	}
	if got.Reason != iban.ReasonInvalidCharacters ||
		got.Position != 14 ||
		got.Length != 11 ||
		got.Expected != iban.CharClassUpperAlphaNumeric ||
		got.Actual != "0500013m026" {
		t.Fatalf("ValidationError = %+v", got)
	}
}
