package iban

import (
	"errors"
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	for _, country := range countrySpecs {
		countryCode := country.code
		t.Run(countryCode, func(t *testing.T) {
			t.Parallel()

			generated, err := Generate(countryCode)
			if err != nil {
				t.Fatal(err)
			}
			if err := Validate(generated); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestGenerateRejectsInvalidCountryCode(t *testing.T) {
	tests := []struct {
		name        string
		countryCode string
		want        error
	}{
		{
			name:        "too short",
			countryCode: "G",
			want:        ErrInvalidCountryCode,
		},
		{
			name:        "lowercase",
			countryCode: "gb",
			want:        ErrInvalidCountryCode,
		},
		{
			name:        "unsupported",
			countryCode: "ZZ",
			want:        ErrUnsupportedCountry,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.countryCode)
			if err == nil {
				t.Fatalf("Generate() error = nil, got %q", got)
			}
			assertCountryCodeError(t, err, tt.countryCode, tt.want)
		})
	}
}

func assertCountryCodeError(t *testing.T, err error, countryCode string, want error) {
	t.Helper()

	if !errors.Is(err, want) {
		t.Fatalf("errors.Is(err, %v) = false, want true", want)
	}

	var got *CountryCodeError
	if !errors.As(err, &got) {
		t.Fatalf("errors.As(err, *CountryCodeError) = false, want true")
	}
	if got.CountryCode != countryCode {
		t.Fatalf("CountryCodeError.CountryCode = %q, want %q", got.CountryCode, countryCode)
	}
}
