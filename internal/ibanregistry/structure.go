package ibanregistry

import (
	"fmt"
	"strconv"
)

// ParseStructure parses a SWIFT IBAN/BBAN structure expression.
func ParseStructure(input string) (Structure, error) {
	var rules []Rule
	pos := 0

	for i := 0; i < len(input); {
		switch c := input[i]; {
		case c >= 'A' && c <= 'Z':
			start := i
			for i < len(input) && input[i] >= 'A' && input[i] <= 'Z' {
				i++
			}
			value := input[start:i]
			rules = append(rules, StaticRule{StartPosition: pos, Value: value})
			pos += len(value)

		case c >= '0' && c <= '9':
			start := i
			for i < len(input) && input[i] >= '0' && input[i] <= '9' {
				i++
			}
			length, err := strconv.Atoi(input[start:i])
			if err != nil {
				return Structure{}, fmt.Errorf("invalid range length %q: %w", input[start:i], err)
			}
			if length <= 0 {
				return Structure{}, fmt.Errorf("range length must be positive: %d", length)
			}
			if i >= len(input) || input[i] != '!' {
				return Structure{}, fmt.Errorf("expected ! after range length at byte %d", i)
			}
			i++
			if i >= len(input) {
				return Structure{}, fmt.Errorf("expected range type at byte %d", i)
			}
			format, err := parseFormat(input[i])
			if err != nil {
				return Structure{}, err
			}
			i++
			rules = append(rules, RangeRule{StartPosition: pos, Length: length, Format: format})
			pos += length

		default:
			return Structure{}, fmt.Errorf("unexpected byte %q at byte %d", c, i)
		}
	}

	return Structure{Rules: rules, Length: pos}, nil
}

func parseFormat(c byte) (Format, error) {
	switch c {
	case 'a':
		return UpperCaseLetters, nil
	case 'n':
		return Digit, nil
	case 'c':
		return AlphaNumeric, nil
	default:
		return 0, fmt.Errorf("unknown range type %q", c)
	}
}

func consolidateRules(rules []Rule) []Rule {
	if len(rules) <= 1 {
		return rules
	}

	result := make([]Rule, 1, len(rules))
	result[0] = rules[0]
	for _, current := range rules[1:] {
		last := result[len(result)-1]
		if last.End() != current.Start() {
			result = append(result, current)
			continue
		}
		switch l := last.(type) {
		case StaticRule:
			if r, ok := current.(StaticRule); ok {
				result[len(result)-1] = StaticRule{StartPosition: l.StartPosition, Value: l.Value + r.Value}
				continue
			}
		case RangeRule:
			if r, ok := current.(RangeRule); ok && l.Format == r.Format {
				result[len(result)-1] = RangeRule{StartPosition: l.StartPosition, Length: l.Length + r.Length, Format: l.Format}
				continue
			}
		}
		result = append(result, current)
	}
	return result
}

func consolidateIBANRules(rules []Rule) []Rule {
	if len(rules) <= 2 {
		return rules
	}
	result := make([]Rule, 0, len(rules))
	result = append(result, rules[0], rules[1])
	return append(result, consolidateRules(rules[2:])...)
}

func structuresEqual(left, right Structure) bool {
	if left.Length != right.Length || len(left.Rules) != len(right.Rules) {
		return false
	}
	for i := range left.Rules {
		if !rulesEqual(left.Rules[i], right.Rules[i]) {
			return false
		}
	}
	return true
}

func rulesEqual(left, right Rule) bool {
	switch l := left.(type) {
	case StaticRule:
		r, ok := right.(StaticRule)
		return ok && l.StartPosition == r.StartPosition && l.Value == r.Value
	case RangeRule:
		r, ok := right.(RangeRule)
		return ok && l.StartPosition == r.StartPosition && l.Length == r.Length && l.Format == r.Format
	default:
		return false
	}
}

func ibanBBANStructure(iban Structure) (Structure, error) {
	var rules []Rule
	for _, rule := range iban.Rules {
		if rule.End() <= 4 {
			continue
		}
		if rule.Start() < 4 {
			return Structure{}, fmt.Errorf("rule crosses IBAN check digit boundary")
		}
		switch r := rule.(type) {
		case StaticRule:
			r.StartPosition -= 4
			rules = append(rules, r)
		case RangeRule:
			r.StartPosition -= 4
			rules = append(rules, r)
		default:
			return Structure{}, fmt.Errorf("unknown rule type %T", rule)
		}
	}
	return Structure{Rules: rules, Length: iban.Length - 4}, nil
}
