package iban

import (
	"errors"
	"testing"
)

func TestChecksum(t *testing.T) {
	t.Parallel()

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
		{iban: "DK0206715394960066", want: "02"},
		{iban: "NI45BAPR00000013000003558124", want: "45"},
	}
	for _, tt := range tests {
		t.Run(tt.iban, func(t *testing.T) {
			t.Parallel()
			if got := checksum(tt.iban); got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceChecksum(t *testing.T) {
	t.Parallel()

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
		t.Run(tt.iban, func(t *testing.T) {
			t.Parallel()
			got, err := ReplaceChecksum(tt.iban)
			if err != nil {
				t.Fatalf("ReplaceChecksum() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceChecksumBytesMutatesInput(t *testing.T) {
	iban := []byte("GB00BUKB20201555555555")

	got := replaceChecksumBytes(iban)

	if got != "GB33BUKB20201555555555" {
		t.Fatalf("replaceChecksumBytes() = %q, want GB33BUKB20201555555555", got)
	}
	if string(iban) != got {
		t.Fatalf("replaceChecksumBytes() did not mutate input: %q", iban)
	}
}

func TestFormatCheckDigits(t *testing.T) {
	tests := []struct {
		name        string
		checkDigits int
		want        string
	}{
		{name: "zero", checkDigits: 0, want: "99"},
		{name: "one", checkDigits: 1, want: "98"},
		{name: "single digit", checkDigits: 7, want: "07"},
		{name: "two digits", checkDigits: 42, want: "42"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got [2]byte
			formatCheckDigits(tt.checkDigits, got[:])
			if string(got[:]) != tt.want {
				t.Fatalf("formatCheckDigits(%d) = %q, want %q", tt.checkDigits, got, tt.want)
			}
		})
	}
}

func TestReplaceChecksumRejectsInvalidStructure(t *testing.T) {
	tests := []struct {
		name string
		iban string
		want ValidationError
	}{
		{
			name: "too short",
			iban: "GB1",
			want: ValidationError{
				Reason:         ReasonInvalidLength,
				Length:         4,
				ExpectedLength: 4,
				ActualLength:   3,
			},
		},
		{
			name: "non-digit check digits",
			iban: "GBXXNWBK60161331926819",
			want: ValidationError{
				Reason:   ReasonInvalidCharacters,
				Position: 2,
				Length:   2,
				Expected: CharClassDigit,
				Actual:   "XX",
			},
		},
		{
			name: "invalid BBAN character",
			iban: "GB29NWBK6016133192681X",
			want: ValidationError{
				Reason:   ReasonInvalidCharacters,
				Position: 8,
				Length:   14,
				Expected: CharClassDigit,
				Actual:   "6016133192681X",
			},
		},
		{
			name: "lowercase alphanumeric character",
			iban: "FR0020041010050500013m02606",
			want: ValidationError{
				Reason:   ReasonInvalidCharacters,
				Position: 14,
				Length:   11,
				Expected: CharClassUpperAlphaNumeric,
				Actual:   "0500013m026",
			},
		},
		{
			name: "unsupported country",
			iban: "ZZ29NWBK60161331926819",
			want: ValidationError{
				Reason:   ReasonUnsupportedCountry,
				Position: 0,
				Length:   2,
				Expected: CharClassUpperAlpha,
				Actual:   "ZZ",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := ReplaceChecksum(tt.iban); err == nil {
				t.Fatalf("ReplaceChecksum() error = nil, got %q", got)
			} else {
				assertValidationError(t, err, tt.want)
			}
		})
	}
}

func assertValidationError(t *testing.T, err error, want ValidationError) {
	t.Helper()

	if !errors.Is(err, ErrInvalidIBAN) {
		t.Fatalf("errors.Is(err, ErrInvalidIBAN) = false, want true")
	}

	var got *ValidationError
	if !errors.As(err, &got) {
		t.Fatalf("errors.As(err, *ValidationError) = false, want true")
	}

	if got.Reason != want.Reason ||
		got.Position != want.Position ||
		got.Length != want.Length ||
		got.Expected != want.Expected ||
		got.Actual != want.Actual ||
		got.ExpectedValue != want.ExpectedValue ||
		got.ExpectedLength != want.ExpectedLength ||
		got.ActualLength != want.ActualLength {
		t.Fatalf("ValidationError = %+v, want %+v", got, want)
	}
}

func benchmarkIBANChecksum(b *testing.B, input string) {
	b.Helper()

	for b.Loop() {
		checksum(input)
	}
}

func BenchmarkIBANChecksumAL(b *testing.B) { benchmarkIBANChecksum(b, "AL47212110090000000235698741") }
func BenchmarkIBANChecksumGB(b *testing.B) { benchmarkIBANChecksum(b, "GB26MIDL40051512345674") }

func benchmarkCalculateCheckDigits(b *testing.B, input string) {
	b.Helper()
	var checkBuf [2]byte

	b.ReportAllocs()
	for b.Loop() {
		calculateCheckDigits(input, checkBuf[:])
	}
}

func BenchmarkCalculateCheckDigitsAL(b *testing.B) {
	benchmarkCalculateCheckDigits(b, "AL47212110090000000235698741")
}

func BenchmarkCalculateCheckDigitsBR(b *testing.B) {
	benchmarkCalculateCheckDigits(b, "BR9700360305000010009795493P1")
}

func BenchmarkCalculateCheckDigitsGB(b *testing.B) {
	benchmarkCalculateCheckDigits(b, "GB26MIDL40051512345674")
}

func BenchmarkCalculateCheckDigitsMaxLetters(b *testing.B) {
	benchmarkCalculateCheckDigits(b, "ZZ00ABCDEFGHIJKLMNOPQRSTUVWXYZABCD")
}

func BenchmarkCalculateCheckDigitsLetters(b *testing.B) {
	benchmarkCalculateCheckDigits(b, "MT84MALT011000012345MTLCAST001S")
}

func BenchmarkReplaceChecksum(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		if _, err := ReplaceChecksum("GB00BUKB20201555555555"); err != nil {
			b.Fatal(err)
		}
	}
}
