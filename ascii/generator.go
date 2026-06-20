package ascii

import "math/rand/v2"

const (
	digits            = "0123456789"
	upperCaseLetters  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaUpperNumeric = digits + upperCaseLetters
)

func AppendRandomDigits(dst []byte, length int) []byte {
	return appendRandomWithCharset(dst, digits, length)
}

func AppendRandomUpperCaseLetters(dst []byte, length int) []byte {
	return appendRandomWithCharset(dst, upperCaseLetters, length)
}

func AppendRandomAlphaNumeric(dst []byte, length int) []byte {
	return appendRandomWithCharset(dst, alphaUpperNumeric, length)
}

func RandomDigits(length int) string {
	return randomStringWithCharset(digits, length)
}

func RandomUpperCaseLetters(length int) string {
	return randomStringWithCharset(upperCaseLetters, length)
}

func RandomAlphaNumeric(length int) string {
	return randomStringWithCharset(alphaUpperNumeric, length)
}

func randomStringWithCharset(charset string, length int) string {
	return string(appendRandomWithCharset(make([]byte, 0, max(0, length)), charset, length))
}

func appendRandomWithCharset(dst []byte, charset string, length int) []byte {
	for range length {
		dst = append(dst, charset[rand.IntN(len(charset))])
	}
	return dst
}
