//go:build validation

package iban

import (
	"testing"
)

func TestIsSEPACountry(t *testing.T) {
	tests := []struct {
		countryCode string
		want        bool
		wantErr     bool
	}{
		{
			countryCode: "BG",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "ES",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "HR",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "CY",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "CZ",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "DK",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "EE",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "FI",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "FR",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "GF",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "DE",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "GI",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "GR",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "GP",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "HU",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "IS",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "IE",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "IT",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "LV",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "LI",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "LT",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "LU",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "PT",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "MT",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "MQ",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "YT",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "MC",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "NL",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "NO",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "PL",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "PT",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "RE",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "RO",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "BL",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "MF",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "SM",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "SK",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "SI",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "ES",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "SE",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "CH",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "GB",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "VA",
			want:        true,
			wantErr:     false,
		},
		{
			countryCode: "AE",
			want:        false,
			wantErr:     false,
		},
		{
			countryCode: "ZZ",
			want:        false,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.countryCode, func(t *testing.T) {
			got, err := IsSEPACountryCode(tt.countryCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsSEPACountry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsSEPACountryCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
