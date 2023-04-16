package iban

import "github.com/jacoelho/banking/pool"

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
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}

		if i > 0 {
			_, _ = sb.WriteRune(sep)
		}

		_, _ = sb.Write(runes[i:nn])

	}
	return sb.String()
}
