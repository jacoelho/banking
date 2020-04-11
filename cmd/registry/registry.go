package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jacoelho/iban/registry"
	"golang.org/x/text/encoding/charmap"
	yaml "gopkg.in/yaml.v2"
)

type Country struct {
	Code string `yaml:"code"`
	Name string `yaml:"name"`
	IBAN string `yaml:"IBAN"`
	BBAN string ` yaml:"BBAN"`
}

func main() {
	fileName := flag.String("filename", "", "registry file")

	flag.Parse()

	fileReader, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("failed to open file: %s", err)
		os.Exit(1)
	}
	defer fileReader.Close()

	entries, err := registry.Decode(charmap.Windows1252.NewDecoder().Reader(fileReader))
	if err != nil {
		fmt.Printf("failed to parse file: %s", err)
		os.Exit(1)
	}

	var countries []Country
	for _, c := range entries {
		c := Country{
			Code: c.CountryCode,
			Name: c.CountryName,
			IBAN: c.IBAN.Structure,
			BBAN: c.BBAN.Structure,
		}

		countries = append(countries, c)
	}

	wrap := struct {
		Countries []Country `yaml:"countries"`
	}{
		Countries: countries,
	}
	yaml.NewEncoder(os.Stdout).Encode(&wrap)
}
