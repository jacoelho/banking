package registry

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		source  io.Reader
		want    []Entry
		wantErr bool
	}{
		"empty reader": {
			source:  nil,
			want:    nil,
			wantErr: true,
		},
		"only headers": {
			source: strings.NewReader(`Name of country	Country code as defined in ISO 3166	Domestic account number example	BBAN	BBAN structure 	BBAN length	Bank identifier position within the BBAN	Bank identifier length	Bank identifier example	BBAN example	IBAN	IBAN structure	IBAN length	IBAN electronic format example	IBAN print format example	SEPA Country	Valid from date	Contact details`),
			want:    nil,
			wantErr: false,
		},
		"one entry": {
			source: strings.NewReader(`Name of country	Country code as defined in ISO 3166	Domestic account number example	BBAN	BBAN structure 	BBAN length	Bank identifier position within the BBAN	Bank identifier length	Bank identifier example	BBAN example	IBAN	IBAN structure	IBAN length	IBAN electronic format example	IBAN print format example	SEPA Country	Valid from date	Contact details
Albania	AL	0000000235698741	0	8!n16!c	24	"Bank Identifier 1-3, Branch Identifier:4-7, Check Digit 8"	8!n	212-1100-9	212110090000000235698741	0	AL2!n8!n16!c	28	AL47212110090000000235698741	AL47 2121 1009 0000 0002 3569 8741	No		"Miho Valer , Deputy Director, Payment Systems Department, BANK OF ALBANIA, Kompleksi Halili, Rruga e Dibres, 1000 TIRANA, ALBANIA, Tel: 355 4 2419301/2/3 ext 3061, Fax: 355 4 2419408 , Email: vmiho@bankofalbania.org"`),
			want: []Entry{
				{
					CountryName:                  "Albania",
					CountryCode:                  "AL",
					DomesticAccountNumberExample: "0000000235698741",
					BBAN: BBAN{
						Structure:              "8!n16!c",
						Length:                 "24",
						BankIdentifierPosition: "Bank Identifier 1-3, Branch Identifier:4-7, Check Digit 8",
						BankIdentifierLength:   "8!n",
						BankIdentifierExample:  "212-1100-9",
						Example:                "212110090000000235698741",
					},
					IBAN: IBAN{
						Structure:               "AL2!n8!n16!c",
						Length:                  "28",
						ElectronicFormatExample: "AL47212110090000000235698741",
						PrintFormatExample:      "AL47 2121 1009 0000 0002 3569 8741",
					},
				},
			},
			wantErr: false,
		},
	}
	for tC, tt := range tests {
		t.Run(tC, func(t *testing.T) {
			got, err := Parse(tt.source)
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
