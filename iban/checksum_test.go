package iban

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	tests := []struct {
		iban string
		want string
	}{
		{iban: "PT50000201231234567890154", want: "50"},
		{iban: "GB39BARC20040418131256", want: "39"},
		{iban: "GB08BARC20040141721778", want: "08"},
		{iban: "GB75BARC20031847335253", want: "75"},
		{iban: "GB39BARC20039557137528", want: "39"},
		{iban: "GB83BARC20038474225535", want: "83"},
		{iban: "IQ98NBIQ850123456789012", want: "98"},
		{iban: "BR9700360305000010009795493P1", want: "97"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.iban, func(t *testing.T) {
			t.Parallel()
			if got := checksum(tt.iban); got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceChecksum(t *testing.T) {
	tests := []struct {
		iban string
		want string
	}{
		{
			iban: "GB00BUKB20201555555555",
			want: "GB33BUKB20201555555555",
		},
		{
			iban: "DE00512108001245126199",
			want: "DE75512108001245126199",
		},
		{
			iban: "FR0130006000011234567890189",
			want: "FR7630006000011234567890189",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.iban, func(t *testing.T) {
			t.Parallel()
			if got, _ := ReplaceChecksum(tt.iban); got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkIBANChecksum(b *testing.B, input string) {
	for i := 0; i < b.N; i++ {
		checksum(input)
	}
}

func BenchmarkIBANChecksumAL(b *testing.B) { benchmarkIBANChecksum(b, "AL47212110090000000235698741") }
func BenchmarkIBANChecksumGB(b *testing.B) { benchmarkIBANChecksum(b, "GB26MIDL40051512345674") }
