package iban

// Generate generates an IBAN based on ISO 3166-1 country code.
func Generate(countryCode string) (string, error) {
	if !validCountryCode(countryCode) {
		return "", invalidCountryCode(countryCode)
	}
	country, ok := lookupCountry(countryCode)
	if !ok {
		return "", unsupportedCountry(countryCode)
	}
	return country.generate(), nil
}
