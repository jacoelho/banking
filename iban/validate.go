package iban

// validateIBANStructure validates an IBAN without comparing checksum.
func validateIBANStructure(iban string) error {
	if len(iban) < 2 {
		return invalidIBANLength(2, len(iban))
	}
	code := iban[0:2]
	if !validCountryCode(code) {
		return invalidIBANCharacters(0, 2, CharClassUpperAlpha, code)
	}
	country, ok := lookupCountry(code)
	if !ok {
		return unsupportedIBANCountry(code)
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
