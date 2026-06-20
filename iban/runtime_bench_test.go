package iban_test

import (
	"testing"

	"github.com/jacoelho/banking/iban"
)

func BenchmarkIBANValidateGB(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		if err := iban.Validate("GB29NWBK60161331926819"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIBANValidateBR(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		if err := iban.Validate("BR9700360305000010009795493P1"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIBANGetBBANGB(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		if _, err := iban.GetBBAN("GB29NWBK60161331926819"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIBANIsSEPAGB(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		if _, err := iban.IsSEPA("GB29NWBK60161331926819"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIBANIsSEPACountryCodeGB(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		if _, err := iban.IsSEPACountryCode("GB"); err != nil {
			b.Fatal(err)
		}
	}
}
