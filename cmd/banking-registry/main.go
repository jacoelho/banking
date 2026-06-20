package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jacoelho/banking/internal/ibangen"
	"github.com/jacoelho/banking/internal/ibanregistry"
)

const defaultEncoding = string(ibanregistry.EncodingAuto)

func run(registryFile, dstDir, encodingName, normalizedPath string) error {
	if registryFile == "" {
		return fmt.Errorf("registry-file flag is required")
	}
	if dstDir == "" {
		return fmt.Errorf("dst-directory flag is required")
	}

	source, usedEncoding, err := ibanregistry.LoadFile(registryFile, ibanregistry.Encoding(encodingName))
	if err != nil {
		return fmt.Errorf("load registry: %w", err)
	}
	if encodingName == string(ibanregistry.EncodingAuto) {
		fmt.Fprintf(os.Stderr, "registry encoding: %s\n", usedEncoding)
	}

	if normalizedPath != "" {
		data := append(ibanregistry.NormalizedBytes(source), '\n')
		if writeErr := os.WriteFile(normalizedPath, data, 0o644); writeErr != nil {
			return fmt.Errorf("write normalized registry: %w", writeErr)
		}
	}

	registry, err := ibanregistry.ParseTSV(source)
	if err != nil {
		return fmt.Errorf("parse registry: %w", err)
	}
	if err := ibangen.Generate(dstDir, registry); err != nil {
		return fmt.Errorf("generate IBAN code: %w", err)
	}
	return nil
}

func runCLI(args []string) error {
	flags := flag.NewFlagSet("banking-registry", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	registryFile := flags.String("registry-file", "", "registry TSV file")
	dstDir := flags.String("dst-directory", "iban", "destination directory")
	encoding := flags.String("encoding", defaultEncoding, "registry encoding: utf8, latin1, or auto")
	normalizedPath := flags.String("write-normalized-registry", "", "optional path for normalized UTF-8 TSV output")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if err := run(*registryFile, *dstDir, *encoding, *normalizedPath); err != nil {
		flags.Usage()
		return err
	}
	return nil
}

func main() {
	if err := runCLI(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
