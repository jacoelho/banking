package tsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jacoelho/banking/registry"
	"github.com/jacoelho/banking/registry/parser"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// rowLabel defines the expected labels for TSV rows
type rowLabel string

const (
	rowCountryName     rowLabel = "Name of country"
	rowCountryCode     rowLabel = "IBAN prefix country code (ISO 3166)"
	rowCountryIncludes rowLabel = "Country code includes other countries/territories"
	rowSEPA            rowLabel = "SEPA country"
	rowBBANStructure   rowLabel = "BBAN structure"
	rowBankPosition    rowLabel = "Bank identifier position within the BBAN"
	rowBranchPosition  rowLabel = "Branch identifier position within the BBAN"
	rowIBANStructure   rowLabel = "IBAN structure"
)

func requiredLabels() []rowLabel {
	return []rowLabel{
		rowCountryName,
		rowCountryCode,
		rowCountryIncludes,
		rowSEPA,
		rowBBANStructure,
		rowBankPosition,
		rowBranchPosition,
		rowIBANStructure,
	}
}

var territoryNames = map[string]string{
	"GF": "French Guyana",
	"GP": "Guadeloupe",
	"MQ": "Martinique",
	"RE": "Reunion",
	"PF": "French Polynesia",
	"TF": "French Southern Territories",
	"YT": "Mayotte",
	"NC": "New Caledonia",
	"BL": "Saint Barthelemy",
	"MF": "Saint Martin",
	"PM": "Saint Pierre Et Miquelon",
	"WF": "Wallis And Futuna Islands",
	"AX": "Aland Islands",
	"IM": "Isle Of Man",
	"JE": "Jersey",
	"GG": "Guernsey",
}

// Parser parses TSV IBAN registry files.
type Parser struct {
	reader *csv.Reader
}

// NewParser returns a Parser that reads Latin1-encoded TSV data.
func NewParser(r io.Reader) *Parser {
	decoder := charmap.ISO8859_1.NewDecoder()
	reader := transform.NewReader(r, decoder)

	csvReader := csv.NewReader(reader)
	csvReader.Comma = '\t'
	csvReader.LazyQuotes = true
	csvReader.FieldsPerRecord = -1

	return &Parser{reader: csvReader}
}

// ParseCountries parses the TSV file and returns a slice of Country entries.
func (p *Parser) ParseCountries() ([]registry.Country, error) {
	rowMap := make(map[rowLabel][]string)
	expectedLabels := requiredLabels()

	for {
		record, err := p.reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read TSV: %w", err)
		}

		if len(record) == 0 {
			continue
		}

		label := strings.TrimSpace(record[0])
		for _, expectedLabel := range expectedLabels {
			if label == string(expectedLabel) {
				rowMap[expectedLabel] = record
				break
			}
		}
	}

	for _, label := range expectedLabels {
		if _, exists := rowMap[label]; !exists {
			return nil, fmt.Errorf("missing required row: %q", label)
		}
	}

	countryCodes := rowMap[rowCountryCode]
	countries := make([]registry.Country, 0, len(countryCodes)-1)

	for i := 1; i < len(countryCodes); i++ {
		code := getColumn(countryCodes, i)
		if code == "" {
			return nil, fmt.Errorf("empty country code at column %d", i)
		}

		country := registry.Country{
			Code:   code,
			Name:   normaliseCountryName(getColumn(rowMap[rowCountryName], i)),
			IBAN:   getColumn(rowMap[rowIBANStructure], i),
			IsSEPA: parseBoolean(getColumn(rowMap[rowSEPA], i)),
		}

		if bban := getColumn(rowMap[rowBBANStructure], i); bban != "" {
			country.BBAN = bban
		}

		var bankEnd, branchEnd int

		if bankPos := getColumn(rowMap[rowBankPosition], i); bankPos != "" && bankPos != "N/A" {
			start, end, err := convertPosition(bankPos)
			if err != nil {
				return nil, fmt.Errorf("country %s: invalid bank position %q: %w", code, bankPos, err)
			}
			bankEnd = end
			country.BankCode = fmt.Sprintf("%d:%d", start, end)
		}

		if branchPos := getColumn(rowMap[rowBranchPosition], i); branchPos != "" && branchPos != "N/A" {
			start, end, err := convertPosition(branchPos)
			if err != nil {
				return nil, fmt.Errorf("country %s: invalid branch position %q: %w", code, branchPos, err)
			}
			branchEnd = end
			country.BranchCode = fmt.Sprintf("%d:%d", start, end)
		}

		accountPos, err := deriveAccountPosition(&country, bankEnd, branchEnd)
		if err != nil {
			return nil, fmt.Errorf("country %s: %w", country.Code, err)
		}
		if accountPos != "" {
			country.AccountNumber = accountPos
		}

		countries = append(countries, country)
	}

	return expandTerritories(countries, rowMap)
}

