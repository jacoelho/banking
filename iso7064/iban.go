package iso7064

import "strconv"

func IBANChecksum(iban string) string {
	t := []byte(iban)
	value := append(t[4:], t[0], t[1], '0', '0')

	check := 98 - Mod97(Normalize(string(value)))

	checksumString := strconv.FormatInt(int64(int(check)), 10)
	if len(checksumString) < 2 {
		return "0" + checksumString
	}

	return checksumString
}
