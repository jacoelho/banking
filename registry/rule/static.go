package rule

import (
	"fmt"
	"strconv"
	"strings"
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

func (s *StaticRule) WriteTo(sb *strings.Builder) {
	sb.WriteString("if staticValue := iban[")
	sb.WriteString(strconv.Itoa(s.StartPosition))
	sb.WriteString(":")
	sb.WriteString(strconv.Itoa(s.StartPosition + len(s.Value)))
	sb.WriteString("]; staticValue != \"")
	sb.WriteString(s.Value)
	sb.WriteString("\" {\n")
	sb.WriteString(`return errors.New("invalid value")`)
	sb.WriteString("\n}")
}
