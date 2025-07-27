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

	for i := 0; i < len(s); i += chunkSize {
		if i > 0 {
			sb.WriteRune(sep)
		}

		end := min(i+chunkSize, len(s))
		sb.WriteString(s[i:end])
	}

	return sb.String()
}
