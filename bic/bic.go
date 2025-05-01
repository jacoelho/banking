package bic

import (
	"errors"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/iso3166"
)

var ErrBICNotValid = errors.New("bic is not valid")

const HeadOfficeBranchCode = "XXX"

// BIC represents Business Identifier Code
// The standard is defined in ISO 9362:2014
// BusinessPartyPrefix is limited to alphabetic values following SWIFT rules.
type BIC struct {
	BusinessPartyPrefix string
	CountryCode         string
	BusinessPartySuffix string
	BranchCode          string
}

// isSameBranchCode compares bic branch codes
// in this context XXX and 123 will be equal
func (b BIC) isSameBranchCode(o BIC) bool {
	if b.BranchCode == HeadOfficeBranchCode || o.BranchCode == HeadOfficeBranchCode {
		return true
	}

	return b.BranchCode == o.CountryCode
}

// SameInstitution checks if two bic belong to the same institutions
// ESSEDE5FXXX and ESSEDE5F100 will return true.
func (b BIC) SameInstitution(other BIC) bool {
	if !b.IsValid() || !other.IsValid() {
		return false
	}

	return b.BusinessPartyPrefix == other.BusinessPartyPrefix &&
		b.CountryCode == other.CountryCode &&
		b.BusinessPartySuffix == other.BusinessPartySuffix &&
		b.isSameBranchCode(other)
}

// BIC8 returns the bic formatted as bic with length 8.
func (b BIC) BIC8() string {
	return b.BusinessPartyPrefix + b.CountryCode + b.BusinessPartySuffix
}

// Implements fmt.Stringer
// returns the bic in bic 11 format.
func (b BIC) String() string {
	return b.BIC8() + b.BranchCode
}

// IsValid returns if a bic is valid.
func (b BIC) IsValid() bool {
	length := len(b.BusinessPartyPrefix) +
		len(b.CountryCode) +
		len(b.BusinessPartySuffix) +
		len(b.BranchCode)

	if length != 11 {
		return false
	}

	if !(iso3166.IsCountryCode(b.CountryCode) &&
		ascii.IsUpperCase(b.BusinessPartyPrefix) &&
		ascii.IsUpperAlphaNumeric(b.BusinessPartySuffix) &&
		ascii.IsUpperAlphaNumeric(b.BranchCode)) {
		return false
	}

	return true
}

// Parse parses a string as a bic.
// If a BIC8 is passed, branch code will be set to XXX.
func Parse(bic string) (BIC, error) {
	if len(bic) != 8 && len(bic) != 11 {
		return BIC{}, ErrBICNotValid
	}

	b := BIC{
		BusinessPartyPrefix: bic[0:4],
		CountryCode:         bic[4:6],
		BusinessPartySuffix: bic[6:8],
		BranchCode:          HeadOfficeBranchCode,
	}

	if len(bic) == 11 {
		b.BranchCode = bic[8:]
	}

	if !b.IsValid() {
		return BIC{}, ErrBICNotValid
	}

	return b, nil
}

// MarshalText implements encoding.TextMarshaler.
func (b BIC) MarshalText() (text []byte, err error) {
	return []byte(b.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (b *BIC) UnmarshalText(data []byte) error {
	bic, err := Parse(string(data))
	if err != nil {
		return err
	}
	*b = bic
	return nil
}
