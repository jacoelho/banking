package iso7064

import "testing"

func TestMod11_2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "0794",
			want:  0,
		},
		{
			input: "079",
			want:  10,
		},
		{
			input: "000000029079593",
			want:  10,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.input, func(t *testing.T) {
			if got := Mod112(tt.input); got != tt.want {
				t.Errorf("Mod11_10() = %v, want %v", got, tt.want)
			}
		})
	}
}
