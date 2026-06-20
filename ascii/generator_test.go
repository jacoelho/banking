package ascii

import (
	"testing"
)

func TestAppendRandomStringGenerators(t *testing.T) {
	tests := []struct {
		name    string
		fn      func([]byte, int) []byte
		length  int
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
			if len(got) != tt.length {
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

func TestRandomStringGenerators(t *testing.T) {
	tests := []struct {
		name    string
		fn      func(int) string
		allowed func(byte) bool
	}{
		{
			name:    "digits",
			fn:      RandomDigits,
			allowed: func(b byte) bool { return b >= '0' && b <= '9' },
		},
		{
			name:    "upper case letters",
			fn:      RandomUpperCaseLetters,
			allowed: func(b byte) bool { return b >= 'A' && b <= 'Z' },
		},
		{
			name: "alpha numeric",
			fn:   RandomAlphaNumeric,
			allowed: func(b byte) bool {
				return b >= '0' && b <= '9' || b >= 'A' && b <= 'Z'
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn(32)
			if len(got) != 32 {
				t.Fatalf("length = %d, want 32", len(got))
			}
			for i := range len(got) {
				if !tt.allowed(got[i]) {
					t.Fatalf("generated %q contains disallowed byte %q", got, got[i])
				}
			}
		})
	}
}

func TestRandomStringGeneratorsHandleNonPositiveLength(t *testing.T) {
	if got := RandomDigits(0); got != "" {
		t.Fatalf("RandomDigits(0) = %q, want empty", got)
	}
	if got := RandomDigits(-1); got != "" {
		t.Fatalf("RandomDigits(-1) = %q, want empty", got)
	}
	if got := AppendRandomDigits([]byte("prefix"), -1); string(got) != "prefix" {
		t.Fatalf("AppendRandomDigits with negative length = %q, want prefix", got)
	}
}
