package lexer

import (
	"strings"

	"github.com/jacoelho/banking/ascii"
	"github.com/jacoelho/banking/registry/token"
)

const (
	eof rune = 0
)

type Lexer struct {
	input *strings.Reader
}

func New(input string) *Lexer {
	return &Lexer{
		input: strings.NewReader(input),
	}
}

func (l *Lexer) read() rune {
	ch, _, err := l.input.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (l *Lexer) unread() {
	_ = l.input.UnreadRune()
}

func (l *Lexer) scanUpperCaseLetters() token.Token {
	var (
		b strings.Builder
	)

	for {
		r := l.read()

		if r == eof {
			break
		}

		if !ascii.IsUpperCaseLetter(r) {
			l.unread()
			break
		}

		b.WriteRune(r)
	}

	return token.Token{
		Type:    token.STRING,
		Literal: b.String(),
	}
}

func (l *Lexer) scanDigit() token.Token {
	var (
		b strings.Builder
	)

	for {
		r := l.read()

		if r == eof {
			break
		}

		if !ascii.IsDigit(r) {
			l.unread()
			break
		}

		b.WriteRune(r)
	}

	return token.Token{
		Type:    token.INTEGER,
		Literal: b.String(),
	}
}

func (l *Lexer) scanSymbol() token.Token {
	r := l.read()

	return token.Token{
		Type:    token.SYMBOL,
		Literal: string(r),
	}
}

func (l *Lexer) Scan() token.Token {
	switch ch := l.read(); {
	case ascii.IsDigit(ch):
		l.unread()
		return l.scanDigit()
	case ascii.IsUpperCaseLetter(ch):
		l.unread()
		return l.scanUpperCaseLetters()
	case ascii.IsLowerCase(ch):
		l.unread()
		return l.scanSymbol()
	case ch == '!':
		return token.Token{
			Type:    token.BANG,
			Literal: "!",
		}
	case ch == eof:
		return token.Token{
			Type:    token.EOF,
			Literal: "",
		}
	default:
		return token.Token{
			Type:    token.ILLEGAL,
			Literal: "",
		}
	}
}
