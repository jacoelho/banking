package iso7064

import "strconv"

const max = 1<<31 - 1

// Mod9710 implements ISO 7064 Mod 97.10 check digit scheme.
// caller must ensure string only contains values between 0-9
func Mod9710(s string) int {
	var result uint

	for _, r := range s {
		result = result*10 + uint(r-'0')
		if result > max {
			result = result % 97
		}
	}

	return int(result % 97)
}

func Mod9710Encoded(s string) string {
	result := Mod9710(s)
	if result == 0 {
		return "00"
	}

	v := strconv.FormatInt(int64(result), 10)
	if result < 9 {
		return "0" + v
	}
	return v
}
