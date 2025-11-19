package iban_test

import (
	"testing"

	"github.com/jacoelho/banking/iban"
)

func benchmarkGenerate(b *testing.B, input string) {
	for b.Loop() {
		if _, err := iban.Generate(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIBANGenerateGB(b *testing.B) { benchmarkGenerate(b, "GB") }
func BenchmarkIBANGenerateBR(b *testing.B) { benchmarkGenerate(b, "BR") }
