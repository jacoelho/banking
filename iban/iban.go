package iban

// PaperFormat returns iban in paper format
// Follows EBS  204 - 5 IBAN Format  2) Paper Format
func PaperFormat(iban string) string {
	return chunkString(iban, 4, ' ')
}

func chunkString(s string, chunkSize int, sep byte) string {
	if s == "" {
		return ""
	}

	out := make([]byte, 0, len(s)+len(s)/chunkSize)

	for i := 0; i < len(s); i += chunkSize {
		if i > 0 {
			out = append(out, sep)
		}

		end := min(i+chunkSize, len(s))
		out = append(out, s[i:end]...)
	}

	return string(out)
}
