package rule

import (
	"strings"
	"testing"
)

func TestRuleValidation(t *testing.T) {
	tests := []struct {
		name  string
		rules []Rule
	}{
		{
			name: "static",
			rules: []Rule{
				&RangeRule{
					StartPosition: 5,
					Length:        5,
					Format:        AlphaNumeric,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sb = new(strings.Builder)

			for _, v := range tt.rules {
				v.WriteTo(sb)
			}

			t.Fatalf(sb.String())
		})
	}
}
