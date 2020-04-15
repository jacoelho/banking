package iban

import "testing"

func Test_chunkString(t *testing.T) {
	tests := []struct {
		input     string
		chunkSize int
		sep       rune
		want      string
	}{
		{
			input:     "12345",
			chunkSize: 1,
			sep:       ' ',
			want:      "1 2 3 4 5",
		},
		{
			input:     "12345",
			chunkSize: 2,
			sep:       ' ',
			want:      "12 34 5",
		},
		{
			input:     "12345",
			chunkSize: 3,
			sep:       ' ',
			want:      "123 45",
		},
		{
			input:     "12345",
			chunkSize: 4,
			sep:       ' ',
			want:      "1234 5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := chunkString(tt.input, tt.chunkSize, tt.sep); got != tt.want {
				t.Errorf("chunkString() = %v, want %v", got, tt.want)
			}
		})
	}
}
