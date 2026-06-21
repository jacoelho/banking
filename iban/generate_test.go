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

func TestGenerateWithBBAN(t *testing.T) {
	t.Parallel()

	got, err := GenerateWithBBAN("GB", BBANParts{
		BankCode:      "NWBK",
		BranchCode:    "601613",
		AccountNumber: "31926819",
	})
	if err != nil {
		t.Fatal(err)
	}
	if got != "GB29NWBK60161331926819" {
		t.Fatalf("GenerateWithBBAN() = %q, want %q", got, "GB29NWBK60161331926819")
	}
}

func TestGenerateWithBBANPreservesPartialParts(t *testing.T) {
	t.Parallel()

	got, err := GenerateWithBBAN("GB", BBANParts{
		BankCode:      "NWBK",
		AccountNumber: "31926819",
	})
	if err != nil {
		t.Fatal(err)
	}

	bban, err := GetBBAN(got)
	if err != nil {
		t.Fatal(err)
	}
	if bban.BankCode != "NWBK" {
		t.Fatalf("BankCode = %q, want %q", bban.BankCode, "NWBK")
	}
	if bban.AccountNumber != "31926819" {
		t.Fatalf("AccountNumber = %q, want %q", bban.AccountNumber, "31926819")
	}
}

func TestGenerateWithBBANPreservesUpperAlphaNumericPart(t *testing.T) {
	t.Parallel()

	got, err := GenerateWithBBAN("FR", BBANParts{
		AccountNumber: "12345ABCDEFGHIJK12",
	})
	if err != nil {
		t.Fatal(err)
	}

	bban, err := GetBBAN(got)
	if err != nil {
		t.Fatal(err)
	}
	if bban.AccountNumber != "12345ABCDEFGHIJK12" {
		t.Fatalf("AccountNumber = %q, want %q", bban.AccountNumber, "12345ABCDEFGHIJK12")
	}
}

