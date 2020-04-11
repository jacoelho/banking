package rule

import (
	"fmt"
	"strconv"
	"strings"
)

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

func (r *RangeRule) StartPos() int {
	return r.StartPosition
}

func (r *RangeRule) EndPos() int {
	return r.StartPosition + r.Length
}

func (r *RangeRule) Type() RuleType {
	return RuleRange
}

func (r *RangeRule) String() string {
	return fmt.Sprintf("%d: start position: %d Length: %d, type: %s", r.Type(), r.StartPosition, r.Length, r.Format)
}

func (r *RangeRule) WriteTo(sb *strings.Builder) {
	sb.WriteString("if rangeValue := iban[")
	sb.WriteString(strconv.Itoa(r.StartPosition))
	sb.WriteString(":")
	sb.WriteString(strconv.Itoa(r.StartPosition + r.Length))
	sb.WriteString("]; ")

	switch r.Format {
	case Digit:
		sb.WriteString("!every(rangeValue, isDigit)")
	case AlphaNumeric:
		sb.WriteString("!every(rangeValue, isAlphaNumeric)")
	case UpperCaseLetters:
		sb.WriteString("!every(rangeValue, isUpperCaseLetter)")
	}

	sb.WriteString(" {\n")
	sb.WriteString(`return errors.New("adsa")`)
	sb.WriteString(" }\n")
}
