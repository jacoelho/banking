package ibangen

import (
	"flag"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"testing"

	"github.com/jacoelho/banking/internal/ibanregistry"
)

var updateGolden = flag.Bool("update", false, "update golden files")

func TestGenerateProducesValidGo(t *testing.T) {
	dir := t.TempDir()
	if err := Generate(dir, testRegistry()); err != nil {
		t.Fatal(err)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		if _, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.ParseComments); err != nil {
			t.Fatalf("generated file %s does not parse: %v", entry.Name(), err)
		}
	}
}

func TestGenerateMatchesGolden(t *testing.T) {
	dir := t.TempDir()
	if err := Generate(dir, testRegistry()); err != nil {
		t.Fatal(err)
	}

	got := readGeneratedFile(t, dir, registryDataFileName)
	goldenPath := filepath.Join("testdata", "registry_data.go.golden")
	if *updateGolden {
		if err := os.MkdirAll(filepath.Dir(goldenPath), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(goldenPath, []byte(got), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	want, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatal(err)
	}
	if got != string(want) {
		t.Fatalf("generated %s does not match golden\n--- got ---\n%s\n--- want ---\n%s", registryDataFileName, got, want)
	}
}

func TestGenerateRemovesOldGeneratedCountryFiles(t *testing.T) {
	dir := t.TempDir()
	old := filepath.Join(dir, "country_united_kingdom.go")
	if err := os.WriteFile(old, []byte(Header+"\n\npackage iban\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := Generate(dir, testRegistry()); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(old); !os.IsNotExist(err) {
		t.Fatalf("old generated file still exists: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, registryDataFileName)); err != nil {
		t.Fatalf("generated data file missing: %v", err)
	}
}

func TestGenerateOverwritesGeneratedDataFile(t *testing.T) {
	dir := t.TempDir()
	oldData := Header + "\n\npackage iban\n\nconst oldData = true\n"
	if err := os.WriteFile(filepath.Join(dir, registryDataFileName), []byte(oldData), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := Generate(dir, testRegistry()); err != nil {
		t.Fatal(err)
	}
	if got := readGeneratedFile(t, dir, registryDataFileName); got == oldData {
		t.Fatalf("%s was not overwritten", registryDataFileName)
	}
}

func TestGenerateRejectsManualRegistryDataFile(t *testing.T) {
	dir := t.TempDir()
	manual := filepath.Join(dir, registryDataFileName)
	if err := os.WriteFile(manual, []byte("package iban\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := Generate(dir, testRegistry()); err == nil {
		t.Fatal("Generate() error = nil")
	}
}

func TestGenerateRejectsManualLegacyCountryFile(t *testing.T) {
	dir := t.TempDir()
	manual := filepath.Join(dir, "country_manual.go")
	if err := os.WriteFile(manual, []byte("package iban\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := Generate(dir, testRegistry()); err == nil {
		t.Fatal("Generate() error = nil")
	}
}

func readGeneratedFile(t *testing.T, dir, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(dir, name))
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func testRegistry() ibanregistry.Registry {
	return ibanregistry.Registry{
		Countries: []ibanregistry.CountrySpec{
			{
				Code: "GB",
				Name: "United Kingdom",
				IBAN: ibanregistry.Structure{
					Length: 22,
					Rules: []ibanregistry.Rule{
						ibanregistry.StaticRule{StartPosition: 0, Value: "GB"},
						ibanregistry.RangeRule{StartPosition: 2, Length: 2, Format: ibanregistry.Digit},
						ibanregistry.RangeRule{StartPosition: 4, Length: 4, Format: ibanregistry.UpperCaseLetters},
						ibanregistry.RangeRule{StartPosition: 8, Length: 14, Format: ibanregistry.Digit},
					},
				},
				BBAN: ibanregistry.Structure{
					Length: 18,
					Rules: []ibanregistry.Rule{
						ibanregistry.RangeRule{StartPosition: 0, Length: 4, Format: ibanregistry.UpperCaseLetters},
						ibanregistry.RangeRule{StartPosition: 4, Length: 14, Format: ibanregistry.Digit},
					},
				},
				BankCode: ibanregistry.OptionalSpan{
					Present: true,
					Span:    ibanregistry.Span{Start: 0, End: 4},
				},
				BranchCode: ibanregistry.OptionalSpan{
					Present: true,
					Span:    ibanregistry.Span{Start: 4, End: 10},
				},
				AccountNumber: ibanregistry.Span{Start: 10, End: 18},
				IsSEPA:        true,
			},
		},
	}
}
