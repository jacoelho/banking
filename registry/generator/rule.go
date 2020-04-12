package generator

import (
	"fmt"
	"text/template"

	"github.com/jacoelho/banking/registry/rule"
)

func templateFunctions() template.FuncMap {
	return template.FuncMap{
		"rule": func(r rule.Rule, s string) string {
			switch v := r.(type) {
			case *rule.StaticRule:
				return fmt.Sprintf(`%s != "%s"`, s, v.Value)
			case *rule.RangeRule:
				var fn string
				switch v.Format {
				case rule.Digit:
					fn = "ascii.IsDigit"
				case rule.AlphaNumeric:
					fn = "ascii.IsAlphaNumeric"
				case rule.UpperCaseLetters:
					fn = "ascii.IsUpperCaseLetter"
				}
				return fmt.Sprintf(`!ascii.Every(%s, %s)`, s, fn)
			}
			return `invalid code fix me`
		},
		"generator": func(r rule.Rule, builderName string) string {
			switch v := r.(type) {
			case *rule.StaticRule:
				return fmt.Sprintf(`%s.WriteString("%s")`, builderName, v.Value)
			case *rule.RangeRule:
				var fn string
				switch v.Format {
				case rule.Digit:
					fn = "Digits"
				case rule.AlphaNumeric:
					fn = "AlphaNumeric"
				case rule.UpperCaseLetters:
					fn = "UpperCaseLetters"
				}
				return fmt.Sprintf(`generator.%s(%s, %d)`, fn, builderName, v.Length)
			}
			return `invalid code fix me`
		},
	}
}