func TestGenerateWithEmptyBBANParts(t *testing.T) {
	t.Parallel()

	for _, country := range countrySpecs {
		countryCode := country.code
		t.Run(countryCode, func(t *testing.T) {
			t.Parallel()

			generated, err := GenerateWithBBAN(countryCode, BBANParts{})
			if err != nil {
				t.Fatal(err)
			}
			if err := Validate(generated); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestGenerateWithBBANGeneratesNonFieldBBANCharacters(t *testing.T) {
	t.Parallel()

	tests := []struct {
		countryCode string
		parts       BBANParts
	}{
		{
			countryCode: "IT",
			parts: BBANParts{
				BankCode:      "05428",
				BranchCode:    "11101",
				AccountNumber: "000000123456",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.countryCode, func(t *testing.T) {
			t.Parallel()

			generated, err := GenerateWithBBAN(tt.countryCode, tt.parts)
			if err != nil {
				t.Fatal(err)
			}
			bban, err := GetBBAN(generated)
			if err != nil {
				t.Fatal(err)
			}
			if bban.BankCode != tt.parts.BankCode {
				t.Fatalf("BankCode = %q, want %q", bban.BankCode, tt.parts.BankCode)
			}
			if bban.BranchCode != tt.parts.BranchCode {
				t.Fatalf("BranchCode = %q, want %q", bban.BranchCode, tt.parts.BranchCode)
			}
			if bban.AccountNumber != tt.parts.AccountNumber {
				t.Fatalf("AccountNumber = %q, want %q", bban.AccountNumber, tt.parts.AccountNumber)
			}

			wantSuffix := tt.parts.BankCode + tt.parts.BranchCode + tt.parts.AccountNumber
			if bban.BBAN[1:] != wantSuffix {
				t.Fatalf("BBAN suffix = %q, want %q", bban.BBAN[1:], wantSuffix)
			}
			if bban.BBAN[0] < 'A' || bban.BBAN[0] > 'Z' {
				t.Fatalf("BBAN first byte = %q, want uppercase letter", bban.BBAN[0])
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

func TestGenerateWithBBANRejectsInvalidCountryCode(t *testing.T) {
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
			got, err := GenerateWithBBAN(tt.countryCode, BBANParts{
				BankCode: "bad input is ignored until country is valid",
			})
			if err == nil {
				t.Fatalf("GenerateWithBBAN() error = nil, got %q", got)
			}
			assertCountryCodeError(t, err, tt.countryCode, tt.want)
		})
	}
}

func TestGenerateWithBBANRejectsInvalidParts(t *testing.T) {
	tests := []struct {
		name        string
		countryCode string
		parts       BBANParts
		want        BBANPartsError
	}{
		{
			name:        "wrong bank code length",
			countryCode: "GB",
			parts: BBANParts{
				BankCode: "NWB",
			},
			want: BBANPartsError{
				CountryCode:    "GB",
				Field:          "BankCode",
				Length:         3,
				ExpectedLength: 4,
				ActualLength:   3,
				Actual:         "NWB",
			},
		},
		{
			name:        "lowercase bank code",
			countryCode: "GB",
			parts: BBANParts{
				BankCode: "nwbk",
			},
			want: BBANPartsError{
				CountryCode:    "GB",
				Field:          "BankCode",
				Length:         4,
				ExpectedLength: 4,
				ActualLength:   4,
				Expected:       CharClassUpperAlpha,
				Actual:         "nwbk",
			},
		},
		{
			name:        "branch supplied for absent span",
			countryCode: "DE",
			parts: BBANParts{
				BranchCode: "123",
			},
			want: BBANPartsError{
				CountryCode:    "DE",
				Field:          "BranchCode",
				Length:         3,
				ExpectedLength: 0,
				ActualLength:   3,
				Actual:         "123",
			},
		},
		{
			name:        "mixed rule account number",
			countryCode: "FR",
			parts: BBANParts{
				AccountNumber: "12345!!!!!!!!!!!12",
			},
			want: BBANPartsError{
				CountryCode:    "FR",
				Field:          "AccountNumber",
				Position:       5,
				Length:         11,
				ExpectedLength: 18,
				ActualLength:   18,
				Expected:       CharClassUpperAlphaNumeric,
				Actual:         "!!!!!!!!!!!",
			},
		},
		{
			name:        "lowercase mixed rule account number",
			countryCode: "FR",
			parts: BBANParts{
				AccountNumber: "12345abcabcabcde12",
			},
			want: BBANPartsError{
				CountryCode:    "FR",
				Field:          "AccountNumber",
				Position:       5,
				Length:         11,
				ExpectedLength: 18,
				ActualLength:   18,
				Expected:       CharClassUpperAlphaNumeric,
				Actual:         "abcabcabcde",
			},
		},
		{
			name:        "trailing digit rule in mixed account number",
			countryCode: "FR",
			parts: BBANParts{
				AccountNumber: "12345ABCDEFGHIJKAB",
			},
			want: BBANPartsError{
				CountryCode:    "FR",
				Field:          "AccountNumber",
				Position:       16,
				Length:         2,
				ExpectedLength: 18,
				ActualLength:   18,
				Expected:       CharClassDigit,
				Actual:         "AB",
			},
		},
		{
			name:        "first invalid field in struct order",
			countryCode: "GB",
			parts: BBANParts{
				BankCode:      "nwbk",
				BranchCode:    "ABCDEF",
				AccountNumber: "bad",
			},
			want: BBANPartsError{
				CountryCode:    "GB",
				Field:          "BankCode",
				Length:         4,
				ExpectedLength: 4,
				ActualLength:   4,
				Expected:       CharClassUpperAlpha,
				Actual:         "nwbk",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateWithBBAN(tt.countryCode, tt.parts)
			if err == nil {
				t.Fatalf("GenerateWithBBAN() error = nil, got %q", got)
			}
			assertBBANPartsError(t, err, tt.want)
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

func assertBBANPartsError(t *testing.T, err error, want BBANPartsError) {
	t.Helper()

	if !errors.Is(err, ErrInvalidBBANParts) {
		t.Fatalf("errors.Is(err, ErrInvalidBBANParts) = false, want true")
	}

	var got *BBANPartsError
	if !errors.As(err, &got) {
		t.Fatalf("errors.As(err, *BBANPartsError) = false, want true")
	}
	if *got != want {
		t.Fatalf("BBANPartsError = %+v, want %+v", *got, want)
	}
}
