package iban

const decimalDigits = "0123456789"

// ReplaceChecksum returns the IBAN with corrected check digits.
func ReplaceChecksum(iban string) (string, error) {
	if len(iban) < 4 {
		return "", &ErrValidationLength{Expected: 4, Actual: len(iban)}
	}
	if err := validateIBANStructure(iban); err != nil {
		return "", err
	}

	return replaceChecksumBytes([]byte(iban)), nil
}

func replaceChecksumBytes(iban []byte) string {
	calculateCheckDigits(iban, iban[2:4])
	return string(iban)
}

func validateChecksum(iban string) error {
	expected := checksumBytes(iban)
	if expected[0] == iban[2] && expected[1] == iban[3] {
		return nil
	}
	return &ErrValidationChecksum{Expected: string(expected[:]), Actual: iban[2:4]}
}

// calculateCheckDigits calculates IBAN check digits and writes them to checkBuf.
func calculateCheckDigits[T ~string | ~[]byte](iban T, checkBuf []byte) {
	var stack [68]byte
	digits := stack[:]
	if len(iban)*2 > len(digits) {
		digits = make([]byte, len(iban)*2)
	}
	digitLen := normalizeRearrangedIBAN(iban, digits)
	modulo := checksumModulo97(digits, digitLen)
	checkDigits := 98 - modulo
	formatCheckDigits(checkDigits, checkBuf)
}

func checksumBytes(iban string) [2]byte {
	var checkBuf [2]byte
	calculateCheckDigits(iban, checkBuf[:])
	return checkBuf
}

// checksum calculates and returns the IBAN check digits as a string.
func checksum(iban string) string {
	checkBuf := checksumBytes(iban)
	return string(checkBuf[:])
}

func normalizeRearrangedIBAN[T ~string | ~[]byte](iban T, buf []byte) int {
	pos := 0
	for i := 4; i < len(iban); i++ {
		pos = normalizeIBANChar(iban[i], buf, pos)
	}
	for i := range 2 {
		pos = normalizeIBANChar(iban[i], buf, pos)
	}
	buf[pos] = '0'
	buf[pos+1] = '0'
	return pos + 2
}

func normalizeIBANChar(c byte, buf []byte, pos int) int {
	switch {
	case c >= 'A' && c <= 'Z':
		value := c - 'A' + 10
		buf[pos] = '0' + value/10
		buf[pos+1] = '0' + value%10
		return pos + 2
	case c >= '0' && c <= '9':
		buf[pos] = c
		return pos + 1
	default:
		return pos
	}
}

func checksumModulo97(digits []byte, length int) int {
	remainder := uint64(0)
	for i := 0; i < length; i += 12 {
		chunkSize := min(12, length-i)
		chunk := uint64(0)
		remainderMultiplier := uint64(1)
		for j := range chunkSize {
			digit := uint64(digits[i+chunkSize-1-j] - '0')
			chunk += digit * remainderMultiplier
			remainderMultiplier *= 10
		}
		remainder = (remainder*remainderMultiplier + chunk) % 97
	}
	return int(remainder)
}

// formatCheckDigits formats check digits with IBAN rules: 0 as 99, 1 as 98, others as-is.
func formatCheckDigits(checkDigits int, buf []byte) {
	switch {
	case checkDigits == 0:
		buf[0], buf[1] = '9', '9'
	case checkDigits == 1:
		buf[0], buf[1] = '9', '8'
	case checkDigits < 10:
		buf[0], buf[1] = '0', decimalDigits[checkDigits]
	default:
		buf[0], buf[1] = decimalDigits[checkDigits/10], decimalDigits[checkDigits%10]
	}
}
