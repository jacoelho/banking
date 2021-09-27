// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
)

// Generate IBAN based on ISO 3166-1 country code
func Generate(countryCode string) (string, error) {
	switch countryCode {
	case ad:
		return generateAndorraIBAN()
	case ae:
		return generateUnitedArabEmiratesIBAN()
	case al:
		return generateAlbaniaIBAN()
	case at:
		return generateAustriaIBAN()
	case az:
		return generateRepublicOfAzerbaijanIBAN()
	case ba:
		return generateBosniaAndHerzegovinaIBAN()
	case be:
		return generateBelgiumIBAN()
	case bg:
		return generateBulgariaIBAN()
	case bh:
		return generateKingdomOfBahrainIBAN()
	case br:
		return generateBrazilIBAN()
	case by:
		return generateBelarusIBAN()
	case ch:
		return generateSwitzerlandIBAN()
	case cr:
		return generateCostaRicaIBAN()
	case cy:
		return generateCyprusIBAN()
	case cz:
		return generateCzechRepublicIBAN()
	case de:
		return generateGermanyIBAN()
	case dk:
		return generateDenmarkIBAN()
	case do:
		return generateDominicanRepublicIBAN()
	case ee:
		return generateEstoniaIBAN()
	case eg:
		return generateEgyptIBAN()
	case es:
		return generateSpainIBAN()
	case fi:
		return generateFinlandIBAN()
	case fo:
		return generateFaroeIslandsIBAN()
	case fr:
		return generateFranceIBAN()
	case gb:
		return generateUnitedKingdomIBAN()
	case ge:
		return generateGeorgiaIBAN()
	case gi:
		return generateGibraltarIBAN()
	case gl:
		return generateGreenlandIBAN()
	case gr:
		return generateGreeceIBAN()
	case gt:
		return generateGuatemalaIBAN()
	case hr:
		return generateCroatiaIBAN()
	case hu:
		return generateHungaryIBAN()
	case ie:
		return generateIrelandIBAN()
	case il:
		return generateIsraelIBAN()
	case iq:
		return generateIraqIBAN()
	case is:
		return generateIcelandIBAN()
	case it:
		return generateItalyIBAN()
	case jo:
		return generateJordanIBAN()
	case kw:
		return generateKuwaitIBAN()
	case kz:
		return generateKazakhstanIBAN()
	case lb:
		return generateLebanonIBAN()
	case lc:
		return generateSaintLuciaIBAN()
	case li:
		return generatePrincipalityOfLiechtensteinIBAN()
	case lt:
		return generateLithuaniaIBAN()
	case lu:
		return generateLuxembourgIBAN()
	case lv:
		return generateLatviaIBAN()
	case ly:
		return generateLibyaIBAN()
	case mc:
		return generateMonacoIBAN()
	case md:
		return generateMoldovaIBAN()
	case me:
		return generateMontenegroIBAN()
	case mk:
		return generateMacedoniaIBAN()
	case mr:
		return generateMauritaniaIBAN()
	case mt:
		return generateMaltaIBAN()
	case mu:
		return generateMauritiusIBAN()
	case nl:
		return generateTheNetherlandsIBAN()
	case no:
		return generateNorwayIBAN()
	case pk:
		return generatePakistanIBAN()
	case pl:
		return generatePolandIBAN()
	case ps:
		return generateStateOfPalestineIBAN()
	case pt:
		return generatePortugalIBAN()
	case qa:
		return generateQatarIBAN()
	case ro:
		return generateRomaniaIBAN()
	case rs:
		return generateSerbiaIBAN()
	case sa:
		return generateSaudiArabiaIBAN()
	case sc:
		return generateSeychellesIBAN()
	case sd:
		return generateSudanIBAN()
	case se:
		return generateSwedenIBAN()
	case si:
		return generateSloveniaIBAN()
	case sk:
		return generateSlovakRepublicIBAN()
	case sm:
		return generateSanMarinoIBAN()
	case st:
		return generateSaoTomeAndPrincipeIBAN()
	case sv:
		return generateElSalvadorIBAN()
	case tl:
		return generateTimorLesteIBAN()
	case tn:
		return generateTunisiaIBAN()
	case tr:
		return generateTurkeyIBAN()
	case ua:
		return generateUkraineIBAN()
	case va:
		return generateVaticanCityStateIBAN()
	case vg:
		return generateBritishVirginIslandsIBAN()
	case xk:
		return generateRepublicOfKosovoIBAN()
	case gf:
		return generateFrenchGuyanaIBAN()
	case gp:
		return generateGuadeloupeIBAN()
	case mq:
		return generateMartiniqueIBAN()
	case re:
		return generateReunionIBAN()
	case fp:
		return generateFrenchPolynesiaIBAN()
	case tf:
		return generateFrenchSouthernTerritoriesIBAN()
	case yt:
		return generateMayotteIBAN()
	case nc:
		return generateNewCaledoniaIBAN()
	case bl:
		return generateSaintBarthelemyIBAN()
	case mf:
		return generateSaintMartinIBAN()
	case pm:
		return generateSaintPierreEtMiquelonIBAN()
	case wf:
		return generateWallisAndFutunaIslandsIBAN()

	default:
		return "", fmt.Errorf("%s is not supported: %w", countryCode, ErrValidation)
	}
}
