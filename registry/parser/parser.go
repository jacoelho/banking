package parser

import (
	"fmt"
	"strconv"

	"github.com/jacoelho/banking/registry/lexer"
	"github.com/jacoelho/banking/registry/rule"
	"github.com/jacoelho/banking/registry/token"
)

// Parser parses ISO 13616 IBAN grammar into rules
type Parser struct {
	lexer        *lexer.Lexer
	pos          int
	errors       []error
	currentToken token.Token
	peekToken    token.Token
}

// New creates a string Parser
// input should match ISO 13616 grammer, eg: `3!n3!n8!n2!n`
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

func (p *Parser) ReducedParse() ([]rule.Rule, error) {
	rules, err := p.Parse()
	if err != nil {
		return nil, err
	}

	i := 0
	for _, r := range rules {
		if i == 0 {
			rules[i] = r
			i++
			continue
		}

		switch cur := r.(type) {
		case *rule.RangeRule:
			if prev, ok := rules[i-1].(*rule.RangeRule); ok && cur.Format == prev.Format {
				prev.Length += cur.Length
				continue
			}
		case *rule.StaticRule:
			if prev, ok := rules[i-1].(*rule.StaticRule); ok {
				prev.Value += cur.Value
				continue
			}
		}
		rules[i] = r
		i++
	}

	return rules[:i], nil
}

func (p *Parser) Parse() ([]rule.Rule, error) {
	var rules []rule.Rule
	for !p.isCurrentTokenType(token.EOF) {
		r := p.parseRule()
		if r != nil {
			rules = append(rules, r)
		}

		p.nextToken()
	}

	if p.errors != nil {
		return nil, fmt.Errorf("parsing error %v", p.errors)
	}

	return rules, nil
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
