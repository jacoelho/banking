package parser

import (
	"fmt"
	"strconv"

	"github.com/jacoelho/iban/registry/lexer"
	"github.com/jacoelho/iban/registry/rule"
	"github.com/jacoelho/iban/registry/token"
)

type Parser struct {
	lexer        *lexer.Lexer
	errors       []error
	currentToken token.Token
	peekToken    token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer: l,
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.Scan()
}

func (p *Parser) isCurrentTokenType(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) isPeekTokenType(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if !p.isPeekTokenType(t) {
		p.errors = append(p.errors, fmt.Errorf("expect next token to be '%s', got '%s'", t, p.peekToken.Type))
		return false
	}

	p.nextToken()
	return true
}

func (p *Parser) Parse() ([]rule.Rule, []error) {
	var rules []rule.Rule
	for !p.isCurrentTokenType(token.EOF) {
		r := p.parseRule()
		if r != nil {
			rules = append(rules, r)
		}

		p.nextToken()
	}

	return rules, p.errors
}

func (p *Parser) parseRule() rule.Rule {
	switch p.currentToken.Type {
	case token.WORD:
		return p.parseStatic()
	case token.INTEGER:
		return p.parseRange()
	default:
		return nil
	}
}

func (p *Parser) parseStatic() *rule.Static {
	return &rule.Static{
		StartPosition: p.currentToken.Position,
		Value:         p.currentToken.Literal,
	}
}

func (p *Parser) parseRange() *rule.RangeRule {
	pos := p.currentToken.Position

	length, err := strconv.Atoi(p.currentToken.Literal)
	if err != nil {
		return nil
	}

	if !p.expectPeek(token.BANG) {
		return nil
	}

	if !p.expectPeek(token.SYMBOL) {
		return nil
	}

	var rangeType rule.RangeType
	switch p.currentToken.Literal {
	case "a":
		rangeType = rule.UpperCaseLetters
	case "n":
		rangeType = rule.Digit
	case "c":
		rangeType = rule.AlphaNumeric
	default:
		return nil
	}

	return &rule.RangeRule{
		StartPosition: pos,
		Length:        length,
		Type:          rangeType,
	}
}
