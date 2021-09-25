package bic_test

import (
	"github.com/jacoelho/banking/bic"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    bic.BIC
		wantErr bool
	}{
		{
			name:  "bic.BIC 8",
			input: "ABCDBEB1",
			want: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "XXX",
			},
			wantErr: false,
		},
		{
			name:  "bic.BIC 11 XXX",
			input: "ABCDBEB1XXX",
			want: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "XXX",
			},
			wantErr: false,
		},
		{
			name:  "bic.BIC 11 branch code with letters",
			input: "ABCDBEB1ABC",
			want: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "ABC",
			},
			wantErr: false,
		},
		{
			name:  "bic.BIC 11 branch code with digits",
			input: "ABCDBEB1ABC",
			want: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "ABC",
			},
			wantErr: false,
		},
		{
			name:  "bic.BIC 11 branch code with mixed characters",
			input: "ABCDBEB1A1C",
			want: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "A1C",
			},
			wantErr: false,
		},
		{
			name:    "only alphabetic values allowed in the first 4 characters (SWIFT)",
			input:   "A2CDBEB1",
			wantErr: true,
		},
		{
			name:    "country code (ZZ) not valid",
			input:   "ABCDZZB1",
			wantErr: true,
		},
		{
			name:    "less than 8 characters",
			input:   "ABCDBE",
			wantErr: true,
		},
		{
			name:    "more than 8, less than 11 characters",
			input:   "ABCDBEB1A1",
			wantErr: true,
		},
		{
			name:    "more than 11 characters",
			input:   "ABCDBEB1A123",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bic.Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name                string
		BusinessPartyPrefix string
		CountryCode         string
		BusinessPartySuffix string
		BranchCode          string
		want                bool
	}{
		{
			name:                "valid bic.BIC",
			BusinessPartyPrefix: "ABCD",
			CountryCode:         "BE",
			BusinessPartySuffix: "B1",
			BranchCode:          "XXX",
			want:                true,
		},
		{
			name: "zero value",
			want: false,
		},
		{
			name:                "missing business party prefix",
			BusinessPartyPrefix: "",
			CountryCode:         "BE",
			BusinessPartySuffix: "B1",
			BranchCode:          "XXX",
			want:                false,
		},
		{
			name:                "missing business party suffix",
			BusinessPartyPrefix: "ABCD",
			CountryCode:         "BE",
			BusinessPartySuffix: "",
			BranchCode:          "XXX",
			want:                false,
		},
		{
			name:                "missing country code",
			BusinessPartyPrefix: "ABCD",
			CountryCode:         "",
			BusinessPartySuffix: "B1",
			BranchCode:          "XXX",
			want:                false,
		},
		{
			name:                "missing branch code",
			BusinessPartyPrefix: "ABCD",
			CountryCode:         "BE",
			BusinessPartySuffix: "B1",
			BranchCode:          "",
			want:                false,
		},
		{
			name:                "branch length not valid",
			BusinessPartyPrefix: "ABCD",
			CountryCode:         "BE",
			BusinessPartySuffix: "B1",
			BranchCode:          "XXXX",
			want:                false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bic.BIC{
				BusinessPartyPrefix: tt.BusinessPartyPrefix,
				CountryCode:         tt.CountryCode,
				BusinessPartySuffix: tt.BusinessPartySuffix,
				BranchCode:          tt.BranchCode,
			}
			if got := b.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSameInstitution(t *testing.T) {
	tests := []struct {
		name  string
		left  bic.BIC
		right bic.BIC
		want  bool
	}{
		{
			name: "same values",
			left: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "XXX",
			},
			right: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "XXX",
			},
			want: true,
		},
		{
			name: "left contains a specific branch code",
			left: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "123",
			},
			right: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "XXX",
			},
			want: true,
		},
		{
			name: "right contains a specific branch code",
			left: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "XXX",
			},
			right: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "123",
			},
			want: true,
		},
		{
			name: "different countries",
			left: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "XXX",
			},
			right: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "DE",
				BusinessPartySuffix: "B1",
				BranchCode:          "XXX",
			},
			want: false,
		},
		{
			name: "different business party prefix",
			left: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "XXX",
			},
			right: bic.BIC{
				BusinessPartyPrefix: "DCBA",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "XXX",
			},
			want: false,
		},
		{
			name: "left is not valid",
			left: bic.BIC{
				BusinessPartyPrefix: "ABCD",
			},
			right: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "XXX",
			},
			want: false,
		},
		{
			name: "right is not valid",
			left: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "XXX",
			},
			right: bic.BIC{
				BusinessPartyPrefix: "ABCD",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.left.SameInstitution(tt.right); got != tt.want {
				t.Errorf("SameInstitution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBIC8(t *testing.T) {
	tests := []struct {
		name   string
		bic bic.BIC
		want   string
	}{
		{
			name: "happy case",
			bic: bic.BIC{
				BusinessPartyPrefix: "ABCD",
				CountryCode:         "BE",
				BusinessPartySuffix: "B1",
				BranchCode:          "XXX",
			},
			want: "ABCDBEB1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bic.BIC8(); got != tt.want {
				t.Errorf("bic.BIC8() = %v, want %v", got, tt.want)
			}
		})
	}
}
