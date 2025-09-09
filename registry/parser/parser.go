package parser

import (
	"fmt"
	"strconv"

	"github.com/jacoelho/banking/registry/lexer"
	"github.com/jacoelho/banking/registry/rule"
	"github.com/jacoelho/banking/registry/token"
)

// Parse parses an IBAN structure string into rules.
// The input should match ISO 13616 grammar, e.g: "3!n3!n8!n2!n"
func Parse(input string) (*ParseResult, error) {
	lex := lexer.New(input)
	return parseStream(lex)
}

// ParseReduced parses an IBAN structure and consolidates adjacent rules of the same type.
func ParseReduced(input string) (*ParseResult, error) {
	result, err := Parse(input)
	if err != nil {
		return nil, err
	}

	return &ParseResult{
		Rules:  consolidateRules(result.Rules),
		Length: result.Length,
	}, nil
}

func parseStream(lex *lexer.Lexer) (*ParseResult, error) {
	var rules []rule.Rule
	pos := 0

	for {
		tok := lex.Scan()
		if tok.Type == token.EOF {
			break
		}

		r, newPos, err := parseToken(tok, lex, pos)
		if err != nil {
			return nil, newParseError(pos, tok, "", err.Error())
		}

		rules = append(rules, r)
		pos = newPos
	}

	return &ParseResult{
		Rules:  rules,
		Length: pos,
	}, nil
}

func parseToken(tok token.Token, lex *lexer.Lexer, pos int) (rule.Rule, int, error) {
	switch tok.Type {
	case token.STRING:
		return parseStatic(tok, pos)
	case token.INTEGER:
		// Try to parse as range rule (INTEGER ! SYMBOL)
		return parseRange(tok, lex, pos)
	default:
		return nil, pos, fmt.Errorf("unexpected token type %s", tok.String())
	}
}

func parseStatic(tok token.Token, pos int) (rule.Rule, int, error) {
	return &rule.StaticRule{
		StartPosition: pos,
		Value:         tok.Literal,
	}, pos + len(tok.Literal), nil
}

func parseRange(lengthTok token.Token, lex *lexer.Lexer, pos int) (rule.Rule, int, error) {
	length, err := strconv.Atoi(lengthTok.Literal)
	if err != nil {
		return nil, pos, fmt.Errorf("invalid range length: %w", err)
	}

	bangTok := lex.Scan()
	if bangTok.Type != token.BANG {
		return nil, pos, fmt.Errorf("expected '!' after range length, got %s", bangTok.String())
	}

	symbolTok := lex.Scan()
	if symbolTok.Type != token.SYMBOL {
		return nil, pos, fmt.Errorf("expected range type symbol, got %s", symbolTok.String())
	}

	rangeType, err := parseRangeType(symbolTok.Literal)
	if err != nil {
		return nil, pos, err
	}

	return &rule.RangeRule{
		StartPosition: pos,
		Length:        length,
		Format:        rangeType,
	}, pos + length, nil
}

func parseRangeType(symbol string) (rule.RangeRuleType, error) {
	switch symbol {
	case "a":
		return rule.UpperCaseLetters, nil
	case "n":
		return rule.Digit, nil
	case "c":
		return rule.AlphaNumeric, nil
	default:
		return 0, fmt.Errorf("unknown range type: %s", symbol)
	}
}

func consolidateRules(rules []rule.Rule) []rule.Rule {
	if len(rules) <= 1 {
		return rules
	}

	result := make([]rule.Rule, 1, len(rules))
	result[0] = rules[0]

	for _, current := range rules[1:] {
		last := result[len(result)-1]

		if merged := tryMergeRules(last, current); merged != nil {
			result[len(result)-1] = merged
		} else {
			result = append(result, current)
		}
	}

	return result
}

func tryMergeRules(prev, curr rule.Rule) rule.Rule {
	switch p := prev.(type) {
	case *rule.RangeRule:
		if c, ok := curr.(*rule.RangeRule); ok && p.Format == c.Format {
			return &rule.RangeRule{
				StartPosition: p.StartPosition,
				Length:        p.Length + c.Length,
				Format:        p.Format,
			}
		}
	case *rule.StaticRule:
		if c, ok := curr.(*rule.StaticRule); ok {
			return &rule.StaticRule{
				StartPosition: p.StartPosition,
				Value:         p.Value + c.Value,
			}
		}
	}
	return nil
}
