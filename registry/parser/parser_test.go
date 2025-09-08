package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/jacoelho/banking/registry/rule"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		input      string
		want       []rule.Rule
		wantLength int
		wantErr    bool
	}{
		"GB": {
			input:      "GB2!n4!a6!n8!n",
			wantLength: 22,
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
		"pt": {
			input:      "PT2!n4!n4!n11!n2!n",
			wantLength: 25,
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
			result, err := Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != nil {
				if diff := cmp.Diff(tt.want, result.Rules); diff != "" {
					t.Errorf("Parse() rules mismatch (-want +got):\n%s", diff)
				}
				if result.Length != tt.wantLength {
					t.Errorf("Parse() length = %v, want %v", result.Length, tt.wantLength)
				}
			}
		})
	}
}

func TestParseReduced(t *testing.T) {
	tests := map[string]struct {
		input      string
		want       []rule.Rule
		wantLength int
		wantErr    bool
	}{
		"GB": {
			input:      "GB2!n4!a6!n8!n",
			wantLength: 22,
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
					Length:        14,
					Format:        rule.Digit,
				},
			},
		},
		"pt": {
			input:      "PT2!n4!n4!n11!n2!n",
			wantLength: 25,
			want: []rule.Rule{
				&rule.StaticRule{
					StartPosition: 0,
					Value:         "PT",
				},
				&rule.RangeRule{
					StartPosition: 2,
					Length:        23,
					Format:        rule.Digit,
				},
			},
			wantErr: false,
		},
	}

	for tc, tt := range tests {
		t.Run(tc, func(t *testing.T) {
			result, err := ParseReduced(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseReduced() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != nil {
				if diff := cmp.Diff(tt.want, result.Rules); diff != "" {
					t.Errorf("ParseReduced() rules mismatch (-want +got):\n%s", diff)
				}
				if result.Length != tt.wantLength {
					t.Errorf("ParseReduced() length = %v, want %v", result.Length, tt.wantLength)
				}
			}
		})
	}
}
