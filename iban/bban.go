package iban

import (
	"github.com/jacoelho/banking/pool"
)

type BBAN struct {
	BBAN             string
	BankCode         string
	BranchCode       string
	NationalChecksum string
	AccountNumber    string
}

func (b BBAN) String() string {
	return chunkString(b.BBAN, 4, ' ')
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
