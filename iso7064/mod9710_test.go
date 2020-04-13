package iso7064

import "testing"

func TestMod9710(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "3214282912345698765432161182",
			want:  1,
		},
		{
			input: "100100100987654321131400",
			want:  69,
		},
		{
			input: "36155444216779151",
			want:  81,
		},
		{
			input: "77277287827223785",
			want:  35,
		},
		{
			input: "2218132140051512345674161100",
			want:  72,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := Mod9710(tt.input); got != tt.want {
				t.Errorf("Mod9710() = %v, want %v", got, tt.want)
			}
		})
	}
}
