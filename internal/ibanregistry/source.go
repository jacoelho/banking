package ibanregistry

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// Encoding names supported registry source encodings.
type Encoding string

const (
	EncodingUTF8   Encoding = "utf8"
	EncodingLatin1 Encoding = "latin1"
	EncodingAuto   Encoding = "auto"
)

// LoadFile loads and decodes a registry source file.
func LoadFile(path string, encoding Encoding) (string, Encoding, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", "", err
	}
	return Decode(data, encoding)
}

// Decode decodes registry bytes and normalizes line endings to LF.
func Decode(data []byte, encoding Encoding) (string, Encoding, error) {
	var decoded string
	var used Encoding

	switch encoding {
	case EncodingUTF8:
		if !utf8.Valid(data) {
			return "", "", fmt.Errorf("registry source is not valid UTF-8")
		}
		decoded = string(data)
		used = EncodingUTF8
	case EncodingLatin1:
		decoded = decodeLatin1(data)
		used = EncodingLatin1
	case EncodingAuto:
		if utf8.Valid(data) {
			decoded = string(data)
			used = EncodingUTF8
		} else {
			decoded = decodeLatin1(data)
			used = EncodingLatin1
		}
	default:
		return "", "", fmt.Errorf("unsupported encoding %q", encoding)
	}

	return normalizeLineEndings(decoded), used, nil
}

func decodeLatin1(data []byte) string {
	var b strings.Builder
	b.Grow(len(data))
	for _, c := range data {
		b.WriteRune(rune(c))
	}
	return b.String()
}

func normalizeLineEndings(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	return strings.ReplaceAll(s, "\r", "\n")
}

// NormalizedBytes returns a stable UTF-8 representation for decoded source.
func NormalizedBytes(source string) []byte {
	return bytes.TrimRight([]byte(normalizeLineEndings(source)), "\n")
}
