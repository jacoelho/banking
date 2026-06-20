package ascii

import (
	"testing"
)

func TestAppendRandomStringGenerators(t *testing.T) {
	tests := []struct {
		name    string
		fn      func([]byte, uint8) []byte
		length  uint8
		allowed func(byte) bool
	}{
		{
			name:    "digits",
			fn:      AppendRandomDigits,
			length:  32,
			allowed: func(b byte) bool { return b >= '0' && b <= '9' },
		},
		{
			name:    "upper case letters",
			fn:      AppendRandomUpperCaseLetters,
			length:  32,
			allowed: func(b byte) bool { return b >= 'A' && b <= 'Z' },
		},
		{
			name:   "alpha numeric",
			fn:     AppendRandomAlphaNumeric,
			length: 32,
			allowed: func(b byte) bool {
				return b >= '0' && b <= '9' || b >= 'A' && b <= 'Z'
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn([]byte("prefix"), tt.length)
			if string(got[:6]) != "prefix" {
				t.Fatalf("prefix = %q, want prefix", got[:6])
			}
			got = got[6:]
			if len(got) != int(tt.length) {
				t.Fatalf("length = %d, want %d", len(got), tt.length)
			}
			for _, b := range got {
				if !tt.allowed(b) {
					t.Fatalf("generated %q contains disallowed byte %q", got, b)
				}
			}
		})
	}
}

func TestAppendRandomStringGeneratorsHandleZeroLength(t *testing.T) {
	if got := AppendRandomDigits([]byte("prefix"), 0); string(got) != "prefix" {
		t.Fatalf("AppendRandomDigits with zero length = %q, want prefix", got)
	}
}
