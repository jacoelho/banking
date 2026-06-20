package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunCLIDefaultEncodingAcceptsLatin1Registry(t *testing.T) {
	dir := t.TempDir()
	registryPath := filepath.Join(dir, "registry.txt")
	dst := filepath.Join(dir, "iban")

	source := append([]byte("ignored\t"), 0xc5, '\n')
	source = append(source, []byte(minimalRegistryTSV())...)
	if err := os.WriteFile(registryPath, source, 0o644); err != nil {
		t.Fatal(err)
	}

	if err := runCLI([]string{"-registry-file", registryPath, "-dst-directory", dst}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dst, "registry_data.go")); err != nil {
		t.Fatalf("generated registry data file missing: %v", err)
	}
}

func minimalRegistryTSV() string {
	return strings.Join([]string{
		"Name of country\tFinland",
		"IBAN prefix country code (ISO 3166)\tFI",
		"Country code includes other countries/territories\tN/A",
		"SEPA country\tYes",
		"SEPA country also includes\tN/A",
		"BBAN structure\t3!n11!n",
		"Bank identifier position within the BBAN\t1-3",
		"Branch identifier position within the BBAN\tN/A",
		"IBAN structure\tFI2!n3!n11!n",
	}, "\n") + "\n"
}
