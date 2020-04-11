package parser

import (
	"fmt"
	"strconv"

	"github.com/jacoelho/iban/registry/token"

	"github.com/jacoelho/iban/registry/lexer"
	"github.com/jacoelho/iban/registry/rule"
)

type Parser struct {
	lexer        *lexer.Lexer
	errors       []error
	currentToken token.Token
	peekToken    token.Token
}

func New(input string) *Parser {
	p := &Parser{
		lexer: lexer.New(input),
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.Scan()
}

func (p *Parser) isCurrentTokenType(t token.ItemType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) isPeekTokenType(t token.ItemType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.ItemType) bool {
	if !p.isPeekTokenType(t) {
		p.errors = append(p.errors, fmt.Errorf("expect next token to be '%v', got '%s'", t, p.peekToken.String()))
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
	case token.STRING:
		return p.parseStatic()
	case token.INTEGER:
		return p.parseRange()
	default:
		return nil
	}
}

func (p *Parser) parseStatic() *rule.StaticRule {
	return &rule.StaticRule{
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

	var rangeType rule.RangeRuleType
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
		Format:        rangeType,
	}
}
