package iso7064

import "strconv"

// Mod1110 implements ISO 7064 Mod 11,10 check digit scheme.
// caller must ensure string only contains values between 0-9
func Mod1110(s string) int {
	helper := func(x int) int {
		if val := x % 10; val != 0 {
			return val
		}
		return 10
	}

	var sum = 10
	for _, r := range s {
		sum = 2 * helper(sum+int(r-'0')) % 11
	}

	return (11 - sum) % 10
}

// Mod1110Encoded returns Mod 11,10 encoded
func Mod1110Encoded(s string) string {
	return strconv.Itoa(Mod1110(s))
}
