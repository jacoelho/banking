package ascii

import (
	"math/rand"
)

const (
	digits            = "0123456789"
	lowerCaseLetters  = "abcdefghijklmnopqrstuvwxyz"
	upperCaseLetters  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaUpperNumeric = digits + upperCaseLetters
)

type RuneWriter interface {
	WriteRune(r rune) (int, error)
}

func Digits(sb RuneWriter, length int) {
	stringWithCharset(sb, digits, length)
}

func LowerCaseLetters(sb RuneWriter, length int) {
	stringWithCharset(sb, lowerCaseLetters, length)
}

func UpperCaseLetters(sb RuneWriter, length int) {
	stringWithCharset(sb, upperCaseLetters, length)
}

func AlphaNumeric(sb RuneWriter, length int) {
	stringWithCharset(sb, alphaUpperNumeric, length)
}

func stringWithCharset(sb RuneWriter, charset string, length int) {
	for range length {
		_, _ = sb.WriteRune(rune(charset[rand.Intn(len(charset))]))
	}
}
