package tsv

import (
	"strings"
	"testing"

	"github.com/jacoelho/banking/registry"
)

func TestConvertPosition(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantStart int
		wantEnd   int
		wantErr   bool
	}{
		{
			name:      "valid position",
			input:     "1-4",
			wantStart: 0,
			wantEnd:   4,
		},
		{
			name:      "with whitespace",
			input:     " 3-7 ",
			wantStart: 2,
			wantEnd:   7,
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
		},
		{
			name:    "N/A",
			input:   "N/A",
			wantErr: true,
		},
		{
			name:    "invalid format",
			input:   "1",
			wantErr: true,
		},
		{
			name:    "invalid number",
			input:   "abc-4",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, end, err := convertPosition(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertPosition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if start != tt.wantStart {
					t.Errorf("convertPosition() start = %v, want %v", start, tt.wantStart)
				}
				if end != tt.wantEnd {
					t.Errorf("convertPosition() end = %v, want %v", end, tt.wantEnd)
				}
			}
		})
	}
}

func TestParseBoolean(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "yes", input: "yes", want: true},
		{name: "no", input: "no", want: false},
		{name: "case insensitive", input: "YES", want: true},
		{name: "whitespace", input: "  Yes  ", want: true},
		{name: "empty", input: "", want: false},
		{name: "invalid", input: "Maybe", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseBoolean(tt.input)
			if got != tt.want {
				t.Errorf("parseBoolean(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestNormalizeCountryName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "parenthetical suffix",
			input: "United Arab Emirates (The)",
			want:  "United Arab Emirates",
		},
		{
			name:  "quoted comma pattern",
			input: `"Moldova, Republic of"`,
			want:  "Moldova",
		},
		{
			name:  "no changes",
			input: "Albania",
			want:  "Albania",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "whitespace and parenthetical",
			input: "Netherlands (The)  ",
			want:  "Netherlands",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normaliseCountryName(tt.input)
			if got != tt.want {
				t.Errorf("normalizeCountryName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestParseCountries(t *testing.T) {
	// Create minimal TSV data with label column (column 0) + 3 country columns
	tsvData := strings.Join([]string{
		"Data element\tDescription\tDescription\tDescription",                               // Row 0: Headers
		"Name of country\tAndorra\tUnited Arab Emirates\tAlbania",                           // Row 1: Country names (label + data)
		"IBAN prefix country code (ISO 3166)\tAD\tAE\tAL",                                   // Row 2: Country codes (label + data)
		"Country code includes other countries/territories\tN/A\tN/A\tN/A",                  // Row 3: Includes
		"SEPA country\tYes\tNo\tNo",                                                         // Row 4: SEPA
		"SEPA country also includes\tN/A\tN/A\tN/A",                                         // Row 5: SEPA includes
		"Domestic account number example\texample\texample\texample",                        // Row 6: Domestic example
		"BBAN\tBBAN\tBBAN\tBBAN",                                                            // Row 7: Label
		"BBAN structure\t4!n4!n12!c\t3!n16!n\t8!n16!c",                                      // Row 8: BBAN structure
		"BBAN length\t20\t19\t24",                                                           // Row 9: BBAN length
		"Bank identifier position within the BBAN\t1-4\t1-3\t1-3",                           // Row 10: Bank position
		"Bank identifier pattern\t4!n\t3!n\t3!n",                                            // Row 11: Bank pattern
		"Branch identifier position within the BBAN\t5-8\t\t4-7",                            // Row 12: Branch position
		"Branch identifier pattern\t4!n\t\t5!n",                                             // Row 13: Branch pattern
		"Bank identifier example\t0001\t033\t212",                                           // Row 14: Bank example
		"Branch identifier example\t2030\t\t11009",                                          // Row 15: Branch example
		"BBAN example\t00012030200359100100\t0331234567890123456\t212110090000000235698741", // Row 16: BBAN example
		"IBAN\tIBAN\tIBAN\tIBAN",                                                            // Row 17: Label
		"IBAN structure\tAD2!n4!n4!n12!c\tAE2!n3!n16!n\tAL2!n8!n16!c",                       // Row 18: IBAN structure
		"IBAN length\t24\t23\t28",                                                           // Row 19: IBAN length
		"Effective date\tApr-07\tOct-11\tApr-09",                                            // Row 20: Effective date
		"IBAN electronic format example\tAD1200012030200359100100\tAE070331234567890123456\tAL47212110090000000235698741",            // Row 21: IBAN electronic
		"IBAN print format example\tAD12 0001 2030 2003 5910 0100\tAE07 0331 2345 6789 0123 456\tAL47 2121 1009 0000 0002 3569 8741", // Row 22: IBAN print
	}, "\n")

	parser := NewParser(strings.NewReader(tsvData))
	countries, err := parser.ParseCountries()
	if err != nil {
		t.Fatalf("ParseCountries() error = %v", err)
	}

	if len(countries) != 3 {
		t.Fatalf("ParseCountries() got %d countries, want 3", len(countries))
	}

	// Test Andorra (SEPA country with branch code)
	ad := countries[0]
	if ad.Code != "AD" {
		t.Errorf("Country 0 Code = %v, want AD", ad.Code)
	}
	if ad.Name != "Andorra" {
		t.Errorf("Country 0 Name = %v, want Andorra", ad.Name)
	}
	if ad.IBAN != "AD2!n4!n4!n12!c" {
		t.Errorf("Country 0 IBAN = %v, want AD2!n4!n4!n12!c", ad.IBAN)
	}
	if !ad.IsSEPA {
		t.Errorf("Country 0 IsSEPA = false, want true")
	}
	if ad.BankCode != "0:4" {
		t.Errorf("Country 0 BankCode = %v, want 0:4", ad.BankCode)
	}
	if ad.BranchCode != "4:8" {
		t.Errorf("Country 0 BranchCode = %v, want 4:8", ad.BranchCode)
	}

	// Test UAE (non-SEPA, no branch code)
	ae := countries[1]
	if ae.Code != "AE" {
		t.Errorf("Country 1 Code = %v, want AE", ae.Code)
	}
	if ae.IsSEPA {
		t.Errorf("Country 1 IsSEPA = true, want false")
	}
	if ae.BranchCode != "" {
		t.Errorf("Country 1 BranchCode = %v, want empty", ae.BranchCode)
	}

	// Test Albania (non-SEPA with branch code)
	al := countries[2]
	if al.Code != "AL" {
		t.Errorf("Country 2 Code = %v, want AL", al.Code)
	}
	if al.BranchCode != "3:7" {
		t.Errorf("Country 2 BranchCode = %v, want 3:7", al.BranchCode)
	}
}

func TestDeriveAccountPosition(t *testing.T) {
	tests := []struct {
		name      string
		country   registry.Country
		bankEnd   int
		branchEnd int
		want      string
		wantErr   bool
	}{
		{
			name: "with bank and branch",
			country: registry.Country{
				IBAN: "AD2!n4!n4!n12!c",
			},
			bankEnd:   4,
			branchEnd: 8,
			want:      "8:20",
		},
		{
			name: "bank only",
			country: registry.Country{
				IBAN: "AE2!n3!n16!n",
			},
			bankEnd: 3,
			want:    "3:19",
		},
		{
			name: "no codes",
			country: registry.Country{
				IBAN: "DE2!n8!n10!n",
			},
			want: "0:18",
		},
		{
			name: "empty IBAN",
			country: registry.Country{
				IBAN: "",
			},
			wantErr: true,
		},
		{
			name: "invalid IBAN structure",
			country: registry.Country{
				IBAN: "XX2!x4!y",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deriveAccountPosition(&tt.country, tt.bankEnd, tt.branchEnd)
			if (err != nil) != tt.wantErr {
				t.Errorf("deriveAccountPosition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("deriveAccountPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseCountries_Errors(t *testing.T) {
	tests := []struct {
		name    string
		tsvData string
		wantErr string
	}{
		{
			name: "missing country code row",
			tsvData: strings.Join([]string{
				"Data element\tDescription",
				"Name of country\tAndorra",
				"SEPA country\tYes",
			}, "\n"),
			wantErr: "missing required row",
		},
		{
			name: "empty country code",
			tsvData: strings.Join([]string{
				"Data element\tDescription\tDescription",
				"Name of country\tAndorra\tAlbania",
				"IBAN prefix country code (ISO 3166)\tAD\t",
				"Country code includes other countries/territories\tN/A\tN/A",
				"SEPA country\tYes\tNo",
				"SEPA country also includes\tN/A\tN/A",
				"Domestic account number example\texample\texample",
				"BBAN\tBBAN\tBBAN",
				"BBAN structure\t4!n4!n12!c\t8!n16!c",
				"BBAN length\t20\t24",
				"Bank identifier position within the BBAN\t1-4\t1-3",
				"Bank identifier pattern\t4!n\t3!n",
				"Branch identifier position within the BBAN\t5-8\t4-7",
				"Branch identifier pattern\t4!n\t5!n",
				"Bank identifier example\t0001\t212",
				"Branch identifier example\t2030\t11009",
				"BBAN example\t00012030200359100100\t212110090000000235698741",
				"IBAN\tIBAN\tIBAN",
				"IBAN structure\tAD2!n4!n4!n12!c\tAL2!n8!n16!c",
			}, "\n"),
			wantErr: "empty country code",
		},
		{
			name: "invalid bank position format",
			tsvData: strings.Join([]string{
				"Data element\tDescription",
				"Name of country\tAndorra",
				"IBAN prefix country code (ISO 3166)\tAD",
				"Country code includes other countries/territories\tN/A",
				"SEPA country\tYes",
				"SEPA country also includes\tN/A",
				"Domestic account number example\texample",
				"BBAN\tBBAN",
				"BBAN structure\t4!n4!n12!c",
				"BBAN length\t20",
				"Bank identifier position within the BBAN\tinvalid",
				"Bank identifier pattern\t4!n",
				"Branch identifier position within the BBAN\t5-8",
				"Branch identifier pattern\t4!n",
				"Bank identifier example\t0001",
				"Branch identifier example\t2030",
				"BBAN example\t00012030200359100100",
				"IBAN\tIBAN",
				"IBAN structure\tAD2!n4!n4!n12!c",
			}, "\n"),
			wantErr: "invalid bank position",
		},
		{
			name: "invalid branch position format",
			tsvData: strings.Join([]string{
				"Data element\tDescription",
				"Name of country\tAndorra",
				"IBAN prefix country code (ISO 3166)\tAD",
				"Country code includes other countries/territories\tN/A",
				"SEPA country\tYes",
				"SEPA country also includes\tN/A",
				"Domestic account number example\texample",
				"BBAN\tBBAN",
				"BBAN structure\t4!n4!n12!c",
				"BBAN length\t20",
				"Bank identifier position within the BBAN\t1-4",
				"Bank identifier pattern\t4!n",
				"Branch identifier position within the BBAN\tbad-format",
				"Branch identifier pattern\t4!n",
				"Bank identifier example\t0001",
				"Branch identifier example\t2030",
				"BBAN example\t00012030200359100100",
				"IBAN\tIBAN",
				"IBAN structure\tAD2!n4!n4!n12!c",
			}, "\n"),
			wantErr: "invalid branch position",
		},
		{
			name: "empty IBAN structure",
			tsvData: strings.Join([]string{
				"Data element\tDescription",
				"Name of country\tAndorra",
				"IBAN prefix country code (ISO 3166)\tAD",
				"Country code includes other countries/territories\tN/A",
				"SEPA country\tYes",
				"SEPA country also includes\tN/A",
				"Domestic account number example\texample",
				"BBAN\tBBAN",
				"BBAN structure\t4!n4!n12!c",
				"BBAN length\t20",
				"Bank identifier position within the BBAN\t1-4",
				"Bank identifier pattern\t4!n",
				"Branch identifier position within the BBAN\t5-8",
				"Branch identifier pattern\t4!n",
				"Bank identifier example\t0001",
				"Branch identifier example\t2030",
				"BBAN example\t00012030200359100100",
				"IBAN\tIBAN",
				"IBAN structure\t",
			}, "\n"),
			wantErr: "cannot derive account position: IBAN is empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := NewParser(strings.NewReader(tt.tsvData))
			_, err := parser.ParseCountries()
			if err == nil {
				t.Errorf("ParseCountries() expected error containing %q, got nil", tt.wantErr)
				return
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("ParseCountries() error = %q, want error containing %q", err.Error(), tt.wantErr)
			}
		})
	}
}
