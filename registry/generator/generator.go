package generator

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"text/template"

	"github.com/jacoelho/banking/registry"

	"github.com/jacoelho/banking/registry/parser"
	"github.com/jacoelho/banking/registry/rule"
)

const generatedPackage = "iban"

var once sync.Once
var validationTemplate *template.Template

func GenerateValidationForCountry(w io.Writer, country registry.Country) error {
	once.Do(func() {
		validationTemplate = template.Must(template.New("").Funcs(templateFunctions()).Parse(validateTmpl))
	})

	rules, parseErr := parser.New(country.IBAN).Parse()
	if parseErr != nil {
		return fmt.Errorf("%v", parseErr)
	}

	var validationData = struct {
		CountryName string
		PackageName string
		Length      int
		Rules       []rule.Rule
	}{
		CountryName: strings.ReplaceAll(country.Name, " ", ""),
		PackageName: generatedPackage,
		Length:      rules[len(rules)-1].EndPos(),
		Rules:       rules,
	}

	return validationTemplate.ExecuteTemplate(w, "", validationData)
}
