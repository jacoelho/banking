package rule

import "fmt"

var _ Rule = (*Static)(nil)

type Static struct {
	StartPosition int
	Value         string
}

func (s *Static) Type() string {
	return "Static"
}

func (s *Static) String() string {
	return fmt.Sprintf("%s'%s' on position %d", s.Type(), s.Value, s.StartPosition)
}
