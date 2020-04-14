package main

import (
	"flag"
	"fmt"
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

	writer, err := os.OpenFile(targetFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	if err := generator.GenerateValidate(writer, countries.Countries); err != nil {
		writer.Close()
		if errRemove := os.Remove(targetFile); errRemove != nil {
			log.Fatalf("while handling %s, got %s", err, errRemove)
		}
		log.Fatal(err)
	}

	writer.Close()

	// single generate file
	targetFile = path.Join(*destinationDirectory, "generate.go")

	writer, err = os.OpenFile(targetFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	if err := generator.GenerateGenerate(writer, countries.Countries); err != nil {
		writer.Close()
		if errRemove := os.Remove(targetFile); errRemove != nil {
			log.Fatalf("while handling %s, got %s", err, errRemove)
		}
		log.Fatal(err)
	}

	writer.Close()

	// single generate file
	targetFile = path.Join(*destinationDirectory, "bban_helper.go")

	writer, err = os.OpenFile(targetFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	if err := generator.GenerateGetBBAN(writer, countries.Countries); err != nil {
		writer.Close()
		if errRemove := os.Remove(targetFile); errRemove != nil {
			log.Fatalf("while handling %s, got %s", err, errRemove)
		}
		log.Fatal(err)
	}

	writer.Close()
}
