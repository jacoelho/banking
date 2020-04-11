package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/jacoelho/iban/registry"
	"golang.org/x/text/encoding/charmap"
)

func main() {
	fileName := flag.String("filename", "", "registry file")

	flag.Parse()

	fileReader, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("failed to open file: %s", err)
		os.Exit(1)
	}
	defer fileReader.Close()

	entries, err := registry.Parse(charmap.Windows1252.NewDecoder().Reader(fileReader))
	if err != nil {
		fmt.Printf("failed to parse file: %s", err)
		os.Exit(1)
	}

	_ = json.NewEncoder(os.Stdout).Encode(&entries)
}
