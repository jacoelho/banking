package iban

const (
	bbanPartBankCode      = "BankCode"
	bbanPartBranchCode    = "BranchCode"
	bbanPartAccountNumber = "AccountNumber"
)

type bbanPart struct {
	field     string
	value     string
	component bbanComponent
}

// GenerateWithBBAN generates an IBAN based on ISO 3166-1 country code and
// optional BBAN field constraints.
func GenerateWithBBAN(countryCode string, parts BBANParts) (string, error) {
	if !validCountryCode(countryCode) {
		return "", invalidCountryCode(countryCode)
	}
	country, ok := lookupCountry(countryCode)
	if !ok {
		return "", unsupportedCountry(countryCode)
	}
	if err := country.validateBBANParts(countryCode, parts); err != nil {
		return "", err
	}
	return country.generateWithBBANParts(parts), nil
}

func (c *countrySpec) bbanParts(parts BBANParts) [3]bbanPart {
	return [3]bbanPart{
		{field: bbanPartBankCode, value: parts.BankCode, component: c.bankCode},
		{field: bbanPartBranchCode, value: parts.BranchCode, component: c.branchCode},
		{field: bbanPartAccountNumber, value: parts.AccountNumber, component: c.accountNumber},
	}
}

func (c *countrySpec) validateBBANParts(countryCode string, parts BBANParts) error {
	for _, part := range c.bbanParts(parts) {
		if err := c.validateBBANPart(countryCode, part); err != nil {
			return err
		}
	}
	return nil
}

func (c *countrySpec) validateBBANPart(countryCode string, part bbanPart) error {
	if part.value == "" {
		return nil
	}

	expectedLength := part.component.length()
	if len(part.value) != expectedLength {
		return invalidBBANParts(bbanPartsErrorData{
			countryCode:    countryCode,
			field:          part.field,
			length:         len(part.value),
			expectedLength: expectedLength,
			actualLength:   len(part.value),
			actual:         part.value,
		})
	}

	start, end := part.component.ibanBounds()
	for _, rule := range c.rules {
		ruleStart := int(rule.start)
		ruleEnd := ruleStart + int(rule.length)
		overlapStart := max(start, ruleStart)
		overlapEnd := min(end, ruleEnd)
		if overlapStart >= overlapEnd {
			continue
		}

		subject := part.value[overlapStart-start : overlapEnd-start]
		violation, ok := rule.validate(subject, overlapStart-ruleStart)
		if ok {
			continue
		}
		return invalidBBANParts(bbanPartsErrorData{
			countryCode:    countryCode,
			field:          part.field,
			position:       overlapStart - start + violation.offset,
			length:         violation.length,
			expectedLength: expectedLength,
			actualLength:   len(part.value),
			expected:       violation.expected,
			expectedValue:  violation.expectedValue,
			actual:         violation.actual,
		})
	}
	return nil
}

func (c *countrySpec) generateWithBBANParts(parts BBANParts) string {
	iban := []byte(c.generate())
	for _, part := range c.bbanParts(parts) {
		applyBBANPart(iban, part)
	}
	return replaceChecksumBytes(iban)
}

func applyBBANPart(iban []byte, part bbanPart) {
	if part.value == "" {
		return
	}
	start, end := part.component.ibanBounds()
	copy(iban[start:end], part.value)
}

func (c bbanComponent) length() int {
	if !c.present() {
		return 0
	}
	return int(c.end - c.start)
}

func (c bbanComponent) ibanBounds() (int, int) {
	start := int(c.start) + ibanBBANOffset
	end := int(c.end) + ibanBBANOffset
	return start, end
}
