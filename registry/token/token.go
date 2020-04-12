package token

import "fmt"

type ItemType int

const (
	ILLEGAL ItemType = iota
	EOF
	STRING  // GB AL
	INTEGER // 12
	SYMBOL  // a, n, c
	BANG    // !
)

type Token struct {
	Type    ItemType
	Literal string
}

func (t Token) String() string {
	switch t.Type {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case STRING:
		return fmt.Sprintf("string <%s>", t.Literal)
	case INTEGER:
		return fmt.Sprintf("integer <%s>", t.Literal)
	case SYMBOL:
		return fmt.Sprintf("symbol <%s>", t.Literal)
	case BANG:
		return "BANG"
	default:
		return fmt.Sprintf("<%s>", t.Literal)
	}
}
