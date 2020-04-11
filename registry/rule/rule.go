package rule

import (
	"fmt"
)

type Rule interface {
	Type() RuleType
	String() string
	StartPos() int
	EndPos() int
}

type RuleType int

const (
	RuleStatic RuleType = iota
	RuleRange
)

func (r RuleType) String() string {
	switch r {
	case RuleStatic:
		return "Static"
	case RuleRange:
		return "Range"
	}

	return fmt.Sprintf("unknown <%d>", r)
}
