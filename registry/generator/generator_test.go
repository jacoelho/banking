package generator

import (
	"bytes"
	"testing"

	"github.com/jacoelho/banking/registry"
)

func TestGenerateValidationForCountry(t *testing.T) {
	tests := []struct {
		name    string
		country registry.Country
		result  string
		wantErr bool
	}{
		{
			name: "test",
			country: registry.Country{
				Code:       "AL",
				Name:       "Albania",
				IBAN:       "AL2!n8!n16!c",
				BankCode:   "0:4",
				BranchCode: "4:6",
				//NationalChecksum: "6:8",
				AccountNumber: "8:10",
			},
			result:  "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Skip()
			w := &bytes.Buffer{}
			err := GenerateCodeForCountry(w, tt.country)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateCodeForCountry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.result {
				t.Errorf("GenerateCodeForCountry() gotW = %v, want %v", gotW, tt.result)
			}
		})
	}
}
