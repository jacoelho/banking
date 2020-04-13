package iso7064

import "testing"

func TestMod11_10(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "79462",
			want:  3,
		},
		{
			input: "00200667308",
			want:  5,
		},
		{
			input: "0794",
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := Mod1110(tt.input); got != tt.want {
				t.Errorf("Mod1110() = %v, want %v", got, tt.want)
			}
		})
	}
}