func getColumn(row []string, index int) string {
	if index < 0 || index >= len(row) {
		return ""
	}
	return strings.TrimSpace(row[index])
}

// parseBoolean reports whether value is "yes".
func parseBoolean(value string) bool {
	return strings.EqualFold(strings.TrimSpace(value), "yes")
}

// normaliseCountryName removes qualifiers and formatting from country names.
func normaliseCountryName(name string) string {
	name = strings.Map(func(r rune) rune {
		if r == '"' {
			return -1
		}
		if r == '-' {
			return ' '
		}
		return r
	}, name)

	if idx := strings.Index(name, ","); idx != -1 {
		name = name[:idx]
	}

	if idx := strings.Index(name, "("); idx != -1 {
		name = name[:idx]
	}

	return strings.TrimSpace(name)
}

// convertPosition converts TSV position notation to 0-based start/end indices.
func convertPosition(tsvPos string) (start, end int, err error) {
	tsvPos = strings.TrimSpace(tsvPos)
	if tsvPos == "" || tsvPos == "N/A" {
		return 0, 0, fmt.Errorf("empty position")
	}

	parts := strings.Split(tsvPos, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid position format: %s", tsvPos)
	}

	startPos, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid start position: %w", err)
	}

	endPos, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid end position: %w", err)
	}

	return startPos - 1, endPos, nil
}

// deriveAccountPosition returns the account number position in BBAN coordinates.
func deriveAccountPosition(country *registry.Country, bankEnd, branchEnd int) (string, error) {
	if country.IBAN == "" {
		return "", fmt.Errorf("cannot derive account position: IBAN is empty")
	}

	parseResult, err := parser.Parse(country.IBAN)
	if err != nil {
		return "", fmt.Errorf("failed to parse IBAN structure %q: %w", country.IBAN, err)
	}

	bbanLen := parseResult.Length - 4
	accountStart := max(branchEnd, bankEnd)

	return fmt.Sprintf("%d:%d", accountStart, bbanLen), nil
}

// expandTerritories parses territory codes and creates country entries for each.
func expandTerritories(countries []registry.Country, rowMap map[rowLabel][]string) ([]registry.Country, error) {
	expanded := make([]registry.Country, 0, len(countries))
	expanded = append(expanded, countries...)

	for i, country := range countries {
		includesRow := rowMap[rowCountryIncludes]
		if len(includesRow) <= i+1 {
			continue
		}

		includesValue := getColumn(includesRow, i+1)
		if includesValue == "" || includesValue == "N/A" {
			continue
		}

		territoryCodes := parseTerritories(includesValue)
		for _, code := range territoryCodes {
			name, exists := territoryNames[code]
			if !exists {
				return nil, fmt.Errorf("unknown territory code %q for country %s - update territoryNames map", code, country.Code)
			}

			territory := registry.Country{
				Code:          code,
				Name:          name,
				IBAN:          replaceCountryCode(country.IBAN, code),
				BBAN:          country.BBAN,
				BankCode:      country.BankCode,
				BranchCode:    country.BranchCode,
				AccountNumber: country.AccountNumber,
				IsSEPA:        country.IsSEPA,
			}

			expanded = append(expanded, territory)
		}
	}

	return expanded, nil
}

// parseTerritories extracts territory codes from comma-separated includes field.
func parseTerritories(value string) []string {
	value = strings.Trim(value, "\"")
	parts := strings.Split(value, ",")

	codes := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if idx := strings.Index(part, "("); idx != -1 {
			part = part[:idx]
			part = strings.TrimSpace(part)
		}
		if part != "" {
			codes = append(codes, part)
		}
	}

	return codes
}

// replaceCountryCode replaces the country code prefix in an IBAN structure.
func replaceCountryCode(iban, newCode string) string {
	if len(iban) < 2 {
		return iban
	}
	return newCode + iban[2:]
}
