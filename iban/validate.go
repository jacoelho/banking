// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
)

// Validate an IBAN
func Validate(iban string) error {
	if len(iban) < 2 {
		return fmt.Errorf("unexpected iban length: %w", ErrValidation)
	}

	code := iban[0:2]
	switch code {
	case ad:
		return validateAndorraIBAN(iban)
	case ae:
		return validateUnitedArabEmiratesIBAN(iban)
	case al:
		return validateAlbaniaIBAN(iban)
	case at:
		return validateAustriaIBAN(iban)
	case az:
		return validateRepublicOfAzerbaijanIBAN(iban)
	case ba:
		return validateBosniaAndHerzegovinaIBAN(iban)
	case be:
		return validateBelgiumIBAN(iban)
	case bg:
		return validateBulgariaIBAN(iban)
	case bh:
		return validateKingdomOfBahrainIBAN(iban)
	case br:
		return validateBrazilIBAN(iban)
	case by:
		return validateBelarusIBAN(iban)
	case ch:
		return validateSwitzerlandIBAN(iban)
	case cr:
		return validateCostaRicaIBAN(iban)
	case cy:
		return validateCyprusIBAN(iban)
	case cz:
		return validateCzechRepublicIBAN(iban)
	case de:
		return validateGermanyIBAN(iban)
	case dk:
		return validateDenmarkIBAN(iban)
	case do:
		return validateDominicanRepublicIBAN(iban)
	case ee:
		return validateEstoniaIBAN(iban)
	case eg:
		return validateEgyptIBAN(iban)
	case es:
		return validateSpainIBAN(iban)
	case fi:
		return validateFinlandIBAN(iban)
	case fo:
		return validateFaroeIslandsIBAN(iban)
	case fr:
		return validateFranceIBAN(iban)
	case gb:
		return validateUnitedKingdomIBAN(iban)
	case ge:
		return validateGeorgiaIBAN(iban)
	case gi:
		return validateGibraltarIBAN(iban)
	case gl:
		return validateGreenlandIBAN(iban)
	case gr:
		return validateGreeceIBAN(iban)
	case gt:
		return validateGuatemalaIBAN(iban)
	case hr:
		return validateCroatiaIBAN(iban)
	case hu:
		return validateHungaryIBAN(iban)
	case ie:
		return validateIrelandIBAN(iban)
	case il:
		return validateIsraelIBAN(iban)
	case iq:
		return validateIraqIBAN(iban)
	case is:
		return validateIcelandIBAN(iban)
	case it:
		return validateItalyIBAN(iban)
	case jo:
		return validateJordanIBAN(iban)
	case kw:
		return validateKuwaitIBAN(iban)
	case kz:
		return validateKazakhstanIBAN(iban)
	case lb:
		return validateLebanonIBAN(iban)
	case lc:
		return validateSaintLuciaIBAN(iban)
	case li:
		return validatePrincipalityOfLiechtensteinIBAN(iban)
	case lt:
		return validateLithuaniaIBAN(iban)
	case lu:
		return validateLuxembourgIBAN(iban)
	case lv:
		return validateLatviaIBAN(iban)
	case ly:
		return validateLibyaIBAN(iban)
	case mc:
		return validateMonacoIBAN(iban)
	case md:
		return validateMoldovaIBAN(iban)
	case me:
		return validateMontenegroIBAN(iban)
	case mk:
		return validateMacedoniaIBAN(iban)
	case mr:
		return validateMauritaniaIBAN(iban)
	case mt:
		return validateMaltaIBAN(iban)
	case mu:
		return validateMauritiusIBAN(iban)
	case nl:
		return validateTheNetherlandsIBAN(iban)
	case no:
		return validateNorwayIBAN(iban)
	case pk:
		return validatePakistanIBAN(iban)
	case pl:
		return validatePolandIBAN(iban)
	case ps:
		return validateStateOfPalestineIBAN(iban)
	case pt:
		return validatePortugalIBAN(iban)
	case qa:
		return validateQatarIBAN(iban)
	case ro:
		return validateRomaniaIBAN(iban)
	case rs:
		return validateSerbiaIBAN(iban)
	case sa:
		return validateSaudiArabiaIBAN(iban)
	case sc:
		return validateSeychellesIBAN(iban)
	case sd:
		return validateSudanIBAN(iban)
	case se:
		return validateSwedenIBAN(iban)
	case si:
		return validateSloveniaIBAN(iban)
	case sk:
		return validateSlovakRepublicIBAN(iban)
	case sm:
		return validateSanMarinoIBAN(iban)
	case st:
		return validateSaoTomeAndPrincipeIBAN(iban)
	case sv:
		return validateElSalvadorIBAN(iban)
	case tl:
		return validateTimorLesteIBAN(iban)
	case tn:
		return validateTunisiaIBAN(iban)
	case tr:
		return validateTurkeyIBAN(iban)
	case ua:
		return validateUkraineIBAN(iban)
	case va:
		return validateVaticanCityStateIBAN(iban)
	case vg:
		return validateBritishVirginIslandsIBAN(iban)
	case xk:
		return validateRepublicOfKosovoIBAN(iban)
	case gf:
		return validateFrenchGuyanaIBAN(iban)
	case gp:
		return validateGuadeloupeIBAN(iban)
	case mq:
		return validateMartiniqueIBAN(iban)
	case re:
		return validateReunionIBAN(iban)
	case fp:
		return validateFrenchPolynesiaIBAN(iban)
	case tf:
		return validateFrenchSouthernTerritoriesIBAN(iban)
	case yt:
		return validateMayotteIBAN(iban)
	case nc:
		return validateNewCaledoniaIBAN(iban)
	case bl:
		return validateSaintBarthelemyIBAN(iban)
	case mf:
		return validateSaintMartinIBAN(iban)
	case pm:
		return validateSaintPierreEtMiquelonIBAN(iban)
	case wf:
		return validateWallisAndFutunaIslandsIBAN(iban)

	default:
		return fmt.Errorf("%s is not supported: %w", code, ErrValidation)
	}
}
