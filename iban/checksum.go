package iban

import (
	"fmt"
	"unsafe"

	"github.com/jacoelho/banking/pool"
	"github.com/jacoelho/go-iso7064/iso7064"
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

// checksum calculates checksum digits
func checksum(iban string) string {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	_, _ = sb.WriteString(iban[4:])
	_ = sb.WriteByte(iban[0])
	_ = sb.WriteByte(iban[1])

	digits := iso7064.Modulo97Radix10(normalize(sb.String()))

	// The underlying rules for IBANs is that the account-servicing financial institution should issue an IBAN,
	// as there are a number of areas where different IBANs could be generated from the same account and branch numbers
	// that would satisfy the generic IBAN validation rules.
	// In particular cases where 00 is a valid check digit, 97 will not be a valid check digit, likewise,
	// if 01 is a valid check digit, 98 will not be a valid check digit, similarly with 02 and 99.
	switch digits {
	case "00":
		return "97"
	case "01":
		return "98"
	case "02":
		return "99"
	default:
		return digits
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
