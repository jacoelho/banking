package iban

// GetBBAN retrieves BBAN from a valid IBAN.
func GetBBAN(iban string) (BBAN, error) {
	if err := Validate(iban); err != nil {
		return BBAN{}, err
	}
	country, _ := lookupCountry(iban[0:2])
	return country.bban(iban), nil
}
