package ibanregistry

import (
	"encoding/csv"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

type rowLabel string

const (
	notAvailable                = "N/A"
	rowCountryName     rowLabel = "Name of country"
	rowCountryCode     rowLabel = "IBAN prefix country code (ISO 3166)"
	rowCountryIncludes rowLabel = "Country code includes other countries/territories"
	rowSEPA            rowLabel = "SEPA country"
	rowSEPAIncludes    rowLabel = "SEPA country also includes"
	rowBBANStructure   rowLabel = "BBAN structure"
	rowBankPosition    rowLabel = "Bank identifier position within the BBAN"
	rowBranchPosition  rowLabel = "Branch identifier position within the BBAN"
	rowIBANStructure   rowLabel = "IBAN structure"
)

var requiredRows = []rowLabel{
	rowCountryName,
	rowCountryCode,
	rowCountryIncludes,
	rowSEPA,
	rowSEPAIncludes,
	rowBBANStructure,
	rowBankPosition,
	rowBranchPosition,
	rowIBANStructure,
}

var territoryNames = map[string]string{
	"AX": "Aland Islands",
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
	"IM": "Isle Of Man",
	"JE": "Jersey",
	"GG": "Guernsey",
}

// ParseTSV parses a decoded SWIFT registry TSV into a validated registry model.
func ParseTSV(source string) (Registry, error) {
	reader := csv.NewReader(strings.NewReader(source))
	reader.Comma = '\t'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1

	rowMap, err := readRequiredRows(reader)
	if err != nil {
		return Registry{}, err
	}
	if widthErr := validateRowWidths(rowMap); widthErr != nil {
		return Registry{}, widthErr
	}

	registry, territoryByParent, err := parsePrimaryCountries(rowMap)
	if err != nil {
		return Registry{}, err
	}
	if err := expandTerritories(&registry, territoryByParent); err != nil {
		return Registry{}, err
	}
	if err := applySEPAIncludes(&registry, rowMap, territoryByParent); err != nil {
		return Registry{}, err
	}
	return registry, nil
}

func readRequiredRows(reader *csv.Reader) (map[rowLabel][]string, error) {
	rows := make(map[rowLabel][]string, len(requiredRows))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("reading TSV: %w", err)
		}
		if len(record) == 0 {
			continue
		}
		label := rowLabel(strings.TrimSpace(record[0]))
		if !slices.Contains(requiredRows, label) {
			continue
		}
		if _, exists := rows[label]; exists {
			return nil, fmt.Errorf("duplicate required row %q", label)
		}
		rows[label] = record
	}

	for _, label := range requiredRows {
		if _, ok := rows[label]; !ok {
			return nil, fmt.Errorf("missing required row %q", label)
		}
	}
	return rows, nil
}

func validateRowWidths(rows map[rowLabel][]string) error {
	width := len(rows[rowCountryCode])
	if width < 2 {
		return fmt.Errorf("country code row has no country columns")
	}
	for _, label := range requiredRows {
		if got := len(rows[label]); got != width {
			return fmt.Errorf("row %q has %d columns, want %d", label, got, width)
		}
	}
	return nil
}

func parsePrimaryCountries(rows map[rowLabel][]string) (Registry, map[string][]string, error) {
	codeRow := rows[rowCountryCode]
	registry := Registry{
		Countries: make([]CountrySpec, 0, len(codeRow)-1),
	}
	territoryByParent := make(map[string][]string)
	seenCodes := make(map[string]bool)

	for col := 1; col < len(codeRow); col++ {
		code := getColumn(codeRow, col)
		if err := validateCode(code); err != nil {
			return Registry{}, nil, fmt.Errorf("column %d country code: %w", col, err)
		}
		if seenCodes[code] {
			return Registry{}, nil, fmt.Errorf("duplicate country code %q", code)
		}
		seenCodes[code] = true

		country, err := parseCountry(rows, col, code, normaliseCountryName(getColumn(rows[rowCountryName], col)))
		if err != nil {
			return Registry{}, nil, fmt.Errorf("country %s: %w", code, err)
		}

		sepa, err := parseBoolean(getColumn(rows[rowSEPA], col))
		if err != nil {
			return Registry{}, nil, fmt.Errorf("country %s SEPA: %w", code, err)
		}
		country.IsSEPA = sepa
		registry.Countries = append(registry.Countries, country)

		territories, err := parseTerritoryCodes(getColumn(rows[rowCountryIncludes], col))
		if err != nil {
			return Registry{}, nil, fmt.Errorf("country %s territories: %w", code, err)
		}
		territoryByParent[code] = territories
	}

	return registry, territoryByParent, nil
}

