package generator

import (
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/jacoelho/banking/registry/rule"
)

func templateFunctions() template.FuncMap {
	return template.FuncMap{
		"ToLower": func(s string) string {
			return strings.ToLower(s)
		},
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
					fn = "ascii.IsUpperCase"
				}
				return fmt.Sprintf(`!%s(%s)`, fn, s)
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
				return fmt.Sprintf(`ascii.%s(%s, %d)`, fn, builderName, v.Length)
			}
			return `invalid code fix me`
		},
		"bban": func(s string, variableName string) string {
			if s == "" {
				return `""`
			}

			fields := strings.Split(s, ":")
			if len(fields) != 2 {
				return `invalid code fix me`
			}

			start, err := strconv.Atoi(fields[0])
			if err != nil {
				return `invalid code fix me`
			}

			end, err := strconv.Atoi(fields[1])
			if err != nil {
				return `invalid code fix me`
			}

			return fmt.Sprintf("%s[%d:%d]", variableName, start+4, end+4)
		},
	}
}
