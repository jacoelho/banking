package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/jacoelho/banking/registry"
	"github.com/jacoelho/banking/registry/parser/tsv"
)

const yamlTemplate = `# Generated from {{.SourceFile}} using tsv-to-yaml
# Do not edit manually - regenerate from TSV source

countries:
{{- range .Countries}}
  - code: {{.Code}}
    name: {{.Name}}
    IBAN: {{.IBAN}}
    {{- if .BBAN}}
    BBAN: {{.BBAN}}
    {{- end}}
    {{- if .BankCode}}
    bank_code: {{.BankCode}}
    {{- end}}
    {{- if .BranchCode}}
    branch_code: {{.BranchCode}}
    {{- end}}
    {{- if .AccountNumber}}
    account_number: {{.AccountNumber}}
    {{- end}}
    sepa: {{.IsSEPA}}
{{end}}
`

type templateData struct {
	SourceFile string
	Countries  []registry.Country
}

func main() {
	var (
		inputPath  = flag.String("input", "", "Path to input TSV file (required)")
		outputPath = flag.String("output", "", "Path to output YAML file (required)")
	)

	flag.Parse()

	if *inputPath == "" || *outputPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*inputPath, *outputPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func writeYAML(w io.Writer, sourceFile string, countries []registry.Country) error {
	tmpl, err := template.New("yaml").Parse(yamlTemplate)
	if err != nil {
		return err
	}

	data := templateData{
		SourceFile: filepath.Base(sourceFile),
		Countries:  countries,
	}
	return tmpl.Execute(w, data)
}

func run(inputPath, outputPath string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer inputFile.Close()

	parser := tsv.NewParser(inputFile)
	countries, err := parser.ParseCountries()
	if err != nil {
		return fmt.Errorf("failed to parse TSV: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Parsed %d countries\n", len(countries))

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	if err := writeYAML(outputFile, inputPath, countries); err != nil {
		return fmt.Errorf("failed to write YAML: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Successfully wrote %s\n", outputPath)
	return nil
}
