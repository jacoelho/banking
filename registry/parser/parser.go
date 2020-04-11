package parser

import (
	"fmt"
	"strconv"

	"github.com/jacoelho/iban/registry/ast"
	"github.com/jacoelho/iban/registry/lexer"
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

func (p *Parser) Parse() ([]ast.Rule, error) {
	var rules []ast.Rule
	for !p.isCurrentTokenType(token.EOF) {
		rule := p.parseRule()
		if rule != nil {
			rules = append(rules, rule)
		}

		p.nextToken()
	}

	return rules, nil
}

func (p *Parser) parseRule() ast.Rule {
	switch p.currentToken.Type {
	case token.WORD:
		return p.parseStatic()
	case token.INTEGER:
		return p.parseRange()
	default:
		return nil
	}
}

func (p *Parser) parseStatic() *ast.Static {
	return &ast.Static{
		StartPosition: p.currentToken.Position,
		Value:         p.currentToken.Literal,
	}
}

func (p *Parser) parseRange() *ast.RangeRule {
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

	var rangeType ast.RangeType
	switch p.currentToken.Literal {
	case "a":
		rangeType = ast.UpperCaseLetters
	case "n":
		rangeType = ast.Digit
	case "c":
		rangeType = ast.AlphaNumeric
	default:
		return nil
	}

	return &ast.RangeRule{
		StartPosition: pos,
		Length:        length,
		Type:          rangeType,
	}
}
