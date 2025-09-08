package generator

import (
	"bytes"
	"flag"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jacoelho/banking/registry"
)

var updateRegressionGolden = flag.Bool("update-regression-golden", false, "update regression golden files")

// Regression tests ensure generated output remains consistent over time
func TestGenerateCodeForCountry_Regression(t *testing.T) {
	for _, country := range testCountries {
		t.Run(country.Name, func(t *testing.T) {
			var buf bytes.Buffer
			err := GenerateCodeForCountry(&buf, country)
			if err != nil {
				t.Fatalf("GenerateCodeForCountry failed: %v", err)
			}

			golden := filepath.Join("testdata", "regression_country_"+country.Code+".go.golden")

			if *updateRegressionGolden {
				err := os.WriteFile(golden, buf.Bytes(), 0644)
				if err != nil {
					t.Fatalf("failed to update generator golden file: %v", err)
				}
				return
			}

			expected, err := os.ReadFile(golden)
			if err != nil {
				t.Fatalf("failed to read generator golden file %s: %v", golden, err)
			}

			if diff := cmp.Diff(string(expected), buf.String()); diff != "" {
				t.Errorf("GenerateCodeForCountry output mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGenerateValidate_Regression(t *testing.T) {
	var buf bytes.Buffer
	err := GenerateValidate(&buf, testCountries)
	if err != nil {
		t.Fatalf("GenerateValidate failed: %v", err)
	}
	output := buf.Bytes()

	golden := filepath.Join("testdata", "regression_validate.go.golden")

	if *updateRegressionGolden {
		err := os.WriteFile(golden, output, 0644)
		if err != nil {
			t.Fatalf("failed to update generator golden file: %v", err)
		}
		return
	}

	expected, err := os.ReadFile(golden)
	if err != nil {
		t.Fatalf("failed to read generator golden file %s: %v", golden, err)
	}

	if diff := cmp.Diff(string(expected), string(output)); diff != "" {
		t.Errorf("GenerateValidate output mismatch (-want +got):\n%s", diff)
	}
}

// TestGeneratedCodeCompilation verifies that generated code compiles successfully
func TestGeneratedCodeSyntax(t *testing.T) {
	country := registry.Country{
		Code:             "DE",
		Name:             "Germany",
		IBAN:             "DE2!n8!n10!n",
		BBAN:             "8!n10!n",
		BankCode:         "0:8",
		BranchCode:       "",
		NationalChecksum: "",
		AccountNumber:    "8:18",
		IsSEPA:           true,
	}

	t.Run("generated code compiles successfully", func(t *testing.T) {
		// Generate code
		var buf bytes.Buffer
		err := GenerateCodeForCountry(&buf, country)
		if err != nil {
			t.Fatalf("code generation failed: %v", err)
		}

		// Verify it's valid Go
		fset := token.NewFileSet()
		_, err = parser.ParseFile(fset, "generated.go", buf.String(), parser.ParseComments)
		if err != nil {
			t.Errorf("generated invalid Go code: %v\n%s", err, buf.String())
		}
	})
}
