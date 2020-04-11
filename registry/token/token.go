package token

//go:generate go run golang.org/x/tools/cmd/stringer -type=TokenType
type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	WORD
	INTEGER
	SYMBOL // a, n, c
	BANG   // !
)

type Token struct {
	Type     TokenType
	Literal  string
	Position int
}
