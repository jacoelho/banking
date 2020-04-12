package iso7064

//const max = 999999999
const max = 1<<31 - 1

// Mod97 expects string to be normalized
func Mod97(s string) int {
	var result int

	for _, r := range s {
		num := int(r - '0')

		result = result*10 + num
		if result > max {
			result = result % 97
		}
	}
	return result % 97
}
