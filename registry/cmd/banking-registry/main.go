package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/jacoelho/banking/registry"
	"github.com/jacoelho/banking/registry/generator"
)

func createDirectory(dirName string) error {
	stat, err := os.Stat(dirName)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(dirName, 0700)
		}
		return err
	}

	if !stat.Mode().IsDir() {
		return fmt.Errorf("%s found, but is not a directory", dirName)
	}

	return nil
}

func main() {
	fileName := flag.String("registry-file", "", "registry file (yaml format)")
	destinationDirectory := flag.String("dst-directory", "iban", "destination directory")

	flag.Parse()

	if *fileName == "" || *destinationDirectory == "" {
		log.Fatal("missing flags")
	}

	if err := createDirectory(*destinationDirectory); err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var countries registry.Countries
	if err := yaml.NewDecoder(f).Decode(&countries); err != nil {
		log.Fatal(err)
	}

	for _, country := range countries.Countries {
		targetFileName := "country_" + strings.ReplaceAll(strings.ToLower(country.Name), " ", "_") + ".go"
		targetFile := path.Join(*destinationDirectory, targetFileName)

		writer, err := os.OpenFile(targetFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
		if err != nil {
			log.Fatal(err)
		}

		if err := generator.GenerateCodeForCountry(writer, country); err != nil {
			writer.Close()
			if errRemove := os.Remove(targetFileName); errRemove != nil {
				log.Fatalf("while handling %s, got %s", err, errRemove)
			}
			log.Fatal(err)
		}

		writer.Close()
	}

	// single validate file
	targetFile := path.Join(*destinationDirectory, "validate.go")
	generateFunc(targetFile, generator.GenerateValidate, countries.Countries)

	// single generate file
	targetFile = path.Join(*destinationDirectory, "generate.go")
	generateFunc(targetFile, generator.GenerateGenerate, countries.Countries)

	// single generate file
	targetFile = path.Join(*destinationDirectory, "bban_helper.go")
	generateFunc(targetFile, generator.GenerateGetBBAN, countries.Countries)

	// single generate file
	targetFile = path.Join(*destinationDirectory, "country_constants.go")
	generateFunc(targetFile, generator.GenerateConstants, countries.Countries)

	// single generate file
	targetFile = path.Join(*destinationDirectory, "sepa.go")
	generateFunc(targetFile, generator.GenerateIsSEPA, countries.Countries)
}

func generateFunc(filename string, fn func(io.Writer, []registry.Country) error, countries []registry.Country) {
	writer, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	if err := fn(writer, countries); err != nil {
		writer.Close()
		if errRemove := os.Remove(filename); errRemove != nil {
			log.Fatalf("while handling %s, got %s", err, errRemove)
		}
		log.Fatal(err)
	}

	writer.Close()
}
