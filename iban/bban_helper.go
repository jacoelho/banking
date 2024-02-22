// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
)

// GetBBAN retrieves BBAN from an iban
func GetBBAN(iban string) (BBAN, error) {
	if len(iban) < 2 {
		return BBAN{}, fmt.Errorf("unexpected iban length: %w", ErrValidation)
	}

	code := iban[0:2]
	switch code {
	case ad:
		return getAndorraBBAN(iban)
	case ae:
		return getUnitedArabEmiratesBBAN(iban)
	case al:
		return getAlbaniaBBAN(iban)
	case at:
		return getAustriaBBAN(iban)
	case az:
		return getRepublicOfAzerbaijanBBAN(iban)
	case ba:
		return getBosniaAndHerzegovinaBBAN(iban)
	case be:
		return getBelgiumBBAN(iban)
	case bg:
		return getBulgariaBBAN(iban)
	case bh:
		return getKingdomOfBahrainBBAN(iban)
	case bi:
		return getBurundiBBAN(iban)
	case br:
		return getBrazilBBAN(iban)
	case by:
		return getBelarusBBAN(iban)
	case ch:
		return getSwitzerlandBBAN(iban)
	case cr:
		return getCostaRicaBBAN(iban)
	case cy:
		return getCyprusBBAN(iban)
	case cz:
		return getCzechRepublicBBAN(iban)
	case de:
		return getGermanyBBAN(iban)
	case dj:
		return getDjiboutiBBAN(iban)
	case dk:
		return getDenmarkBBAN(iban)
	case do:
		return getDominicanRepublicBBAN(iban)
	case ee:
		return getEstoniaBBAN(iban)
	case eg:
		return getEgyptBBAN(iban)
	case es:
		return getSpainBBAN(iban)
	case fi:
		return getFinlandBBAN(iban)
	case fk:
		return getFalklandIslandsBBAN(iban)
	case fo:
		return getFaroeIslandsBBAN(iban)
	case fr:
		return getFranceBBAN(iban)
	case gb:
		return getUnitedKingdomBBAN(iban)
	case ge:
		return getGeorgiaBBAN(iban)
	case gi:
		return getGibraltarBBAN(iban)
	case gl:
		return getGreenlandBBAN(iban)
	case gr:
		return getGreeceBBAN(iban)
	case gt:
		return getGuatemalaBBAN(iban)
	case hr:
		return getCroatiaBBAN(iban)
	case hu:
		return getHungaryBBAN(iban)
	case ie:
		return getIrelandBBAN(iban)
	case il:
		return getIsraelBBAN(iban)
	case iq:
		return getIraqBBAN(iban)
	case is:
		return getIcelandBBAN(iban)
	case it:
		return getItalyBBAN(iban)
	case jo:
		return getJordanBBAN(iban)
	case kw:
		return getKuwaitBBAN(iban)
	case kz:
		return getKazakhstanBBAN(iban)
	case lb:
		return getLebanonBBAN(iban)
	case lc:
		return getSaintLuciaBBAN(iban)
	case li:
		return getPrincipalityOfLiechtensteinBBAN(iban)
	case lt:
		return getLithuaniaBBAN(iban)
	case lu:
		return getLuxembourgBBAN(iban)
	case lv:
		return getLatviaBBAN(iban)
	case ly:
		return getLibyaBBAN(iban)
	case mc:
		return getMonacoBBAN(iban)
	case md:
		return getMoldovaBBAN(iban)
	case me:
		return getMontenegroBBAN(iban)
	case mk:
		return getMacedoniaBBAN(iban)
	case mn:
		return getMongoliaBBAN(iban)
	case mr:
		return getMauritaniaBBAN(iban)
	case mt:
		return getMaltaBBAN(iban)
	case mu:
		return getMauritiusBBAN(iban)
	case ni:
		return getNicaraguaBBAN(iban)
	case nl:
		return getTheNetherlandsBBAN(iban)
	case no:
		return getNorwayBBAN(iban)
	case om:
		return getOmanBBAN(iban)
	case pk:
		return getPakistanBBAN(iban)
	case pl:
		return getPolandBBAN(iban)
	case ps:
		return getStateOfPalestineBBAN(iban)
	case pt:
		return getPortugalBBAN(iban)
	case qa:
		return getQatarBBAN(iban)
	case ro:
		return getRomaniaBBAN(iban)
	case rs:
		return getSerbiaBBAN(iban)
	case ru:
		return getRussiaBBAN(iban)
	case sa:
		return getSaudiArabiaBBAN(iban)
	case sc:
		return getSeychellesBBAN(iban)
	case sd:
		return getSudanBBAN(iban)
	case se:
		return getSwedenBBAN(iban)
	case si:
		return getSloveniaBBAN(iban)
	case sk:
		return getSlovakRepublicBBAN(iban)
	case sm:
		return getSanMarinoBBAN(iban)
	case so:
		return getSomaliaBBAN(iban)
	case st:
		return getSaoTomeAndPrincipeBBAN(iban)
	case sv:
		return getElSalvadorBBAN(iban)
	case tl:
		return getTimorLesteBBAN(iban)
	case tn:
		return getTunisiaBBAN(iban)
	case tr:
		return getTurkeyBBAN(iban)
	case ua:
		return getUkraineBBAN(iban)
	case va:
		return getVaticanCityStateBBAN(iban)
	case vg:
		return getBritishVirginIslandsBBAN(iban)
	case xk:
		return getRepublicOfKosovoBBAN(iban)
	case gf:
		return getFrenchGuyanaBBAN(iban)
	case gp:
		return getGuadeloupeBBAN(iban)
	case mq:
		return getMartiniqueBBAN(iban)
	case re:
		return getReunionBBAN(iban)
	case fp:
		return getFrenchPolynesiaBBAN(iban)
	case tf:
		return getFrenchSouthernTerritoriesBBAN(iban)
	case yt:
		return getMayotteBBAN(iban)
	case nc:
		return getNewCaledoniaBBAN(iban)
	case bl:
		return getSaintBarthelemyBBAN(iban)
	case mf:
		return getSaintMartinBBAN(iban)
	case pm:
		return getSaintPierreEtMiquelonBBAN(iban)
	case wf:
		return getWallisAndFutunaIslandsBBAN(iban)

	default:
		return BBAN{}, fmt.Errorf("%s is not supported: %w", code, ErrValidation)
	}
}
