package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/jacoelho/banking/registry"
	"github.com/jacoelho/banking/registry/generator"
)

// File permissions constants
const (
	dirPerm  = 0700 // Directory permission: rwx------
	filePerm = 0600 // File permission: rw-------
)

func createDirectory(dirName string) error {
	stat, err := os.Stat(dirName)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(dirName, dirPerm)
		}
		return err
	}

	if !stat.Mode().IsDir() {
		return fmt.Errorf("%s found, but is not a directory", dirName)
	}

	return nil
}

func validateFlags(fileName, destDir string) error {
	if fileName == "" {
		return errors.New("registry-file flag is required")
	}
	if destDir == "" {
		return errors.New("dst-directory flag is required")
	}
	return nil
}

func run(fileName, destDir string) error {
	if err := createDirectory(destDir); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	f, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open registry file: %w", err)
	}
	defer f.Close()

	var countries registry.Countries
	if err := yaml.NewDecoder(f).Decode(&countries); err != nil {
		return fmt.Errorf("failed to decode registry file: %w", err)
	}

	for _, country := range countries.Countries {
		targetFileName := "country_" + strings.ReplaceAll(strings.ToLower(country.Name), " ", "_") + ".go"
		targetFile := path.Join(destDir, targetFileName)

		if err := generateCountryFile(targetFile, country); err != nil {
			return fmt.Errorf("failed to generate code for %s: %w", country.Name, err)
		}
	}

	// Generate aggregate files
	files := []struct {
		name string
		fn   func(io.Writer, []registry.Country) error
	}{
		{"validate.go", generator.GenerateValidate},
		{"generate.go", generator.GenerateGenerate},
		{"bban_helper.go", generator.GenerateGetBBAN},
		{"sepa.go", generator.GenerateIsSEPA},
	}

	for _, file := range files {
		targetFile := path.Join(destDir, file.name)
		if err := generateFunc(targetFile, file.fn, countries.Countries); err != nil {
			return fmt.Errorf("failed to generate %s: %w", file.name, err)
		}
	}

	return nil
}

func main() {
	fileName := flag.String("registry-file", "", "registry file (yaml format)")
	destinationDirectory := flag.String("dst-directory", "iban", "destination directory")

	flag.Parse()

	if err := validateFlags(*fileName, *destinationDirectory); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*fileName, *destinationDirectory); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// generateCountryFile generates code for a single country with proper resource management
func generateCountryFile(filename string, country registry.Country) error {
	writer, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, filePerm)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer func() {
		if closeErr := writer.Close(); closeErr != nil {
			// Log but don't override the main error
			fmt.Fprintf(os.Stderr, "Warning: failed to close file %s: %v\n", filename, closeErr)
		}
	}()

	if err := generator.GenerateCodeForCountry(writer, country); err != nil {
		// Remove incomplete file on error
		if rmErr := os.Remove(filename); rmErr != nil {
			return fmt.Errorf("failed to generate code and failed to cleanup %s: %w (cleanup error: %v)", filename, err, rmErr)
		}
		return err
	}

	return nil
}

// generateFunc generates aggregate files with proper resource management
func generateFunc(filename string, fn func(io.Writer, []registry.Country) error, countries []registry.Country) error {
	writer, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, filePerm)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer func() {
		if closeErr := writer.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close file %s: %v\n", filename, closeErr)
		}
	}()

	if err := fn(writer, countries); err != nil {
		// Remove incomplete file on error
		if rmErr := os.Remove(filename); rmErr != nil {
			return fmt.Errorf("failed to generate content and failed to cleanup %s: %w (cleanup error: %v)", filename, err, rmErr)
		}
		return err
	}

	return nil
}
