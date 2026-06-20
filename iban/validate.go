package iban

// validateIBANStructure validates an IBAN without comparing checksum.
func validateIBANStructure(iban string) error {
	if len(iban) < 2 {
		return &ErrValidationLength{Expected: 2, Actual: len(iban)}
	}
	code := iban[0:2]
	country, ok := lookupCountry(code)
	if !ok {
		return &ErrUnsupportedCountry{CountryCode: code}
	}
	return country.validate(iban)
}

// Validate validates an IBAN.
func Validate(iban string) error {
	if err := validateIBANStructure(iban); err != nil {
		return err
	}
	return validateChecksum(iban)
}
