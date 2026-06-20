package iban

// Generate generates an IBAN based on ISO 3166-1 country code.
func Generate(countryCode string) (string, error) {
	if len(countryCode) != 2 {
		return "", &ErrValidationLength{Expected: 2, Actual: len(countryCode)}
	}
	country, ok := lookupCountry(countryCode)
	if !ok {
		return "", &ErrUnsupportedCountry{CountryCode: countryCode}
	}
	return country.generate(), nil
}
