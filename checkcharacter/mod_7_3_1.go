package checkcharacter

import (
	"strconv"
)

// 7-3-1 method
// https://www.pangaliit.ee/settlements-and-standards/reference-number-of-the-invoice

var weights = [...]uint{7, 3, 1}

func Mod731(s string) int {
	length := len(s)

	var result uint
	for i := 0; i < length; i++ {
		result += uint(s[length-1-i]-'0') * weights[i%3]
	}

	return 10 - int(result%10)
}

func Mod731Encoded(s string) string {
	return s + strconv.Itoa(Mod731(s))
}
