// +build validation

package iban

import "testing"

func TestGenerate(t *testing.T) {
	tests := []struct {
		name string
		fn   func() string
	}{
		{
			name: "Albania",
			fn:   GenerateAlbaniaIBAN,
		},
		{
			name: "Andorra",
			fn:   GenerateAndorraIBAN,
		},
		{
			name: "Austria",
			fn:   GenerateAustriaIBAN,
		},
		{
			name: "Republic Of Azerbaijan",
			fn:   GenerateRepublicOfAzerbaijanIBAN,
		},
		{
			name: "Kingdom Of Bahrain",
			fn:   GenerateKingdomOfBahrainIBAN,
		},
		{
			name: "Belgium",
			fn:   GenerateBelgiumIBAN,
		},
		{
			name: "Bosnia And Herzegovina",
			fn:   GenerateBosniaAndHerzegovinaIBAN,
		},
		{
			name: "Brazil",
			fn:   GenerateBrazilIBAN,
		},
		{
			name: "Bulgaria",
			fn:   GenerateBulgariaIBAN,
		},
		{
			name: "Costa Rica",
			fn:   GenerateCostaRicaIBAN,
		},
		{
			name: "Croatia",
			fn:   GenerateCroatiaIBAN,
		},
		{
			name: "Cyprus",
			fn:   GenerateCyprusIBAN,
		},
		{
			name: "Czech Republic",
			fn:   GenerateCzechRepublicIBAN,
		},
		{
			name: "Denmark",
			fn:   GenerateDenmarkIBAN,
		},
		{
			name: "Faroe Islands",
			fn:   GenerateFaroeIslandsIBAN,
		},
		{
			name: "Greenland",
			fn:   GenerateGreenlandIBAN,
		},
		{
			name: "Dominican Republic",
			fn:   GenerateDominicanRepublicIBAN,
		},
		{
			name: "Estonia",
			fn:   GenerateEstoniaIBAN,
		},
		{
			name: "Finland",
			fn:   GenerateFinlandIBAN,
		},
		{
			name: "France",
			fn:   GenerateFranceIBAN,
		},
		{
			name: "French Guyana",
			fn:   GenerateFrenchGuyanaIBAN,
		},
		{
			name: "Guadeloupe",
			fn:   GenerateGuadeloupeIBAN,
		},
		{
			name: "Martinique",
			fn:   GenerateMartiniqueIBAN,
		},
		{
			name: "Reunion",
			fn:   GenerateReunionIBAN,
		},
		{
			name: "French Polynesia",
			fn:   GenerateFrenchPolynesiaIBAN,
		},
		{
			name: "French Southern Territories",
			fn:   GenerateFrenchSouthernTerritoriesIBAN,
		},
		{
			name: "Mayotte",
			fn:   GenerateMayotteIBAN,
		},
		{
			name: "New Caledonia",
			fn:   GenerateNewCaledoniaIBAN,
		},
		{
			name: "Saint Barthelemy",
			fn:   GenerateSaintBarthelemyIBAN,
		},
		{
			name: "Saint Martin",
			fn:   GenerateSaintMartinIBAN,
		},
		{
			name: "Saint Pierre Et Miquelon",
			fn:   GenerateSaintPierreEtMiquelonIBAN,
		},
		{
			name: "Wallis And Futuna Islands",
			fn:   GenerateWallisAndFutunaIslandsIBAN,
		},
		{
			name: "Georgia",
			fn:   GenerateGeorgiaIBAN,
		},
		{
			name: "Germany",
			fn:   GenerateGermanyIBAN,
		},
		{
			name: "Gibraltar",
			fn:   GenerateGibraltarIBAN,
		},
		{
			name: "Greece",
			fn:   GenerateGreeceIBAN,
		},
		{
			name: "Guatemala",
			fn:   GenerateGuatemalaIBAN,
		},
		{
			name: "Hungary",
			fn:   GenerateHungaryIBAN,
		},
		{
			name: "Iceland",
			fn:   GenerateIcelandIBAN,
		},
		{
			name: "Ireland",
			fn:   GenerateIrelandIBAN,
		},
		{
			name: "Israel",
			fn:   GenerateIsraelIBAN,
		},
		{
			name: "Italy",
			fn:   GenerateItalyIBAN,
		},
		{
			name: "Jordan",
			fn:   GenerateJordanIBAN,
		},
		{
			name: "Kazakhstan",
			fn:   GenerateKazakhstanIBAN,
		},
		{
			name: "Republic Of Kosovo",
			fn:   GenerateRepublicOfKosovoIBAN,
		},
		{
			name: "Kuwait",
			fn:   GenerateKuwaitIBAN,
		},
		{
			name: "Latvia",
			fn:   GenerateLatviaIBAN,
		},
		{
			name: "Lebanon",
			fn:   GenerateLebanonIBAN,
		},
		{
			name: "Principality Of Liechtenstein",
			fn:   GeneratePrincipalityOfLiechtensteinIBAN,
		},
		{
			name: "Lithuania",
			fn:   GenerateLithuaniaIBAN,
		},
		{
			name: "Luxembourg",
			fn:   GenerateLuxembourgIBAN,
		},
		{
			name: "Macedonia",
			fn:   GenerateMacedoniaIBAN,
		},
		{
			name: "Malta",
			fn:   GenerateMaltaIBAN,
		},
		{
			name: "Mauritania",
			fn:   GenerateMauritaniaIBAN,
		},
		{
			name: "Mauritius",
			fn:   GenerateMauritiusIBAN,
		},
		{
			name: "Moldova",
			fn:   GenerateMoldovaIBAN,
		},
		{
			name: "Monaco",
			fn:   GenerateMonacoIBAN,
		},
		{
			name: "Montenegro",
			fn:   GenerateMontenegroIBAN,
		},
		{
			name: "The Netherlands",
			fn:   GenerateTheNetherlandsIBAN,
		},
		{
			name: "Norway",
			fn:   GenerateNorwayIBAN,
		},
		{
			name: "Pakistan",
			fn:   GeneratePakistanIBAN,
		},
		{
			name: "State Of Palestine",
			fn:   GenerateStateOfPalestineIBAN,
		},
		{
			name: "Poland",
			fn:   GeneratePolandIBAN,
		},
		{
			name: "Portugal",
			fn:   GeneratePortugalIBAN,
		},
		{
			name: "Qatar",
			fn:   GenerateQatarIBAN,
		},
		{
			name: "Romania",
			fn:   GenerateRomaniaIBAN,
		},
		{
			name: "Saint Lucia",
			fn:   GenerateSaintLuciaIBAN,
		},
		{
			name: "San Marino",
			fn:   GenerateSanMarinoIBAN,
		},
		{
			name: "Sao Tome And Principe",
			fn:   GenerateSaoTomeAndPrincipeIBAN,
		},
		{
			name: "Saudi Arabia",
			fn:   GenerateSaudiArabiaIBAN,
		},
		{
			name: "Serbia",
			fn:   GenerateSerbiaIBAN,
		},
		{
			name: "Slovak Republic",
			fn:   GenerateSlovakRepublicIBAN,
		},
		{
			name: "Slovenia",
			fn:   GenerateSloveniaIBAN,
		},
		{
			name: "Spain",
			fn:   GenerateSpainIBAN,
		},
		{
			name: "Sweden",
			fn:   GenerateSwedenIBAN,
		},
		{
			name: "Switzerland",
			fn:   GenerateSwitzerlandIBAN,
		},
		{
			name: "Timor Leste",
			fn:   GenerateTimorLesteIBAN,
		},
		{
			name: "Tunisia",
			fn:   GenerateTunisiaIBAN,
		},
		{
			name: "Turkey",
			fn:   GenerateTurkeyIBAN,
		},
		{
			name: "United Arab Emirates",
			fn:   GenerateUnitedArabEmiratesIBAN,
		},
		{
			name: "United Kingdom",
			fn:   GenerateUnitedKingdomIBAN,
		},
		{
			name: "British Virgin Islands",
			fn:   GenerateBritishVirginIslandsIBAN,
		},
		{
			name: "Seychelles",
			fn:   GenerateSeychellesIBAN,
		},
		{
			name: "Ukraine",
			fn:   GenerateUkraineIBAN,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Validate(tt.fn()); err != nil {
				t.Error(err)
			}
		})
	}
}
