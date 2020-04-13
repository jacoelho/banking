package iban

import (
	"strconv"

	"github.com/jacoelho/banking/iso7064"
)

func Checksum(iban string) string {
	t := []byte(iban)
	value := append(t[4:], t[0], t[1], '0', '0')

	checkDigit := 98 - iso7064.Mod9710(iso7064.Normalize(string(value)))
	if checkDigit == 0 {
		return "00"
	}

	checkString := strconv.FormatInt(int64(checkDigit), 10)
	if len(checkString) < 2 {
		return "0" + checkString
	}

	return checkString
}

func ReplaceChecksum(iban string) string {
	raw := []byte(iban)

	check := Checksum(iban)

	raw[2], raw[3] = check[0], check[1]

	return string(raw)
}
