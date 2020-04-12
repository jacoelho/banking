package lexer

import (
	"reflect"
	"testing"

	"github.com/jacoelho/banking/registry/token"
)

func TestLexer_Scan(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []token.Token
	}{
		"GB": {
			input: "GB2!n4!a6!n8!n",
			want: []token.Token{
				{Type: token.STRING, Literal: "GB"},
				{Type: token.INTEGER, Literal: "2"},
				{Type: token.BANG, Literal: "!"},
				{Type: token.SYMBOL, Literal: "n"},
				{Type: token.INTEGER, Literal: "4"},
				{Type: token.BANG, Literal: "!"},
				{Type: token.SYMBOL, Literal: "a"},
				{Type: token.INTEGER, Literal: "6"},
				{Type: token.BANG, Literal: "!"},
				{Type: token.SYMBOL, Literal: "n"},
				{Type: token.INTEGER, Literal: "8"},
				{Type: token.BANG, Literal: "!"},
				{Type: token.SYMBOL, Literal: "n"},
			},
		},
	}
	for tc, tt := range tests {
		t.Run(tc, func(t *testing.T) {
			l := New(tt.input)

			for _, want := range tt.want {
				got := l.Scan()

				if !reflect.DeepEqual(got, want) {
					t.Fatalf("Scan() = %v, want %v", got, want)
				}
			}

			if got := l.Scan(); got.Type != token.EOF {
				t.Errorf("expected EOF, got %v", got)
			}
		})
	}
}
