package ascii

import "math/rand/v2"

const (
	digits            = "0123456789"
	upperCaseLetters  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaUpperNumeric = digits + upperCaseLetters
)

func AppendRandomDigits(dst []byte, length uint8) []byte {
	return appendRandomWithCharset(dst, digits, length)
}

func AppendRandomUpperCaseLetters(dst []byte, length uint8) []byte {
	return appendRandomWithCharset(dst, upperCaseLetters, length)
}

func AppendRandomAlphaNumeric(dst []byte, length uint8) []byte {
	return appendRandomWithCharset(dst, alphaUpperNumeric, length)
}

func appendRandomWithCharset(dst []byte, charset string, length uint8) []byte {
	for range length {
		dst = append(dst, charset[rand.IntN(len(charset))])
	}
	return dst
}
