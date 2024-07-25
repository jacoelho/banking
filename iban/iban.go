package iban

import (
	"github.com/jacoelho/banking/pool"
)

//go:generate banking-registry -registry-file ../docs/registry.yml -dst-directory .

// PaperFormat returns iban in paper format
// Follows EBS  204 - 5 IBAN Format  2) Paper Format
func PaperFormat(iban string) string {
	return chunkString(iban, 4, ' ')
}

func chunkString(s string, chunkSize int, sep rune) string {
	if len(s) == 0 {
		return ""
	}

	sb := pool.BytesPool.Get()
	defer sb.Free()

	runes := []byte(s)

	for i := 0; i < len(runes); i += chunkSize {
		if i > 0 {
			sb.WriteRune(sep)
		}

		end := i + chunkSize
		if end > len(runes) {
			end = len(runes)
		}

		sb.Write(runes[i:end])
	}

	return sb.String()
}
