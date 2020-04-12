package iso7064

const max = 1<<31 - 1

// Mod97 expects string to be normalized
func Mod97(s string) uint {
	var result uint

	for _, r := range s {
		result = result*10 + uint(r-'0')
		if result > max {
			result = result % 97
		}
	}
	return result % 97
}
