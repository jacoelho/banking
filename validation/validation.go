package validation

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"text/template"

	"gopkg.in/yaml.v2"

	"github.com/jacoelho/iban/registry/parser"

	"github.com/jacoelho/iban/registry/rule"
)

var once sync.Once
var validationTemplate *template.Template

type Country struct {
	Code string `yaml:"code"`
	Name string `yaml:"name"`
	IBAN string `yaml:"IBAN"`
	BBAN string ` yaml:"BBAN"`
}

func Generate() (string, error) {
	once.Do(func() {
		validationTemplate = template.Must(template.New("").Funcs(templateFunctions()).ParseFiles("validate.go.tmpl", "static.go.tmpl"))
	})

	f, err := os.Open("../docs/registry.yml")
	if err != nil {
		return "", err
	}
	defer f.Close()

	var countries struct {
		Countries []Country
	}

	if err := yaml.NewDecoder(f).Decode(&countries); err != nil {
		return "", nil
	}

	for _, v := range countries.Countries {
		fmt.Println(v.Name)

		sb, err := os.OpenFile("../samples/validator_"+strings.ReplaceAll(strings.ToLower(v.Name), " ", "_")+".go", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
		if err != nil {
			return "", err
		}

		rules, parseErr := parser.New(v.IBAN).Parse()
		if parseErr != nil {
			return "", errors.New("parse error")
		}

		var bla = struct {
			CountryName string
			PackageName string
			Length      int
			Rules       []rule.Rule
		}{
			CountryName: strings.ReplaceAll(v.Name, " ", ""),
			PackageName: "sample",
			Length:      rules[len(rules)-1].EndPos(),
			Rules:       rules,
		}

		if err := validationTemplate.ExecuteTemplate(sb, "validate.go.tmpl", bla); err != nil {
			return "", err
		}
		sb.Close()

	}
	return "", nil
}

func templateFunctions() template.FuncMap {
	return template.FuncMap{
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
}
