package iso7064

import "strconv"

// Mod112 implements ISO 7064 Mod 11,2 check digit scheme.
// caller must ensure string only contains values between 0-9 and X
func Mod112(s string) int {
	helper := func(r rune) int {
		if r == 'x' || r == 'X' {
			return 10
		}
		return int(r - '0')
	}

	var check = 0
	for _, r := range s {
		check = 2 * (check + helper(r))
	}

	check = check % 11
	return (12 - check) % 11
}

// Mod112Encoded returns Mod 11,2 encoded
func Mod112Encoded(s string) string {
	check := Mod112(s)
	if check == 10 {
		return "X"
	}
	return strconv.Itoa(check)
}
