package iban

import "unicode"

type CountryCode string

type IBAN interface {
	CountryCode() CountryCode
	String() string
}

type BBAN interface {
	String() string
}

func IsDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func IsLetter(r rune) bool {
	return unicode.IsLetter(r)
}

// BR 2!n 8!n 5!n 10!n 1!a 1!c
// 8!n16!c
func valid(s string) string {
	return ""
}
