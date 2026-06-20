package iban

// IsSEPACountryCode reports whether a country code is a SEPA member.
func IsSEPACountryCode(countryCode string) (bool, error) {
	if !validCountryCode(countryCode) {
		return false, invalidCountryCode(countryCode)
	}
	slot, _ := countryIndexSlot(countryCode)
	entry := countryCodeIndex[slot]
	if entry&countryIndexMask == 0 {
		return false, unsupportedCountry(countryCode)
	}
	return entry&countrySEPAFlag != 0, nil
}

// IsSEPA reports whether a valid IBAN belongs to a SEPA member country.
func IsSEPA(iban string) (bool, error) {
	if err := Validate(iban); err != nil {
		return false, err
	}
	return IsSEPACountryCode(iban[0:2])
}
