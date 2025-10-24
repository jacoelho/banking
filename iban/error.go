package iban

import (
	"errors"
	"fmt"
)

// CharacterType represents the expected character type in IBAN range validation
type CharacterType int

const (
	CharacterTypeDigit CharacterType = iota
	CharacterTypeUpperCase
	CharacterTypeAlphaNumeric
)

// String returns the string representation of CharacterType
func (ct CharacterType) String() string {
	switch ct {
	case CharacterTypeDigit:
		return "Digit"
	case CharacterTypeUpperCase:
		return "UpperCase"
	case CharacterTypeAlphaNumeric:
		return "AlphaNumeric"
	default:
		return "Unknown"
	}
}

// ErrValidationLength represents an IBAN length validation error
type ErrValidationLength struct {
	Expected int
	Actual   int
}

func (e *ErrValidationLength) Error() string {
	return fmt.Sprintf("unexpected length, want: %d, got: %d", e.Expected, e.Actual)
}

// ErrValidationChecksum represents an IBAN checksum validation error
type ErrValidationChecksum struct {
	Expected string
	Actual   string
}

func (e *ErrValidationChecksum) Error() string {
	return fmt.Sprintf("incorrect checksum, expected: %s, got: %s", e.Expected, e.Actual)
}

// ErrValidationRange represents an IBAN range validation error
type ErrValidationRange struct {
	Position int
	Length   int
	Expected CharacterType
	Actual   string
}

func (e *ErrValidationRange) Error() string {
	return fmt.Sprintf("range rule, start pos: %d, length: %d, expected type: %s, found: %s",
		e.Position, e.Length, e.Expected.String(), e.Actual)
}

// ErrValidationStaticValue represents an IBAN static value validation error
type ErrValidationStaticValue struct {
	Position int
	Expected string
	Actual   string
}

func (e *ErrValidationStaticValue) Error() string {
	return fmt.Sprintf("static value rule, pos: %d, expected value: %s, found: %s",
		e.Position, e.Expected, e.Actual)
}

// ErrUnsupportedCountry represents an error for unsupported country codes
type ErrUnsupportedCountry struct {
	CountryCode string
}

func (e *ErrUnsupportedCountry) Error() string {
	return fmt.Sprintf("country code %s is not supported", e.CountryCode)
}

// IsValidationError reports whether err is any validation error from this library
func IsValidationError(err error) bool {
	if err == nil {
		return false
	}

	var lengthErr *ErrValidationLength
	var checksumErr *ErrValidationChecksum
	var rangeErr *ErrValidationRange
	var staticErr *ErrValidationStaticValue
	var unsupportedErr *ErrUnsupportedCountry

	return errors.As(err, &lengthErr) ||
		errors.As(err, &checksumErr) ||
		errors.As(err, &rangeErr) ||
		errors.As(err, &staticErr) ||
		errors.As(err, &unsupportedErr)
}
