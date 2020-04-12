// Code generated by banking/generator; DO NOT EDIT.

package iban

import (
	"fmt"
)

// Generate IBAN based on ISO 3166-1 country code
func Generate(countryCode string) (string, error) {
	var result string

	switch countryCode {
	case "AL":
		result = GenerateAlbaniaIBAN()
	case "AD":
		result = GenerateAndorraIBAN()
	case "AT":
		result = GenerateAustriaIBAN()
	case "AZ":
		result = GenerateRepublicOfAzerbaijanIBAN()
	case "BH":
		result = GenerateKingdomOfBahrainIBAN()
	case "BE":
		result = GenerateBelgiumIBAN()
	case "BA":
		result = GenerateBosniaAndHerzegovinaIBAN()
	case "BR":
		result = GenerateBrazilIBAN()
	case "BG":
		result = GenerateBulgariaIBAN()
	case "CR":
		result = GenerateCostaRicaIBAN()
	case "HR":
		result = GenerateCroatiaIBAN()
	case "CY":
		result = GenerateCyprusIBAN()
	case "CZ":
		result = GenerateCzechRepublicIBAN()
	case "DK":
		result = GenerateDenmarkIBAN()
	case "FO":
		result = GenerateFaroeIslandsIBAN()
	case "GL":
		result = GenerateGreenlandIBAN()
	case "DO":
		result = GenerateDominicanRepublicIBAN()
	case "EE":
		result = GenerateEstoniaIBAN()
	case "FI":
		result = GenerateFinlandIBAN()
	case "FR":
		result = GenerateFranceIBAN()
	case "GF":
		result = GenerateFrenchGuyanaIBAN()
	case "GP":
		result = GenerateGuadeloupeIBAN()
	case "MQ":
		result = GenerateMartiniqueIBAN()
	case "RE":
		result = GenerateReunionIBAN()
	case "FP":
		result = GenerateFrenchPolynesiaIBAN()
	case "TF":
		result = GenerateFrenchSouthernTerritoriesIBAN()
	case "YT":
		result = GenerateMayotteIBAN()
	case "NC":
		result = GenerateNewCaledoniaIBAN()
	case "BL":
		result = GenerateSaintBarthelemyIBAN()
	case "MF":
		result = GenerateSaintMartinIBAN()
	case "PM":
		result = GenerateSaintPierreEtMiquelonIBAN()
	case "WF":
		result = GenerateWallisAndFutunaIslandsIBAN()
	case "GE":
		result = GenerateGeorgiaIBAN()
	case "DE":
		result = GenerateGermanyIBAN()
	case "GI":
		result = GenerateGibraltarIBAN()
	case "GR":
		result = GenerateGreeceIBAN()
	case "GT":
		result = GenerateGuatemalaIBAN()
	case "HU":
		result = GenerateHungaryIBAN()
	case "IS":
		result = GenerateIcelandIBAN()
	case "IE":
		result = GenerateIrelandIBAN()
	case "IL":
		result = GenerateIsraelIBAN()
	case "IT":
		result = GenerateItalyIBAN()
	case "JO":
		result = GenerateJordanIBAN()
	case "KZ":
		result = GenerateKazakhstanIBAN()
	case "XK":
		result = GenerateRepublicOfKosovoIBAN()
	case "KW":
		result = GenerateKuwaitIBAN()
	case "LV":
		result = GenerateLatviaIBAN()
	case "LB":
		result = GenerateLebanonIBAN()
	case "LI":
		result = GeneratePrincipalityOfLiechtensteinIBAN()
	case "LT":
		result = GenerateLithuaniaIBAN()
	case "LU":
		result = GenerateLuxembourgIBAN()
	case "MK":
		result = GenerateMacedoniaIBAN()
	case "MT":
		result = GenerateMaltaIBAN()
	case "MR":
		result = GenerateMauritaniaIBAN()
	case "MU":
		result = GenerateMauritiusIBAN()
	case "MD":
		result = GenerateMoldovaIBAN()
	case "MC":
		result = GenerateMonacoIBAN()
	case "ME":
		result = GenerateMontenegroIBAN()
	case "NL":
		result = GenerateTheNetherlandsIBAN()
	case "NO":
		result = GenerateNorwayIBAN()
	case "PK":
		result = GeneratePakistanIBAN()
	case "PS":
		result = GenerateStateOfPalestineIBAN()
	case "PL":
		result = GeneratePolandIBAN()
	case "PT":
		result = GeneratePortugalIBAN()
	case "QA":
		result = GenerateQatarIBAN()
	case "RO":
		result = GenerateRomaniaIBAN()
	case "LC":
		result = GenerateSaintLuciaIBAN()
	case "SM":
		result = GenerateSanMarinoIBAN()
	case "ST":
		result = GenerateSaoTomeAndPrincipeIBAN()
	case "SA":
		result = GenerateSaudiArabiaIBAN()
	case "RS":
		result = GenerateSerbiaIBAN()
	case "SK":
		result = GenerateSlovakRepublicIBAN()
	case "SI":
		result = GenerateSloveniaIBAN()
	case "ES":
		result = GenerateSpainIBAN()
	case "SE":
		result = GenerateSwedenIBAN()
	case "CH":
		result = GenerateSwitzerlandIBAN()
	case "TL":
		result = GenerateTimorLesteIBAN()
	case "TN":
		result = GenerateTunisiaIBAN()
	case "TR":
		result = GenerateTurkeyIBAN()
	case "AE":
		result = GenerateUnitedArabEmiratesIBAN()
	case "GB":
		result = GenerateUnitedKingdomIBAN()
	case "VG":
		result = GenerateBritishVirginIslandsIBAN()
	case "SC":
		result = GenerateSeychellesIBAN()
	case "UA":
		result = GenerateUkraineIBAN()

	default:
		return "", fmt.Errorf("%s is not supported: %w", countryCode, ErrValidation)
	}

	return result, nil
}
