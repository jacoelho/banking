package iso7064

import (
	"testing"
)

func TestMod731(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "12131295",
			want:  2,
		},
		{
			input: "22102014568",
			want:  5,
		},
		{
			input: "97888932329",
			want:  2,
		},
		{
			input: "978889323291",
			want:  3,
		},
		{
			input: "9788893232912",
			want:  3,
		},
		{
			input: "123456789",
			want:  7,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			if got := Mod731(tt.input); got != tt.want {
				t.Errorf("Mod731() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMod731Encoded(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "12131295",
			want:  "121312952",
		},
		{
			input: "22102014568",
			want:  "221020145685",
		},
		{
			input: "97888932329",
			want:  "978889323292",
		},
		{
			input: "978889323291",
			want:  "9788893232913",
		},
		{
			input: "9788893232912",
			want:  "97888932329123",
		},
		{
			input: "123456789",
			want:  "1234567897",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			if got := Mod731Encoded(tt.input); got != tt.want {
				t.Errorf("Mod731() = %v, want %v", got, tt.want)
			}
		})
	}
}
