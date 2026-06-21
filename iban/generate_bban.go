package iban

const (
	bbanPartBankCode      = "BankCode"
	bbanPartBranchCode    = "BranchCode"
	bbanPartAccountNumber = "AccountNumber"
)

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

func (c *countrySpec) validateBBANParts(countryCode string, parts BBANParts) error {
	if err := c.validateBBANPart(countryCode, bbanPartBankCode, parts.BankCode, c.bankCode); err != nil {
		return err
	}
	if err := c.validateBBANPart(countryCode, bbanPartBranchCode, parts.BranchCode, c.branchCode); err != nil {
		return err
	}
	return c.validateBBANPart(countryCode, bbanPartAccountNumber, parts.AccountNumber, c.accountNumber)
}

func (c *countrySpec) validateBBANPart(countryCode, field, value string, span bbanComponent) error {
	if value == "" {
		return nil
	}

	expectedLength := span.length()
	if len(value) != expectedLength {
		return invalidBBANParts(bbanPartsErrorData{
			countryCode:    countryCode,
			field:          field,
			length:         len(value),
			expectedLength: expectedLength,
			actualLength:   len(value),
			actual:         value,
		})
	}

	start, end := span.ibanBounds()
	for _, rule := range c.rules {
		ruleStart := int(rule.start)
		ruleEnd := ruleStart + int(rule.length)
		overlapStart := max(start, ruleStart)
		overlapEnd := min(end, ruleEnd)
		if overlapStart >= overlapEnd {
			continue
		}

		subject := value[overlapStart-start : overlapEnd-start]
		violation, ok := rule.validate(subject, overlapStart-ruleStart)
		if ok {
			continue
		}
		return invalidBBANParts(bbanPartsErrorData{
			countryCode:    countryCode,
			field:          field,
			position:       overlapStart - start + violation.offset,
			length:         violation.length,
			expectedLength: expectedLength,
			actualLength:   len(value),
			expected:       violation.expected,
			expectedValue:  violation.expectedValue,
			actual:         violation.actual,
		})
	}
	return nil
}

func (c *countrySpec) generateWithBBANParts(parts BBANParts) string {
	iban := []byte(c.generate())
	applyBBANPart(iban, parts.BankCode, c.bankCode)
	applyBBANPart(iban, parts.BranchCode, c.branchCode)
	applyBBANPart(iban, parts.AccountNumber, c.accountNumber)
	return replaceChecksumBytes(iban)
}

func applyBBANPart(iban []byte, value string, span bbanComponent) {
	if value == "" {
		return
	}
	start, end := span.ibanBounds()
	copy(iban[start:end], value)
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
