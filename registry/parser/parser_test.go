package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/jacoelho/iban/registry/lexer"
	"github.com/jacoelho/iban/registry/rule"
)

func TestParser_Parse(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    []rule.Rule
		wantErr bool
	}{
		"GB": {
			input: "GB2!n4!a6!n8!n",
			want: []rule.Rule{
				&rule.Static{
					StartPosition: 0,
					Value:         "GB",
				},
				&rule.RangeRule{
					StartPosition: 2,
					Length:        2,
					Type:          rule.Digit,
				},
				&rule.RangeRule{
					StartPosition: 5,
					Length:        4,
					Type:          rule.UpperCaseLetters,
				},
				&rule.RangeRule{
					StartPosition: 8,
					Length:        6,
					Type:          rule.Digit,
				},
				&rule.RangeRule{
					StartPosition: 11,
					Length:        8,
					Type:          rule.Digit,
				},
			},
		},
		"no symbol": {
			input:   "!2!",
			wantErr: true,
		},
		"only digits": {
			input:   "222",
			wantErr: true,
		},
	}

	for tc, tt := range tests {
		t.Run(tc, func(t *testing.T) {
			p := NewParser(lexer.New(tt.input))
			got, err := p.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != nil {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("Parse() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
