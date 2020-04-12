package iban

import "testing"

func benchmarkGenerate(b *testing.B, input string) {
	for i := 0; i < b.N; i++ {
		if _, err := Generate(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIBANGenerateGB(b *testing.B) { benchmarkGenerate(b, "GB") }
func BenchmarkIBANGenerateBR(b *testing.B) { benchmarkGenerate(b, "BR") }