func parseCountry(rows map[rowLabel][]string, col int, code, name string) (CountrySpec, error) {
	ibanRaw := getColumn(rows[rowIBANStructure], col)
	bbanRaw := getColumn(rows[rowBBANStructure], col)

	iban, err := ParseStructure(ibanRaw)
	if err != nil {
		return CountrySpec{}, fmt.Errorf("IBAN structure %q: %w", ibanRaw, err)
	}
	bban, err := ParseStructure(bbanRaw)
	if err != nil {
		return CountrySpec{}, fmt.Errorf("BBAN structure %q: %w", bbanRaw, err)
	}
	if prefixErr := validateIBANPrefix(code, iban); prefixErr != nil {
		return CountrySpec{}, prefixErr
	}
	ibanBBAN, err := ibanBBANStructure(iban)
	if err != nil {
		return CountrySpec{}, err
	}
	ibanBBAN.Rules = consolidateRules(ibanBBAN.Rules)
	bban.Rules = consolidateRules(bban.Rules)
	if !structuresEqual(ibanBBAN, bban) {
		return CountrySpec{}, fmt.Errorf("IBAN BBAN suffix does not match BBAN structure")
	}
	iban.Rules = consolidateIBANRules(iban.Rules)

	bank, err := parseOptionalSpan(getColumn(rows[rowBankPosition], col), bban.Length)
	if err != nil {
		return CountrySpec{}, fmt.Errorf("bank position: %w", err)
	}
	branch, err := parseOptionalSpan(getColumn(rows[rowBranchPosition], col), bban.Length)
	if err != nil {
		return CountrySpec{}, fmt.Errorf("branch position: %w", err)
	}

	accountStart := 0
	if bank.Present {
		accountStart = max(accountStart, bank.Span.End)
	}
	if branch.Present {
		accountStart = max(accountStart, branch.Span.End)
	}
	account := Span{Start: accountStart, End: bban.Length}
	if err := account.Validate(bban.Length); err != nil {
		return CountrySpec{}, fmt.Errorf("account position: %w", err)
	}

	return CountrySpec{
		Code:          code,
		Name:          name,
		IBAN:          iban,
		BBAN:          bban,
		BankCode:      bank,
		BranchCode:    branch,
		AccountNumber: account,
	}, nil
}

func validateIBANPrefix(code string, iban Structure) error {
	if len(iban.Rules) < 2 {
		return fmt.Errorf("IBAN structure has too few rules")
	}
	static, ok := iban.Rules[0].(StaticRule)
	if !ok || static.StartPosition != 0 || static.Value != code {
		return fmt.Errorf("IBAN structure must start with %s", code)
	}
	check, ok := iban.Rules[1].(RangeRule)
	if !ok || check.StartPosition != 2 || check.Length != 2 || check.Format != Digit {
		return fmt.Errorf("IBAN structure must contain 2 numeric check digits after country code")
	}
	return nil
}

func expandTerritories(registry *Registry, territoryByParent map[string][]string) error {
	seen := make(map[string]bool, len(registry.Countries))
	for _, country := range registry.Countries {
		seen[country.Code] = true
	}

	for _, parent := range registry.Countries {
		for _, code := range territoryByParent[parent.Code] {
			if seen[code] {
				return fmt.Errorf("duplicate expanded country code %q", code)
			}
			name, ok := territoryNames[code]
			if !ok {
				return fmt.Errorf("unknown territory code %q", code)
			}
			territory := parent
			territory.Code = code
			territory.Name = name
			territory.IsSEPA = false
			territory.IBAN = replaceIBANCode(parent.IBAN, code)
			registry.Countries = append(registry.Countries, territory)
			seen[code] = true
		}
	}
	return nil
}

func replaceIBANCode(iban Structure, code string) Structure {
	rules := make([]Rule, len(iban.Rules))
	copy(rules, iban.Rules)
	if len(rules) > 0 {
		if static, ok := rules[0].(StaticRule); ok {
			static.Value = code
			rules[0] = static
		}
	}
	return Structure{Rules: rules, Length: iban.Length}
}

func applySEPAIncludes(registry *Registry, rows map[rowLabel][]string, territoryByParent map[string][]string) error {
	countryIndex := make(map[string]int, len(registry.Countries))
	for i, country := range registry.Countries {
		countryIndex[country.Code] = i
	}
	includedBy := make(map[string]string)

	sepaRow := rows[rowSEPAIncludes]
	codeRow := rows[rowCountryCode]
	for col := 1; col < len(codeRow); col++ {
		parentCode := getColumn(codeRow, col)
		parentSEPA := registry.Countries[countryIndex[parentCode]].IsSEPA
		codes, err := parseSEPAIncludeCodes(getColumn(sepaRow, col), parentCode, parentSEPA, countryIndex, territoryByParent[parentCode])
		if err != nil {
			return fmt.Errorf("country %s SEPA includes: %w", parentCode, err)
		}
		for _, code := range codes {
			if previous, ok := includedBy[code]; ok && previous != parentCode {
				return fmt.Errorf("SEPA code %q included by both %s and %s", code, previous, parentCode)
			}
			includedBy[code] = parentCode
			registry.Countries[countryIndex[code]].IsSEPA = true
		}
	}
	return nil
}

