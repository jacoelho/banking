package iban

import "testing"

func TestGenerate(t *testing.T) {
	tests := []struct {
		name string
		fn   func() (string, error)
	}{
		{
			name: "Albania",
			fn:   generateAlbaniaIBAN,
		},
		{
			name: "Andorra",
			fn:   generateAndorraIBAN,
		},
		{
			name: "Austria",
			fn:   generateAustriaIBAN,
		},
		{
			name: "Republic Of Azerbaijan",
			fn:   generateRepublicOfAzerbaijanIBAN,
		},
		{
			name: "Kingdom Of Bahrain",
			fn:   generateKingdomOfBahrainIBAN,
		},
		{
			name: "Belgium",
			fn:   generateBelgiumIBAN,
		},
		{
			name: "Bosnia And Herzegovina",
			fn:   generateBosniaAndHerzegovinaIBAN,
		},
		{
			name: "Brazil",
			fn:   generateBrazilIBAN,
		},
		{
			name: "Bulgaria",
			fn:   generateBulgariaIBAN,
		},
		{
			name: "Costa Rica",
			fn:   generateCostaRicaIBAN,
		},
		{
			name: "Croatia",
			fn:   generateCroatiaIBAN,
		},
		{
			name: "Cyprus",
			fn:   generateCyprusIBAN,
		},
		{
			name: "Czech Republic",
			fn:   generateCzechRepublicIBAN,
		},
		{
			name: "Denmark",
			fn:   generateDenmarkIBAN,
		},
		{
			name: "Faroe Islands",
			fn:   generateFaroeIslandsIBAN,
		},
		{
			name: "Greenland",
			fn:   generateGreenlandIBAN,
		},
		{
			name: "Dominican Republic",
			fn:   generateDominicanRepublicIBAN,
		},
		{
			name: "Estonia",
			fn:   generateEstoniaIBAN,
		},
		{
			name: "Finland",
			fn:   generateFinlandIBAN,
		},
		{
			name: "France",
			fn:   generateFranceIBAN,
		},
		{
			name: "French Guyana",
			fn:   generateFrenchGuyanaIBAN,
		},
		{
			name: "Guadeloupe",
			fn:   generateGuadeloupeIBAN,
		},
		{
			name: "Martinique",
			fn:   generateMartiniqueIBAN,
		},
		{
			name: "Reunion",
			fn:   generateReunionIBAN,
		},
		{
			name: "French Polynesia",
			fn:   generateFrenchPolynesiaIBAN,
		},
		{
			name: "French Southern Territories",
			fn:   generateFrenchSouthernTerritoriesIBAN,
		},
		{
			name: "Mayotte",
			fn:   generateMayotteIBAN,
		},
		{
			name: "New Caledonia",
			fn:   generateNewCaledoniaIBAN,
		},
		{
			name: "Saint Barthelemy",
			fn:   generateSaintBarthelemyIBAN,
		},
		{
			name: "Saint Martin",
			fn:   generateSaintMartinIBAN,
		},
		{
			name: "Saint Pierre Et Miquelon",
			fn:   generateSaintPierreEtMiquelonIBAN,
		},
		{
			name: "Wallis And Futuna Islands",
			fn:   generateWallisAndFutunaIslandsIBAN,
		},
		{
			name: "Georgia",
			fn:   generateGeorgiaIBAN,
		},
		{
			name: "Germany",
			fn:   generateGermanyIBAN,
		},
		{
			name: "Gibraltar",
			fn:   generateGibraltarIBAN,
		},
		{
			name: "Greece",
			fn:   generateGreeceIBAN,
		},
		{
			name: "Guatemala",
			fn:   generateGuatemalaIBAN,
		},
		{
			name: "Hungary",
			fn:   generateHungaryIBAN,
		},
		{
			name: "Iceland",
			fn:   generateIcelandIBAN,
		},
		{
			name: "Ireland",
			fn:   generateIrelandIBAN,
		},
		{
			name: "Israel",
			fn:   generateIsraelIBAN,
		},
		{
			name: "Italy",
			fn:   generateItalyIBAN,
		},
		{
			name: "Jordan",
			fn:   generateJordanIBAN,
		},
		{
			name: "Kazakhstan",
			fn:   generateKazakhstanIBAN,
		},
		{
			name: "Republic Of Kosovo",
			fn:   generateRepublicOfKosovoIBAN,
		},
		{
			name: "Kuwait",
			fn:   generateKuwaitIBAN,
		},
		{
			name: "Latvia",
			fn:   generateLatviaIBAN,
		},
		{
			name: "Lebanon",
			fn:   generateLebanonIBAN,
		},
		{
			name: "Principality Of Liechtenstein",
			fn:   generatePrincipalityOfLiechtensteinIBAN,
		},
		{
			name: "Lithuania",
			fn:   generateLithuaniaIBAN,
		},
		{
			name: "Luxembourg",
			fn:   generateLuxembourgIBAN,
		},
		{
			name: "Macedonia",
			fn:   generateMacedoniaIBAN,
		},
		{
			name: "Malta",
			fn:   generateMaltaIBAN,
		},
		{
			name: "Mauritania",
			fn:   generateMauritaniaIBAN,
		},
		{
			name: "Mauritius",
			fn:   generateMauritiusIBAN,
		},
		{
			name: "Moldova",
			fn:   generateMoldovaIBAN,
		},
		{
			name: "Monaco",
			fn:   generateMonacoIBAN,
		},
		{
			name: "Montenegro",
			fn:   generateMontenegroIBAN,
		},
		{
			name: "The Netherlands",
			fn:   generateTheNetherlandsIBAN,
		},
		{
			name: "Norway",
			fn:   generateNorwayIBAN,
		},
		{
			name: "Pakistan",
			fn:   generatePakistanIBAN,
		},
		{
			name: "State Of Palestine",
			fn:   generateStateOfPalestineIBAN,
		},
		{
			name: "Poland",
			fn:   generatePolandIBAN,
		},
		{
			name: "Portugal",
			fn:   generatePortugalIBAN,
		},
		{
			name: "Qatar",
			fn:   generateQatarIBAN,
		},
		{
			name: "Romania",
			fn:   generateRomaniaIBAN,
		},
		{
			name: "Saint Lucia",
			fn:   generateSaintLuciaIBAN,
		},
		{
			name: "San Marino",
			fn:   generateSanMarinoIBAN,
		},
		{
			name: "Sao Tome And Principe",
			fn:   generateSaoTomeAndPrincipeIBAN,
		},
		{
			name: "Saudi Arabia",
			fn:   generateSaudiArabiaIBAN,
		},
		{
			name: "Serbia",
			fn:   generateSerbiaIBAN,
		},
		{
			name: "Slovak Republic",
			fn:   generateSlovakRepublicIBAN,
		},
		{
			name: "Slovenia",
			fn:   generateSloveniaIBAN,
		},
		{
			name: "Spain",
			fn:   generateSpainIBAN,
		},
		{
			name: "Sweden",
			fn:   generateSwedenIBAN,
		},
		{
			name: "Switzerland",
			fn:   generateSwitzerlandIBAN,
		},
		{
			name: "Timor Leste",
			fn:   generateTimorLesteIBAN,
		},
		{
			name: "Tunisia",
			fn:   generateTunisiaIBAN,
		},
		{
			name: "Turkey",
			fn:   generateTurkeyIBAN,
		},
		{
			name: "United Arab Emirates",
			fn:   generateUnitedArabEmiratesIBAN,
		},
		{
			name: "United Kingdom",
			fn:   generateUnitedKingdomIBAN,
		},
		{
			name: "British Virgin Islands",
			fn:   generateBritishVirginIslandsIBAN,
		},
		{
			name: "Seychelles",
			fn:   generateSeychellesIBAN,
		},
		{
			name: "Ukraine",
			fn:   generateUkraineIBAN,
		},
		{
			name: "Somalia",
			fn:   generateSomaliaIBAN,
		},
		{
			name: "Falkland Islands",
			fn:   generateFalklandIslandsIBAN,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			generated, err := tt.fn()
			if err != nil {
				t.Error(err)
				return
			}

			if err := Validate(generated); err != nil {
				t.Error(err)
			}
		})
	}
}
