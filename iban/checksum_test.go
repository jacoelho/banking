package iban

import (
	"testing"
)

func benchmarkIBANChecksum(b *testing.B, input string) {
	for i := 0; i < b.N; i++ {
		Checksum(input)
	}
}

func BenchmarkIBANChecksumAL(b *testing.B) { benchmarkIBANChecksum(b, "AL47212110090000000235698741") }
func BenchmarkIBANChecksumGB(b *testing.B) { benchmarkIBANChecksum(b, "GB26MIDL40051512345674") }
