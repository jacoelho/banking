package iban

import (
	"fmt"
	"strconv"
	"unsafe"

	"github.com/jacoelho/banking/pool"
)

// normalize converts an IBAN to its digit representation.
func normalize(s string) string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.Grow(len(s) * 2)

	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			val := r - 'A' + 10

			_, _ = sb.WriteRune('0' + val/10)
			_, _ = sb.WriteRune('0' + val%10)
		} else if r >= '0' && r <= '9' {
			_, _ = sb.WriteRune(r)
		}
	}

	return sb.String()
}

func mod9710(number string) int {
	if len(number)%2 != 0 {
		number = "0" + number
	}

	remainder := 0
	for i := 0; i < len(number); i += 2 {
		_ = number[i+1]
		chunk := int(number[i]-'0')*10 + int(number[i+1]-'0')
		remainder = (remainder*100 + chunk) % 97
	}

	return remainder
}

// checksum calculates checksum digits
func checksum(iban string) string {
	rearrangedIBAN := iban[4:] + iban[:2] + "00"

	normalizedIBAN := normalize(rearrangedIBAN)

	modulo := mod9710(normalizedIBAN)

	// Check digits with the value of '01' or '00' are invalid.
	// To resolve an anomaly in the algorithm,
	// values '01' and '00' are equivalent to '98' and '99', respectively, and the latter must be used.
	checkDigits := 98 - modulo
	switch {
	case checkDigits == 0:
		return "99"
	case checkDigits == 1:
		return "98"
	case checkDigits < 10:
		return "0" + strconv.Itoa(checkDigits)
	default:
		return strconv.Itoa(checkDigits)
	}
}

// ReplaceChecksum returns input iban with the correct check digits
func ReplaceChecksum(iban string) (string, error) {
	if len(iban) < 4 {
		return "", fmt.Errorf("invalid iban length: %w", ErrValidation)
	}

	raw := []byte(iban)

	check := checksum(iban)

	raw[2], raw[3] = check[0], check[1]

	return unsafe.String(unsafe.SliceData(raw), len(raw)), nil
}
