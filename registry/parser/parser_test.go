package parser

import (
	"testing"

	"github.com/jacoelho/iban/registry/ast"
	"github.com/jacoelho/iban/registry/lexer"

	"github.com/google/go-cmp/cmp"
)

func TestParser_Parse(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    []ast.Rule
		wantErr bool
	}{
		"GB": {
			input: "GB2!n4!a6!n8!n",
			want: []ast.Rule{
				&ast.Static{
					StartPosition: 0,
					Value:         "GB",
				},
				&ast.RangeRule{
					StartPosition: 2,
					Length:        2,
					Type:          ast.Digit,
				},
				&ast.RangeRule{
					StartPosition: 5,
					Length:        4,
					Type:          ast.UpperCaseLetters,
				},
				&ast.RangeRule{
					StartPosition: 8,
					Length:        6,
					Type:          ast.Digit,
				},
				&ast.RangeRule{
					StartPosition: 11,
					Length:        8,
					Type:          ast.Digit,
				},
			},
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
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Parse() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
