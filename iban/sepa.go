package iban

// IsSEPACountryCode reports whether a country code is a SEPA member.
func IsSEPACountryCode(countryCode string) (bool, error) {
	if len(countryCode) != 2 {
		return false, &ErrValidationLength{Expected: 2, Actual: len(countryCode)}
	}
	slot, ok := countryIndexSlot(countryCode)
	if !ok {
		return false, &ErrUnsupportedCountry{CountryCode: countryCode}
	}
	entry := countryCodeIndex[slot]
	if entry&countryIndexMask == 0 {
		return false, &ErrUnsupportedCountry{CountryCode: countryCode}
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
