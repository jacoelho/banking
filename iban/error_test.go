package iban_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/jacoelho/banking/iban"
)

func TestIsValidationError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "non-validation error",
			err:  fmt.Errorf("some other error"),
			want: false,
		},
		{
			name: "ErrValidationLength",
			err:  &iban.ErrValidationLength{Expected: 22, Actual: 20},
			want: true,
		},
		{
			name: "ErrValidationChecksum",
			err:  &iban.ErrValidationChecksum{Expected: "97", Actual: "00"},
			want: true,
		},
		{
			name: "ErrValidationRange",
			err:  &iban.ErrValidationRange{Position: 4, Length: 8, Expected: iban.CharacterTypeDigit, Actual: "ABC12345"},
			want: true,
		},
		{
			name: "ErrValidationStaticValue",
			err:  &iban.ErrValidationStaticValue{Position: 0, Expected: "GB", Actual: "XX"},
			want: true,
		},
		{
			name: "ErrUnsupportedCountry",
			err:  &iban.ErrUnsupportedCountry{CountryCode: "XX"},
			want: true,
		},
		{
			name: "wrapped validation error",
			err:  fmt.Errorf("wrapper: %w", &iban.ErrValidationLength{Expected: 22, Actual: 20}),
			want: true,
		},
		{
			name: "wrapped non-validation error",
			err:  fmt.Errorf("wrapper: %w", errors.New("not validation")),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := iban.IsValidationError(tt.err)
			if got != tt.want {
				t.Errorf("IsValidationError() = %v, want %v", got, tt.want)
			}
		})
	}
}
