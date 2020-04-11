package lexer

import (
	"strings"

	"github.com/jacoelho/iban/registry/token"
)

const (
	eof rune = 0
)

type Lexer struct {
	input    *strings.Reader
	position int
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
	l.position++
	return ch
}

func (l *Lexer) unread() {
	_ = l.input.UnreadRune()
	l.position--
}

func (l *Lexer) scanUpperCaseLetters() token.Token {
	var (
		b   strings.Builder
		pos = l.position
	)

	for {
		r := l.read()

		if r == eof {
			break
		}

		if !isUpperCaseLetter(r) {
			l.unread()
			break
		}

		b.WriteRune(r)
	}

	return token.Token{
		Type:     token.STRING,
		Literal:  b.String(),
		Position: pos,
	}
}

func (l *Lexer) scanDigit() token.Token {
	var (
		b   strings.Builder
		pos = l.position
	)

	for {
		r := l.read()

		if r == eof {
			break
		}

		if !isDigit(r) {
			l.unread()
			break
		}

		b.WriteRune(r)
	}

	return token.Token{
		Type:     token.INTEGER,
		Literal:  b.String(),
		Position: pos,
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
	case isDigit(ch):
		l.unread()
		return l.scanDigit()
	case isUpperCaseLetter(ch):
		l.unread()
		return l.scanUpperCaseLetters()
	case isLowerCase(ch):
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

func isLowerCase(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func isUpperCaseLetter(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}
