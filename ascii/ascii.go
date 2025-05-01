package ascii

// isDigitTable maps ASCII bytes '0'–'9' to true.
var isDigitTable = [256]bool{
	'0': true, '1': true, '2': true, '3': true, '4': true,
	'5': true, '6': true, '7': true, '8': true, '9': true,
}

// isUpperCaseTable maps 'A'–'Z' to true.
var isUpperCaseTable = [256]bool{
	'A': true, 'B': true, 'C': true, 'D': true, 'E': true,
	'F': true, 'G': true, 'H': true, 'I': true, 'J': true,
	'K': true, 'L': true, 'M': true, 'N': true, 'O': true,
	'P': true, 'Q': true, 'R': true, 'S': true, 'T': true,
	'U': true, 'V': true, 'W': true, 'X': true, 'Y': true,
	'Z': true,
}

// isLowerCaseTable maps 'a'–'z' to true.
var isLowerCaseTable = [256]bool{
	'a': true, 'b': true, 'c': true, 'd': true, 'e': true,
	'f': true, 'g': true, 'h': true, 'i': true, 'j': true,
	'k': true, 'l': true, 'm': true, 'n': true, 'o': true,
	'p': true, 'q': true, 'r': true, 's': true, 't': true,
	'u': true, 'v': true, 'w': true, 'x': true, 'y': true,
	'z': true,
}

// isNumericOrUpperCaseTable maps '0'-'9' and 'A'-'Z' to true.
var isNumericOrUpperCaseTable = mergeTables(isDigitTable, isUpperCaseTable)

// isAlphaTable maps 'a'-'z' and 'A'-'Z' to true.
var isAlphaTable = mergeTables(isLowerCaseTable, isUpperCaseTable)

// isAlphaNumericTable maps '0'-'9', 'a'-'z' and 'A'-'Z' to true.
var isAlphaNumericTable = mergeTables(isDigitTable, isAlphaTable)

// mergeTables returns a new table where each entry is true if
// it was true in *any* of the input tables.
func mergeTables(tables ...[256]bool) [256]bool {
	var result [256]bool
	for _, tbl := range tables {
		for i, v := range tbl {
			result[i] = result[i] || v
		}
	}
	return result
}

// every tests that every byte in s is present in table.
func every(s string, table *[256]bool) bool {
	for i := 0; i < len(s); i++ {
		if !table[s[i]] {
			return false
		}
	}
	return true
}

// IsDigit reports whether every character in s is an ASCII digit ('0'–'9').
// It returns true if s is empty or if every byte in s is in the digit table.
func IsDigit(s string) bool { return every(s, &isDigitTable) }

// IsUpperCase reports whether every character in s is an ASCII uppercase letter ('A'–'Z').
// It returns true if s is empty or if every byte in s is in the uppercase table.
func IsUpperCase(s string) bool { return every(s, &isUpperCaseTable) }

// IsAlphaNumeric reports whether every character in s is an ASCII letter or digit.
// It returns true if s is empty or if every byte in s is in the alphanumeric table.
func IsAlphaNumeric(s string) bool { return every(s, &isAlphaNumericTable) }

// IsUpperAlphaNumeric reports whether every character in s is an ASCII digit or uppercase letter.
// It returns true if s is empty or if every byte in s is in the digit/uppercase table.
func IsUpperAlphaNumeric(s string) bool { return every(s, &isNumericOrUpperCaseTable) }
