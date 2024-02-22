// Code generated by banking/registry; DO NOT EDIT.

package iban

import (
	"fmt"
)

// IsSEPACountryCode returns if a country code is a SEPA member
func IsSEPACountryCode(countryCode string) (bool, error) {
	if len(countryCode) != 2 {
		return false, fmt.Errorf("unexpected country code length: %w", ErrValidation)
	}

	switch countryCode {
	// Country Code Andorra
	case ad:
		return true, nil
	// Country Code United Arab Emirates
	case ae:
		return false, nil
	// Country Code Albania
	case al:
		return false, nil
	// Country Code Austria
	case at:
		return true, nil
	// Country Code Republic Of Azerbaijan
	case az:
		return false, nil
	// Country Code Bosnia And Herzegovina
	case ba:
		return false, nil
	// Country Code Belgium
	case be:
		return true, nil
	// Country Code Bulgaria
	case bg:
		return true, nil
	// Country Code Kingdom Of Bahrain
	case bh:
		return false, nil
	// Country Code Burundi
	case bi:
		return false, nil
	// Country Code Brazil
	case br:
		return false, nil
	// Country Code Belarus
	case by:
		return false, nil
	// Country Code Switzerland
	case ch:
		return true, nil
	// Country Code Costa Rica
	case cr:
		return false, nil
	// Country Code Cyprus
	case cy:
		return true, nil
	// Country Code Czech Republic
	case cz:
		return true, nil
	// Country Code Germany
	case de:
		return true, nil
	// Country Code Djibouti
	case dj:
		return false, nil
	// Country Code Denmark
	case dk:
		return true, nil
	// Country Code Dominican Republic
	case do:
		return false, nil
	// Country Code Estonia
	case ee:
		return true, nil
	// Country Code Egypt
	case eg:
		return false, nil
	// Country Code Spain
	case es:
		return true, nil
	// Country Code Finland
	case fi:
		return true, nil
	// Country Code Falkland Islands
	case fk:
		return false, nil
	// Country Code Faroe Islands
	case fo:
		return false, nil
	// Country Code France
	case fr:
		return true, nil
	// Country Code United Kingdom
	case gb:
		return true, nil
	// Country Code Georgia
	case ge:
		return false, nil
	// Country Code Gibraltar
	case gi:
		return true, nil
	// Country Code Greenland
	case gl:
		return false, nil
	// Country Code Greece
	case gr:
		return true, nil
	// Country Code Guatemala
	case gt:
		return false, nil
	// Country Code Croatia
	case hr:
		return true, nil
	// Country Code Hungary
	case hu:
		return true, nil
	// Country Code Ireland
	case ie:
		return true, nil
	// Country Code Israel
	case il:
		return false, nil
	// Country Code Iraq
	case iq:
		return false, nil
	// Country Code Iceland
	case is:
		return true, nil
	// Country Code Italy
	case it:
		return true, nil
	// Country Code Jordan
	case jo:
		return false, nil
	// Country Code Kuwait
	case kw:
		return false, nil
	// Country Code Kazakhstan
	case kz:
		return false, nil
	// Country Code Lebanon
	case lb:
		return false, nil
	// Country Code Saint Lucia
	case lc:
		return false, nil
	// Country Code Principality Of Liechtenstein
	case li:
		return true, nil
	// Country Code Lithuania
	case lt:
		return true, nil
	// Country Code Luxembourg
	case lu:
		return true, nil
	// Country Code Latvia
	case lv:
		return true, nil
	// Country Code Libya
	case ly:
		return false, nil
	// Country Code Monaco
	case mc:
		return true, nil
	// Country Code Moldova
	case md:
		return false, nil
	// Country Code Montenegro
	case me:
		return false, nil
	// Country Code Macedonia
	case mk:
		return false, nil
	// Country Code Mongolia
	case mn:
		return false, nil
	// Country Code Mauritania
	case mr:
		return false, nil
	// Country Code Malta
	case mt:
		return true, nil
	// Country Code Mauritius
	case mu:
		return false, nil
	// Country Code Nicaragua
	case ni:
		return false, nil
	// Country Code The Netherlands
	case nl:
		return true, nil
	// Country Code Norway
	case no:
		return true, nil
	// Country Code Oman
	case om:
		return false, nil
	// Country Code Pakistan
	case pk:
		return false, nil
	// Country Code Poland
	case pl:
		return true, nil
	// Country Code State Of Palestine
	case ps:
		return false, nil
	// Country Code Portugal
	case pt:
		return true, nil
	// Country Code Qatar
	case qa:
		return false, nil
	// Country Code Romania
	case ro:
		return true, nil
	// Country Code Serbia
	case rs:
		return false, nil
	// Country Code Russia
	case ru:
		return false, nil
	// Country Code Saudi Arabia
	case sa:
		return false, nil
	// Country Code Seychelles
	case sc:
		return false, nil
	// Country Code Sudan
	case sd:
		return false, nil
	// Country Code Sweden
	case se:
		return true, nil
	// Country Code Slovenia
	case si:
		return true, nil
	// Country Code Slovak Republic
	case sk:
		return true, nil
	// Country Code San Marino
	case sm:
		return true, nil
	// Country Code Somalia
	case so:
		return false, nil
	// Country Code Sao Tome And Principe
	case st:
		return false, nil
	// Country Code El Salvador
	case sv:
		return false, nil
	// Country Code Timor Leste
	case tl:
		return false, nil
	// Country Code Tunisia
	case tn:
		return false, nil
	// Country Code Turkey
	case tr:
		return false, nil
	// Country Code Ukraine
	case ua:
		return false, nil
	// Country Code Vatican City State
	case va:
		return true, nil
	// Country Code British Virgin Islands
	case vg:
		return false, nil
	// Country Code Republic Of Kosovo
	case xk:
		return false, nil
	// Country Code French Guyana
	case gf:
		return true, nil
	// Country Code Guadeloupe
	case gp:
		return true, nil
	// Country Code Martinique
	case mq:
		return true, nil
	// Country Code Reunion
	case re:
		return true, nil
	// Country Code French Polynesia
	case fp:
		return false, nil
	// Country Code French Southern Territories
	case tf:
		return false, nil
	// Country Code Mayotte
	case yt:
		return true, nil
	// Country Code New Caledonia
	case nc:
		return false, nil
	// Country Code Saint Barthelemy
	case bl:
		return true, nil
	// Country Code Saint Martin
	case mf:
		return true, nil
	// Country Code Saint Pierre Et Miquelon
	case pm:
		return false, nil
	// Country Code Wallis And Futuna Islands
	case wf:
		return false, nil
	default:
		return false, fmt.Errorf("%s is not supported: %w", countryCode, ErrValidation)
	}
}

// IsSEPA returns if an iban country is a SEPA member
func IsSEPA(iban string) (bool, error) {
	if len(iban) < 2 {
		return false, fmt.Errorf("unexpected iban length: %w", ErrValidation)
	}

	return IsSEPACountryCode(iban[0:2])
}
