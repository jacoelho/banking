package validation

import (
	"fmt"
	"strings"
	"sync"
	"text/template"

	"github.com/jacoelho/iban/registry/parser"

	"github.com/jacoelho/iban/registry/rule"
)

var once sync.Once
var validationTemplate *template.Template

func Generate() (string, error) {
	sb := new(strings.Builder)

	rules, err := parser.New("AL2!n8!n16!c").Parse()
	if err != nil {
		return "", nil
	}

	var bla = struct {
		CountryName string
		PackageName string
		Rules       []rule.Rule
	}{
		CountryName: "Albania",
		PackageName: "test",
		Rules:       rules,
	}

	funcs := template.FuncMap{
		"rule": func(r rule.Rule, s string) string {
			switch v := r.(type) {
			case *rule.StaticRule:
				return fmt.Sprintf(`%s != "%s"`, s, v.Value)
			case *rule.RangeRule:
				var r string
				switch v.Format {
				case rule.Digit:
					r = "ascii.IsDigit"
				case rule.AlphaNumeric:
					r = "ascii.IsAlphaNumeric"
				case rule.UpperCaseLetters:
					r = "ascii.IsUpperCaseLetter"
				}
				return fmt.Sprintf(`!ascii.Every(%s, %s)`, s, r)
			}
			return `panic("generator error")`
		},
	}

	once.Do(func() {
		validationTemplate = template.Must(template.New("").Funcs(funcs).ParseFiles("validate.go.tmpl", "static.go.tmpl"))
	})

	if err := validationTemplate.ExecuteTemplate(sb, "validate.go.tmpl", bla); err != nil {
		return "", err
	}

	return sb.String(), nil
}
