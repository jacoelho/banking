package ascii

func IsLowerCase(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func IsUpperCaseLetter(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func IsDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func IsAlphaNumeric(r rune) bool {
	return IsDigit(r) || IsLowerCase(r) || IsUpperCaseLetter(r)
}

func IsUpperAlphaNumeric(r rune) bool {
	return IsDigit(r) || IsUpperCaseLetter(r)
}

func Every(s string, fn func(r rune) bool) bool {
	for _, r := range s {
		if !fn(r) {
			return false
		}
	}

	return true
}
