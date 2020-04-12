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

var (
	once                    sync.Once
	validateCountryTemplate *template.Template
)

func GenerateValidationForCountry(w io.Writer, country registry.Country) error {
	once.Do(func() {
		validateCountryTemplate = template.Must(template.New("").Funcs(templateFunctions()).Parse(validateCountryTmpl))
	})

	rules, parseErr := parser.New(country.IBAN).Parse()
	if parseErr != nil {
		return fmt.Errorf("%v", parseErr)
	}

	var validationData = struct {
		PackageName  string
		FunctionName string
		Length       int
		Rules        []rule.Rule
	}{
		FunctionName: validateFunctionName(country.Name),
		PackageName:  generatedPackage,
		Length:       rules[len(rules)-1].EndPos(),
		Rules:        rules,
	}

	return validateCountryTemplate.ExecuteTemplate(w, "", validationData)
}

func validateFunctionName(s string) string {
	return fmt.Sprintf("Validate%sIBAN", strings.ReplaceAll(s, " ", ""))
}

type validateCountry struct {
	Code string
	Fn   string
}

func GenerateValidate(w io.Writer, countries []registry.Country) error {
	tmpl, err := template.New("").Parse(validateTmpl)
	if err != nil {
		return err
	}

	var functions []validateCountry
	for _, country := range countries {
		functions = append(functions, validateCountry{
			Code: country.Code,
			Fn:   validateFunctionName(country.Name),
		})
	}

	var data = struct {
		PackageName string
		Functions   []validateCountry
	}{
		PackageName: generatedPackage,
		Functions:   functions,
	}

	return tmpl.ExecuteTemplate(w, "", data)
}
