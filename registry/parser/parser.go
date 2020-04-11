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
	pos          int
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
	currentPost := p.pos

	p.pos += len(p.currentToken.Literal)
	return &rule.StaticRule{
		StartPosition: currentPost,
		Value:         p.currentToken.Literal,
	}
}

func (p *Parser) parseRange() *rule.RangeRule {
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

	currentPost := p.pos
	p.pos += length

	return &rule.RangeRule{
		StartPosition: currentPost,
		Length:        length,
		Format:        rangeType,
	}
}
