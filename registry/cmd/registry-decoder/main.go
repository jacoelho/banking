package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := flag.String("filename", "", "registry file")

	flag.Parse()

	if *fileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	fileReader, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer fileReader.Close()

	reader := csv.NewReader(fileReader)
	reader.Comma = '|'

	// skip header
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(line)
	}
	fmt.Println("end")
}
