package iban

import (
	"fmt"
	"strconv"
	"unsafe"

	"github.com/jacoelho/banking/pool"
)

func normalize(s string) string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	for _, r := range s {
		switch r {
		case 'A':
			_, _ = sb.WriteString("10")
		case 'B':
			_, _ = sb.WriteString("11")
		case 'C':
			_, _ = sb.WriteString("12")
		case 'D':
			_, _ = sb.WriteString("13")
		case 'E':
			_, _ = sb.WriteString("14")
		case 'F':
			_, _ = sb.WriteString("15")
		case 'G':
			_, _ = sb.WriteString("16")
		case 'H':
			_, _ = sb.WriteString("17")
		case 'I':
			_, _ = sb.WriteString("18")
		case 'J':
			_, _ = sb.WriteString("19")
		case 'K':
			_, _ = sb.WriteString("20")
		case 'L':
			_, _ = sb.WriteString("21")
		case 'M':
			_, _ = sb.WriteString("22")
		case 'N':
			_, _ = sb.WriteString("23")
		case 'O':
			_, _ = sb.WriteString("24")
		case 'P':
			_, _ = sb.WriteString("25")
		case 'Q':
			_, _ = sb.WriteString("26")
		case 'R':
			_, _ = sb.WriteString("27")
		case 'S':
			_, _ = sb.WriteString("28")
		case 'T':
			_, _ = sb.WriteString("29")
		case 'U':
			_, _ = sb.WriteString("30")
		case 'V':
			_, _ = sb.WriteString("31")
		case 'W':
			_, _ = sb.WriteString("32")
		case 'X':
			_, _ = sb.WriteString("33")
		case 'Y':
			_, _ = sb.WriteString("34")
		case 'Z':
			_, _ = sb.WriteString("35")
		default:
			if '0' <= r && r <= '9' {
				_, _ = sb.WriteRune(r)
			}
		}
	}

	return sb.String()
}

var paddings = [...]int{
	1: 10,
	2: 100,
	3: 1000,
	4: 10000,
	5: 100000,
	6: 1000000,
	7: 10000000,
	8: 100000000,
	9: 1000000000,
}

// mod9710Chunked computes mod97-10 checksum in chucks
func mod9710Chunked(s string) int {
	var padding = 1000000000
	var sum int

	for i := 0; i < len(s); i += 9 {
		end := i + 9

		if end > len(s) {
			end = len(s)
			padding = paddings[len(s[i:end])]
		}

		v, err := strconv.Atoi(s[i:end])
		if err != nil {
			panic(err)
		}

		sum = (sum*padding + v) % 97
	}

	return sum
}

// checksum calculates checksum digits
func checksum(iban string) string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	sb.Grow(len(iban))

	_, _ = sb.WriteString(iban[4:])
	_, _ = sb.WriteString(iban[:2])
	_, _ = sb.WriteString("00")

	result := mod9710Chunked(normalize(sb.String()))

	// Check digits with the value of '01' or '00' are invalid.
	// To resolve an anomaly in the algorithm,
	// values '01' and '00' are equivalent to '98' and '99', respectively, and the latter must be used.
	checkDigits := 98 - result
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

	return *(*string)(unsafe.Pointer(&raw)), nil
}