func getColumn(row []string, col int) string {
	if col < 0 || col >= len(row) {
		return ""
	}
	return strings.TrimSpace(row[col])
}

func validateCode(code string) error {
	if len(code) != 2 {
		return fmt.Errorf("must be 2 characters, got %d", len(code))
	}
	for i := range 2 {
		if code[i] < 'A' || code[i] > 'Z' {
			return fmt.Errorf("must be uppercase ASCII, got %q", code)
		}
	}
	return nil
}

func parseBoolean(value string) (bool, error) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "yes":
		return true, nil
	case "no":
		return false, nil
	default:
		return false, fmt.Errorf("expected Yes or No, got %q", value)
	}
}

func normaliseCountryName(name string) string {
	name = strings.Map(func(r rune) rune {
		switch r {
		case '"':
			return -1
		case '-':
			return ' '
		default:
			return r
		}
	}, name)
	if before, _, ok := strings.Cut(name, ","); ok {
		name = before
	}
	if before, _, ok := strings.Cut(name, "("); ok {
		name = before
	}
	return strings.TrimSpace(name)
}

func parseOptionalSpan(value string, limit int) (OptionalSpan, error) {
	value = strings.TrimSpace(value)
	if value == "" || value == notAvailable {
		return OptionalSpan{}, nil
	}
	span, err := parseSpan(value, limit)
	if err != nil {
		return OptionalSpan{}, err
	}
	return OptionalSpan{Span: span, Present: true}, nil
}

func parseSpan(value string, limit int) (Span, error) {
	left, right, ok := strings.Cut(strings.TrimSpace(value), "-")
	if !ok {
		return Span{}, fmt.Errorf("invalid position %q", value)
	}
	start, err := strconv.Atoi(strings.TrimSpace(left))
	if err != nil {
		return Span{}, fmt.Errorf("invalid start: %w", err)
	}
	end, err := strconv.Atoi(strings.TrimSpace(right))
	if err != nil {
		return Span{}, fmt.Errorf("invalid end: %w", err)
	}
	if start < 1 {
		return Span{}, fmt.Errorf("start must be >= 1, got %d", start)
	}
	if end < start {
		return Span{}, fmt.Errorf("end %d is before start %d", end, start)
	}
	span := Span{Start: start - 1, End: end}
	if err := span.Validate(limit); err != nil {
		return Span{}, err
	}
	return span, nil
}

func parseTerritoryCodes(value string) ([]string, error) {
	value = strings.TrimSpace(strings.Trim(value, `"`))
	if value == "" || value == notAvailable {
		return nil, nil
	}
	var codes []string
	seen := make(map[string]bool)
	for token := range strings.SplitSeq(value, ",") {
		code := stripAnnotation(token)
		if err := validateCode(code); err != nil {
			return nil, fmt.Errorf("malformed territory token %q: %w", token, err)
		}
		if _, ok := territoryNames[code]; !ok {
			return nil, fmt.Errorf("unknown territory code %q", code)
		}
		if seen[code] {
			return nil, fmt.Errorf("duplicate territory code %q", code)
		}
		seen[code] = true
		codes = append(codes, code)
	}
	return codes, nil
}

var allowedSEPAAnnotations = map[string]map[string]bool{
	"PT": {
		"Azores":  true,
		"Madeira": true,
	},
}

func parseSEPAIncludeCodes(value, parentCode string, parentSEPA bool, supported map[string]int, ownedTerritories []string) ([]string, error) {
	value = strings.TrimSpace(strings.Trim(value, `"`))
	if value == "" || value == notAvailable {
		if parentSEPA {
			return slices.Clone(ownedTerritories), nil
		}
		return nil, nil
	}
	owned := make(map[string]bool, len(ownedTerritories))
	for _, code := range ownedTerritories {
		owned[code] = true
	}
	var codes []string
	seen := make(map[string]bool)
	for token := range strings.SplitSeq(value, ",") {
		token = strings.TrimSpace(token)
		code := stripAnnotation(token)
		if allowedSEPAAnnotations[parentCode][code] {
			continue
		}
		if err := validateCode(code); err != nil {
			return nil, fmt.Errorf("malformed SEPA token %q: %w", token, err)
		}
		if _, ok := supported[code]; !ok {
			return nil, fmt.Errorf("SEPA code %q is not a supported IBAN code", code)
		}
		if !owned[code] {
			return nil, fmt.Errorf("SEPA code %q is not owned by this country", code)
		}
		if seen[code] {
			return nil, fmt.Errorf("duplicate SEPA code %q", code)
		}
		seen[code] = true
		codes = append(codes, code)
	}
	return codes, nil
}

func stripAnnotation(token string) string {
	token = strings.TrimSpace(token)
	if before, _, ok := strings.Cut(token, "("); ok {
		token = before
	}
	return strings.TrimSpace(token)
}
