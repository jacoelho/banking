package ibanregistry

import (
	"strings"
	"testing"
)

func TestParseTSVExpandsTerritoriesAndSEPA(t *testing.T) {
	registry, err := ParseTSV(validTSV())
	if err != nil {
		t.Fatalf("ParseTSV() error = %v", err)
	}

	countries := make(map[string]CountrySpec)
	for _, country := range registry.Countries {
		countries[country.Code] = country
	}
	for _, code := range []string{"FI", "FR", "PT", "GB", "AX", "GF", "PF", "IM", "JE", "GG"} {
		if _, ok := countries[code]; !ok {
			t.Fatalf("country %s not parsed", code)
		}
	}

	wantSEPA := map[string]bool{
		"AX": true,
		"GF": true,
		"PF": false,
		"IM": true,
		"JE": true,
		"GG": true,
		"PT": true,
	}
	for code, want := range wantSEPA {
		if got := countries[code].IsSEPA; got != want {
			t.Fatalf("SEPA %s = %v, want %v", code, got, want)
		}
	}
}

func TestParseTSVCompactsContiguousRules(t *testing.T) {
	registry, err := ParseTSV(validTSV())
	if err != nil {
		t.Fatalf("ParseTSV() error = %v", err)
	}

	countries := make(map[string]CountrySpec)
	for _, country := range registry.Countries {
		countries[country.Code] = country
	}

	assertRules(t, countries["PT"].IBAN.Rules, []Rule{
		StaticRule{StartPosition: 0, Value: "PT"},
		RangeRule{StartPosition: 2, Length: 2, Format: Digit},
		RangeRule{StartPosition: 4, Length: 21, Format: Digit},
	})
	assertRules(t, countries["PT"].BBAN.Rules, []Rule{
		RangeRule{StartPosition: 0, Length: 21, Format: Digit},
	})
	assertRules(t, countries["GB"].IBAN.Rules, []Rule{
		StaticRule{StartPosition: 0, Value: "GB"},
		RangeRule{StartPosition: 2, Length: 2, Format: Digit},
		RangeRule{StartPosition: 4, Length: 4, Format: UpperCaseLetters},
		RangeRule{StartPosition: 8, Length: 14, Format: Digit},
	})
}

func TestDecodeAutoUsesLatin1ForInvalidUTF8(t *testing.T) {
	got, used, err := Decode([]byte{'A', 0xc5}, EncodingAuto)
	if err != nil {
		t.Fatal(err)
	}
	if used != EncodingLatin1 {
		t.Fatalf("encoding = %s, want latin1", used)
	}
	if got != "AÅ" {
		t.Fatalf("decoded = %q, want %q", got, "AÅ")
	}
}

func TestParseTSVRejectsInvalidInput(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
		{
			name:   "duplicate required row",
			source: validTSV() + "SEPA country\tYes\tYes\tYes\tYes\n",
		},
		{
			name:   "missing SEPA includes row",
			source: removeRow(validTSV(), "SEPA country also includes"),
		},
		{
			name:   "invalid boolean",
			source: strings.Replace(validTSV(), "SEPA country\tYes\tYes\tYes\tYes", "SEPA country\tYes\tMaybe\tYes\tYes", 1),
		},
		{
			name:   "reversed span",
			source: strings.Replace(validTSV(), "Bank identifier position within the BBAN\t1-3", "Bank identifier position within the BBAN\t5-4", 1),
		},
		{
			name:   "column mismatch",
			source: strings.Replace(validTSV(), "IBAN structure\tFI2!n3!n11!n\tFR2!n5!n5!n11!c2!n\tPT2!n4!n4!n11!n2!n\tGB2!n4!a6!n8!n", "IBAN structure\tFI2!n3!n11!n", 1),
		},
		{
			name:   "SEPA include not owned by parent",
			source: strings.Replace(validTSV(), "SEPA country also includes\tAX\tGF", "SEPA country also includes\tGF\tGF", 1),
		},
		{
			name:   "duplicate SEPA include token",
			source: strings.Replace(validTSV(), "SEPA country also includes\tAX", "SEPA country also includes\t\"AX, AX\"", 1),
		},
		{
			name:   "malformed SEPA include token",
			source: strings.Replace(validTSV(), "SEPA country also includes\tAX", "SEPA country also includes\tax", 1),
		},
		{
			name:   "SEPA annotation under wrong parent",
			source: strings.Replace(validTSV(), "SEPA country also includes\tAX", "SEPA country also includes\tAzores", 1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ParseTSV(tt.source); err == nil {
				t.Fatal("ParseTSV() error = nil")
			}
		})
	}
}

func assertRules(t *testing.T, got, want []Rule) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("len(rules) = %d, want %d\ngot: %#v\nwant: %#v", len(got), len(want), got, want)
	}
	for i := range got {
		if !rulesEqual(got[i], want[i]) {
			t.Fatalf("rule %d = %#v, want %#v", i, got[i], want[i])
		}
	}
}

func validTSV() string {
	return strings.Join([]string{
		"ignored\t\"multi\nline\"\tignored\tignored\tignored",
		"Name of country\tFinland\tFrance\tPortugal\tUnited Kingdom",
		"IBAN prefix country code (ISO 3166)\tFI\tFR\tPT\tGB",
		"Country code includes other countries/territories\tAX\t\"GF, PF\"\tN/A\t\"IM, JE, GG\"",
		"SEPA country\tYes\tYes\tYes\tYes",
		"SEPA country also includes\tAX\tGF\t\"Azores, Madeira\"\tN/A",
		"BBAN structure\t3!n11!n\t5!n5!n11!c2!n\t4!n4!n11!n2!n\t4!a6!n8!n",
		"Bank identifier position within the BBAN\t1-3\t1-5\t1-4\t1-4",
		"Branch identifier position within the BBAN\tN/A\tN/A\t5-8\t5-10",
		"IBAN structure\tFI2!n3!n11!n\tFR2!n5!n5!n11!c2!n\tPT2!n4!n4!n11!n2!n\tGB2!n4!a6!n8!n",
	}, "\r\n") + "\r\n"
}

func removeRow(source, label string) string {
	var rows []string
	for row := range strings.SplitSeq(source, "\r\n") {
		if row == "" || strings.HasPrefix(row, label+"\t") {
			continue
		}
		rows = append(rows, row)
	}
	return strings.Join(rows, "\r\n") + "\r\n"
}
