package iban

import (
	"unsafe"
)

// ReplaceChecksum returns the IBAN with corrected check digits.
func ReplaceChecksum(iban string) (string, error) {
	if len(iban) < 4 {
		return "", &ErrValidationLength{Expected: 4, Actual: len(iban)}
	}

	result := []byte(iban)
	calculateCheckDigits(iban, result[2:4])

	return unsafe.String(unsafe.SliceData(result), len(result)), nil
}

// calculateCheckDigits calculates IBAN check digits and writes them to checkBuf.
func calculateCheckDigits(iban string, checkBuf []byte) {
	normalizedBuf := make([]byte, len(iban)*2)
	normalizedLen := normalizeRearrangedIBAN(iban, normalizedBuf)
	normalizedBuf = normalizedBuf[:normalizedLen]

	modulo := mod9710(normalizedBuf, normalizedLen)
	checkDigits := 98 - modulo
	formatCheckDigits(checkDigits, checkBuf)
}

// checksum calculates and returns the IBAN check digits as a string.
func checksum(iban string) string {
	var checkBuf [2]byte
	calculateCheckDigits(iban, checkBuf[:])
	return string(checkBuf[:])
}

// normalizeChar converts a character to its numeric representation (A-Z:10-35, 0-9:0-9) and returns new position.
func normalizeChar(c byte, buf []byte, pos int) int {
	if c >= 'A' && c <= 'Z' {
		val := c - 'A' + 10
		buf[pos] = '0' + val/10
		buf[pos+1] = '0' + val%10
		return pos + 2
	} else if c >= '0' && c <= '9' {
		buf[pos] = c
		return pos + 1
	}
	return pos
}

// normalizeRearrangedIBAN rearranges IBAN to "bban+countrycode+00" format and converts to digits.
func normalizeRearrangedIBAN(iban string, buf []byte) int {
	pos := 0

	for i := 4; i < len(iban); i++ {
		pos = normalizeChar(iban[i], buf, pos)
	}

	for i := range 2 {
		pos = normalizeChar(iban[i], buf, pos)
	}

	buf[pos] = '0'
	buf[pos+1] = '0'
	pos += 2

	return pos
}

// mod9710 calculates modulo 97 for IBAN checksum using 12-digit chunking.
func mod9710(digits []byte, length int) int {
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
		buf[0], buf[1] = '0', '0'+byte(checkDigits)
	default:
		buf[0], buf[1] = '0'+byte(checkDigits/10), '0'+byte(checkDigits%10)
	}
}
