package ast

import "fmt"

type Rule interface {
	String() string
}

type Static struct {
	StartPosition int
	Value         string
}

func (s *Static) String() string {
	return fmt.Sprintf("static: '%s' on position %d", s.Value, s.StartPosition)
}

//go:generate go run golang.org/x/tools/cmd/stringer -type=RangeType
type RangeType int

const (
	Digit RangeType = iota
	UpperCaseLetters
	AlphaNumeric
)

type RangeRule struct {
	StartPosition int
	Length        int
	Type          RangeType
}

func (r *RangeRule) String() string {
	return fmt.Sprintf("range: start position: %d length: %d, type: %s", r.StartPosition, r.Length, r.Type)
}
