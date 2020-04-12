// +build validation

package iban

import (
	"testing"
)

func TestValidateAlbaniaIBAN(t *testing.T) {

	tests := []struct {
		name    string
		iban    string
		wantErr bool
	}{
		{
			name:    "albania",
			iban:    "AL37212110090000000235698A41",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateAlbaniaIBAN(tt.iban); (err != nil) != tt.wantErr {
				t.Errorf("ValidateAlbaniaIBAN() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
