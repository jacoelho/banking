package rule

import (
	"fmt"
)

var _ Rule = (*StaticRule)(nil)

type StaticRule struct {
	StartPosition int
	Value         string
}

func (s *StaticRule) StartPos() int {
	return s.StartPosition
}

func (s *StaticRule) EndPos() int {
	return s.StartPosition + len(s.Value)
}

func (s *StaticRule) Type() string {
	return "Static"
}

func (s *StaticRule) String() string {
	return fmt.Sprintf("static value rule, pos: %d, expected value: %s", s.StartPosition, s.Value)
}
