package ibanregistry

import "fmt"

// Format is the character class used by a registry range rule.
type Format int

const (
	Digit Format = iota
	UpperCaseLetters
	AlphaNumeric
)

func (f Format) String() string {
	switch f {
	case Digit:
		return "digit"
	case UpperCaseLetters:
		return "upper-case"
	case AlphaNumeric:
		return "alpha-numeric"
	default:
		return "unknown"
	}
}

// Rule describes one IBAN or BBAN structure rule.
type Rule interface {
	rule()
	Start() int
	End() int
}

// StaticRule requires a fixed value at StartPosition.
type StaticRule struct {
	StartPosition int
	Value         string
}

func (StaticRule) rule() {}

func (r StaticRule) Start() int { return r.StartPosition }

func (r StaticRule) End() int { return r.StartPosition + len(r.Value) }

// RangeRule requires Length bytes of Format at StartPosition.
type RangeRule struct {
	StartPosition int
	Length        int
	Format        Format
}

func (RangeRule) rule() {}

func (r RangeRule) Start() int { return r.StartPosition }

func (r RangeRule) End() int { return r.StartPosition + r.Length }

// Structure is a parsed IBAN or BBAN registry structure.
type Structure struct {
	Rules  []Rule
	Length int
}

// Span is a half-open [Start, End) slice in BBAN coordinates.
type Span struct {
	Start int
	End   int
}

func (s Span) Validate(limit int) error {
	if s.Start < 0 {
		return fmt.Errorf("start %d is negative", s.Start)
	}
	if s.End <= s.Start {
		return fmt.Errorf("span [%d,%d) is empty or reversed", s.Start, s.End)
	}
	if s.End > limit {
		return fmt.Errorf("span [%d,%d) exceeds length %d", s.Start, s.End, limit)
	}
	return nil
}

// OptionalSpan represents an optional BBAN component span.
type OptionalSpan struct {
	Span    Span
	Present bool
}

func (s OptionalSpan) Validate(limit int) error {
	if !s.Present {
		return nil
	}
	return s.Span.Validate(limit)
}

// CountrySpec is the complete validated model needed to generate IBAN code.
type CountrySpec struct {
	Code          string
	Name          string
	IBAN          Structure
	BBAN          Structure
	BankCode      OptionalSpan
	BranchCode    OptionalSpan
	AccountNumber Span
	IsSEPA        bool
}

// Registry is a validated IBAN registry snapshot.
type Registry struct {
	Countries []CountrySpec
}
