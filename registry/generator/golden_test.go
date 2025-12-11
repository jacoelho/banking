package generator

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jacoelho/banking/registry"
)

var updateGolden = flag.Bool("update-golden", false, "update golden files")

var testCountries = []registry.Country{
	{
		Code:          "AL",
		Name:          "Albania",
		IBAN:          "AL2!n8!n16!c",
		BBAN:          "8!n16!c",
		BankCode:      "0:3",
		BranchCode:    "3:7",
		AccountNumber: "8:24",
		IsSEPA:        false,
	},
	{
		Code:          "DE",
		Name:          "Germany",
		IBAN:          "DE2!n8!n10!n",
		BBAN:          "8!n10!n",
		BankCode:      "0:8",
		BranchCode:    "",
		AccountNumber: "8:18",
		IsSEPA:        true,
	},
	{
		Code:          "GB",
		Name:          "United Kingdom",
		IBAN:          "GB2!n4!a6!n8!n",
		BBAN:          "4!a6!n8!n",
		BankCode:      "0:4",
		BranchCode:    "4:10",
		AccountNumber: "10:18",
		IsSEPA:        false,
	},
}

func TestGenerateCodeForCountry_Golden(t *testing.T) {
	for _, country := range testCountries {
		t.Run(country.Name, func(t *testing.T) {
			var buf bytes.Buffer
			err := GenerateCodeForCountry(&buf, country)
			if err != nil {
				t.Fatalf("GenerateCodeForCountry failed: %v", err)
			}

			golden := filepath.Join("testdata", "country_"+country.Code+".go.golden")

			if *updateGolden {
				err := os.WriteFile(golden, buf.Bytes(), 0644)
				if err != nil {
					t.Fatalf("failed to update golden file: %v", err)
				}
				return
			}

			expected, err := os.ReadFile(golden)
			if err != nil {
				t.Fatalf("failed to read golden file %s: %v", golden, err)
			}

			if diff := cmp.Diff(string(expected), buf.String()); diff != "" {
				t.Errorf("GenerateCodeForCountry output mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGenerateValidate_Golden(t *testing.T) {
	var buf bytes.Buffer
	err := GenerateValidate(&buf, testCountries)
	if err != nil {
		t.Fatalf("GenerateValidate failed: %v", err)
	}

	golden := filepath.Join("testdata", "validate.go.golden")

	if *updateGolden {
		err := os.WriteFile(golden, buf.Bytes(), 0644)
		if err != nil {
			t.Fatalf("failed to update golden file: %v", err)
		}
		return
	}

	expected, err := os.ReadFile(golden)
	if err != nil {
		t.Fatalf("failed to read golden file %s: %v", golden, err)
	}

	if diff := cmp.Diff(string(expected), buf.String()); diff != "" {
		t.Errorf("GenerateValidate output mismatch (-want +got):\n%s", diff)
	}
}

func TestGenerateGenerate_Golden(t *testing.T) {
	var buf bytes.Buffer
	err := GenerateGenerate(&buf, testCountries)
	if err != nil {
		t.Fatalf("GenerateGenerate failed: %v", err)
	}

	golden := filepath.Join("testdata", "generate.go.golden")

	if *updateGolden {
		err := os.WriteFile(golden, buf.Bytes(), 0644)
		if err != nil {
			t.Fatalf("failed to update golden file: %v", err)
		}
		return
	}

	expected, err := os.ReadFile(golden)
	if err != nil {
		t.Fatalf("failed to read golden file %s: %v", golden, err)
	}

	if diff := cmp.Diff(string(expected), buf.String()); diff != "" {
		t.Errorf("GenerateGenerate output mismatch (-want +got):\n%s", diff)
	}
}

func TestGenerateGetBBAN_Golden(t *testing.T) {
	var buf bytes.Buffer
	err := GenerateGetBBAN(&buf, testCountries)
	if err != nil {
		t.Fatalf("GenerateGetBBAN failed: %v", err)
	}

	golden := filepath.Join("testdata", "get_bban.go.golden")

	if *updateGolden {
		err := os.WriteFile(golden, buf.Bytes(), 0644)
		if err != nil {
			t.Fatalf("failed to update golden file: %v", err)
		}
		return
	}

	expected, err := os.ReadFile(golden)
	if err != nil {
		t.Fatalf("failed to read golden file %s: %v", golden, err)
	}

	if diff := cmp.Diff(string(expected), buf.String()); diff != "" {
		t.Errorf("GenerateGetBBAN output mismatch (-want +got):\n%s", diff)
	}
}

func TestGenerateIsSEPA_Golden(t *testing.T) {
	var buf bytes.Buffer
	err := GenerateIsSEPA(&buf, testCountries)
	if err != nil {
		t.Fatalf("GenerateIsSEPA failed: %v", err)
	}

	golden := filepath.Join("testdata", "sepa.go.golden")

	if *updateGolden {
		err := os.WriteFile(golden, buf.Bytes(), 0644)
		if err != nil {
			t.Fatalf("failed to update golden file: %v", err)
		}
		return
	}

	expected, err := os.ReadFile(golden)
	if err != nil {
		t.Fatalf("failed to read golden file %s: %v", golden, err)
	}

	if diff := cmp.Diff(string(expected), buf.String()); diff != "" {
		t.Errorf("GenerateIsSEPA output mismatch (-want +got):\n%s", diff)
	}
}
