package iso7064

import "strings"

func Normalize(s string) string {
	sb := new(strings.Builder)
	sb.Grow(len(s) * 2)

	for _, r := range s {
		switch r {
		case 'A':
			sb.WriteString("10")
		case 'B':
			sb.WriteString("11")
		case 'C':
			sb.WriteString("12")
		case 'D':
			sb.WriteString("13")
		case 'E':
			sb.WriteString("14")
		case 'F':
			sb.WriteString("15")
		case 'G':
			sb.WriteString("16")
		case 'H':
			sb.WriteString("17")
		case 'I':
			sb.WriteString("18")
		case 'J':
			sb.WriteString("19")
		case 'K':
			sb.WriteString("20")
		case 'L':
			sb.WriteString("21")
		case 'M':
			sb.WriteString("22")
		case 'N':
			sb.WriteString("23")
		case 'O':
			sb.WriteString("24")
		case 'P':
			sb.WriteString("25")
		case 'Q':
			sb.WriteString("26")
		case 'R':
			sb.WriteString("27")
		case 'S':
			sb.WriteString("28")
		case 'T':
			sb.WriteString("29")
		case 'U':
			sb.WriteString("30")
		case 'V':
			sb.WriteString("31")
		case 'W':
			sb.WriteString("32")
		case 'X':
			sb.WriteString("33")
		case 'Y':
			sb.WriteString("34")
		case 'Z':
			sb.WriteString("35")
		default:
			if '0' <= r && r <= '9' {
				sb.WriteRune(r)
			}
		}
	}

	return sb.String()
}
