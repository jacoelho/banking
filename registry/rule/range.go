package rule

import "fmt"

type RangeType int

const (
	Digit RangeType = iota
	UpperCaseLetters
	AlphaNumeric
)

func (r RangeType) String() string {
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
	Format        RangeType
}

func (r *RangeRule) Type() string {
	return "Range"
}

func (r *RangeRule) String() string {
	return fmt.Sprintf("%s: start position: %d length: %d, type: %s", r.Type(), r.StartPosition, r.Length, r.Format)
}
