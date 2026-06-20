package iban

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidIBAN reports an invalid IBAN.
	ErrInvalidIBAN = errors.New("invalid IBAN")

	// ErrInvalidCountryCode reports an invalid ISO 3166-1 alpha-2 country code.
	ErrInvalidCountryCode = errors.New("invalid country code")

	// ErrUnsupportedCountry reports a syntactically valid country code that this
	// registry does not support.
	ErrUnsupportedCountry = errors.New("unsupported country")
)

// ValidationReason identifies why IBAN validation failed.
type ValidationReason uint8

const (
	ReasonInvalidLength ValidationReason = iota + 1
	ReasonInvalidChecksum
	ReasonInvalidCharacters
	ReasonUnsupportedCountry
)

func (r ValidationReason) String() string {
	switch r {
	case ReasonInvalidLength:
		return "invalid length"
	case ReasonInvalidChecksum:
		return "invalid checksum"
	case ReasonInvalidCharacters:
		return "invalid characters"
	case ReasonUnsupportedCountry:
		return "unsupported country"
	default:
		return "unknown"
	}
}

// CharClass identifies the expected character class for an IBAN span.
type CharClass uint8

const (
	CharClassDigit CharClass = iota + 1
	CharClassUpperAlpha
	CharClassUpperAlphaNumeric
)

func (c CharClass) String() string {
	switch c {
	case CharClassDigit:
		return "digit"
	case CharClassUpperAlpha:
		return "uppercase letter"
	case CharClassUpperAlphaNumeric:
		return "uppercase letter or digit"
	default:
		return "unknown"
	}
}

// ValidationError contains machine-readable IBAN validation diagnostics.
type ValidationError struct {
	Reason ValidationReason

	// Position and Length identify the invalid IBAN span when applicable.
	Position int
	Length   int

	// Expected identifies the required character class for ReasonInvalidCharacters.
	Expected CharClass
	Actual   string

	// ExpectedValue is set when a precise expected value exists, such as
	// checksum digits or registry static values.
	ExpectedValue string

	// ExpectedLength and ActualLength are set for ReasonInvalidLength.
	ExpectedLength int
	ActualLength   int
}

func (e *ValidationError) Error() string {
	if e == nil {
		return ErrInvalidIBAN.Error()
	}
	switch e.Reason {
	case ReasonInvalidLength:
		return fmt.Sprintf("invalid IBAN length: want %d, got %d", e.ExpectedLength, e.ActualLength)
	case ReasonInvalidChecksum:
		return fmt.Sprintf("invalid IBAN checksum: want %s, got %s", e.ExpectedValue, e.Actual)
	case ReasonInvalidCharacters:
		if e.ExpectedValue != "" {
			return fmt.Sprintf("invalid IBAN value at position %d: want %s, got %s", e.Position, e.ExpectedValue, e.Actual)
		}
		return fmt.Sprintf("invalid IBAN characters at position %d: want %s, got %s", e.Position, e.Expected, e.Actual)
	case ReasonUnsupportedCountry:
		return fmt.Sprintf("unsupported IBAN country %s", e.Actual)
	default:
		return ErrInvalidIBAN.Error()
	}
}

func (e *ValidationError) Is(target error) bool {
	if e == nil {
		return false
	}
	return target == ErrInvalidIBAN ||
		target == ErrUnsupportedCountry && e.Reason == ReasonUnsupportedCountry
}

// CountryCodeError contains machine-readable country-code diagnostics.
type CountryCodeError struct {
	CountryCode string
	Err         error
}

func (e *CountryCodeError) Error() string {
	if e == nil || e.Err == nil {
		return ErrInvalidCountryCode.Error()
	}
	switch {
	case errors.Is(e.Err, ErrUnsupportedCountry):
		return fmt.Sprintf("unsupported country code %s", e.CountryCode)
	case errors.Is(e.Err, ErrInvalidCountryCode):
		return fmt.Sprintf("invalid country code %s", e.CountryCode)
	default:
		return e.Err.Error()
	}
}

func (e *CountryCodeError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

func invalidIBANLength(expected, actual int) error {
	return &ValidationError{
		Reason:         ReasonInvalidLength,
		Length:         expected,
		ExpectedLength: expected,
		ActualLength:   actual,
	}
}

func invalidIBANChecksum(expected, actual string) error {
	return &ValidationError{
		Reason:        ReasonInvalidChecksum,
		Position:      2,
		Length:        2,
		Expected:      CharClassDigit,
		Actual:        actual,
		ExpectedValue: expected,
	}
}

func invalidIBANCharacters(position, length int, expected CharClass, actual string) error {
	return &ValidationError{
		Reason:   ReasonInvalidCharacters,
		Position: position,
		Length:   length,
		Expected: expected,
		Actual:   actual,
	}
}

func invalidIBANValue(position int, expected, actual string) error {
	return &ValidationError{
		Reason:        ReasonInvalidCharacters,
		Position:      position,
		Length:        len(expected),
		Actual:        actual,
		ExpectedValue: expected,
	}
}

func unsupportedIBANCountry(countryCode string) error {
	return &ValidationError{
		Reason:   ReasonUnsupportedCountry,
		Position: 0,
		Length:   len(countryCode),
		Expected: CharClassUpperAlpha,
		Actual:   countryCode,
	}
}

func invalidCountryCode(countryCode string) error {
	return &CountryCodeError{
		CountryCode: countryCode,
		Err:         ErrInvalidCountryCode,
	}
}

func unsupportedCountry(countryCode string) error {
	return &CountryCodeError{
		CountryCode: countryCode,
		Err:         ErrUnsupportedCountry,
	}
}
