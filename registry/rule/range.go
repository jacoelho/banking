package rule

import (
	"fmt"
)

type RangeRuleType int

const (
	Digit RangeRuleType = iota
	UpperCaseLetters
	AlphaNumeric
)

func (r RangeRuleType) String() string {
	switch r {
	case Digit:
		return "Digit"
	case UpperCaseLetters:
		return "UpperCaseLetters"
	case AlphaNumeric:
		return "AlphaNumeric"
	}
	return string(r)
}

var _ Rule = (*RangeRule)(nil)

type RangeRule struct {
	StartPosition int
	Length        int
	Format        RangeRuleType
}

func (r *RangeRule) StartPos() int {
	return r.StartPosition
}

func (r *RangeRule) EndPos() int {
	return r.StartPosition + r.Length
}

func (r *RangeRule) Type() string {
	return "Range"
}

func (r *RangeRule) String() string {
	return fmt.Sprintf("range rule, start pos: %d, length: %d, expected type %s", r.StartPosition, r.Length, r.Format)
}
