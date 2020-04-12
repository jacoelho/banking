package iban

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	tests := []struct {
		iban string
		want string
	}{
		{
			iban: "PT50000201231234567890154",
			want: "50",
		},
	}
	for _, tt := range tests {
		t.Run(tt.iban, func(t *testing.T) {
			if got := Checksum(tt.iban); got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkIBANChecksum(b *testing.B, input string) {
	for i := 0; i < b.N; i++ {
		Checksum(input)
	}
}

func BenchmarkIBANChecksumAL(b *testing.B) { benchmarkIBANChecksum(b, "AL47212110090000000235698741") }
func BenchmarkIBANChecksumGB(b *testing.B) { benchmarkIBANChecksum(b, "GB26MIDL40051512345674") }
