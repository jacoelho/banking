package iban

import "testing"

func TestGenerate(t *testing.T) {
	t.Parallel()

	for _, country := range countrySpecs {
		countryCode := country.code
		t.Run(countryCode, func(t *testing.T) {
			t.Parallel()

			generated, err := Generate(countryCode)
			if err != nil {
				t.Fatal(err)
			}
			if err := Validate(generated); err != nil {
				t.Fatal(err)
			}
		})
	}
}
