package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"

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
				&rule.StaticRule{
					StartPosition: 0,
					Value:         "GB",
				},
				&rule.RangeRule{
					StartPosition: 2,
					Length:        2,
					Format:        rule.Digit,
				},
				&rule.RangeRule{
					StartPosition: 4,
					Length:        4,
					Format:        rule.UpperCaseLetters,
				},
				&rule.RangeRule{
					StartPosition: 8,
					Length:        6,
					Format:        rule.Digit,
				},
				&rule.RangeRule{
					StartPosition: 14,
					Length:        8,
					Format:        rule.Digit,
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
		"pt": {
			input: "PT2!n4!n4!n11!n2!n",
			want: []rule.Rule{
				&rule.StaticRule{
					StartPosition: 0,
					Value:         "PT",
				},
				&rule.RangeRule{
					StartPosition: 2,
					Length:        2,
					Format:        rule.Digit,
				},
				&rule.RangeRule{
					StartPosition: 4,
					Length:        4,
					Format:        rule.Digit,
				},
				&rule.RangeRule{
					StartPosition: 8,
					Length:        4,
					Format:        rule.Digit,
				},
				&rule.RangeRule{
					StartPosition: 12,
					Length:        11,
					Format:        rule.Digit,
				},
				&rule.RangeRule{
					StartPosition: 23,
					Length:        2,
					Format:        rule.Digit,
				},
			},
			wantErr: false,
		},
	}

	for tc, tt := range tests {
		t.Run(tc, func(t *testing.T) {
			p := New(tt.input)
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
