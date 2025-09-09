package parser

import (
	"fmt"

	"github.com/jacoelho/banking/registry/rule"
	"github.com/jacoelho/banking/registry/token"
)

// ParseResult contains the result of parsing an IBAN structure
type ParseResult struct {
	Rules  []rule.Rule // Parsed rules
	Length int         // Total length of the IBAN structure
}

// ParseError provides detailed error information for parsing failures
type ParseError struct {
	Position int         // Character position where error occurred
	Token    token.Token // Token that caused the error
	Expected string      // What was expected
	Context  string      // Error context/message
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error at position %d: %s (got %s, expected %s)",
		e.Position, e.Context, e.Token.String(), e.Expected)
}

// newParseError creates a ParseError with context
func newParseError(pos int, tok token.Token, expected, context string) *ParseError {
	return &ParseError{
		Position: pos,
		Token:    tok,
		Expected: expected,
		Context:  context,
	}
}
