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

func GenerateCodeForCountry(w io.Writer, country registry.Country) error {
	once.Do(func() {
		validateCountryTemplate = template.Must(template.New("").Funcs(templateFunctions()).Parse(validateCountryTmpl))
	})

	rules, parseErr := parser.New(country.IBAN).ReducedParse()
	if parseErr != nil {
		return fmt.Errorf("%v", parseErr)
	}

	var validationData = struct {
		PackageName      string
		FunctionValidate string
		FunctionGenerate string
		FunctionBBAN     string
		Country          registry.Country
		Length           int
		Rules            []rule.Rule
	}{
		FunctionValidate: validateFunctionName(country.Name),
		FunctionGenerate: generateFunctionName(country.Name),
		FunctionBBAN:     getBBANFunctionName(country.Name),
		PackageName:      generatedPackage,
		Country:          country,
		Length:           rules[len(rules)-1].EndPos(),
		Rules:            rules,
	}

	return validateCountryTemplate.ExecuteTemplate(w, "", validationData)
}

func validateFunctionName(s string) string {
	return fmt.Sprintf("validate%sIBAN", strings.ReplaceAll(s, " ", ""))
}

func generateFunctionName(s string) string {
	return fmt.Sprintf("generate%sIBAN", strings.ReplaceAll(s, " ", ""))
}

func getBBANFunctionName(s string) string {
	return fmt.Sprintf("get%sBBAN", strings.ReplaceAll(s, " ", ""))
}

type validateCountry struct {
	Code string
	Fn   string
}

func GenerateValidate(w io.Writer, countries []registry.Country) error {
	tmpl, err := template.New("").Funcs(templateFunctions()).Parse(validateTmpl)
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

func GenerateIsSEPA(w io.Writer, countries []registry.Country) error {
	tmpl, err := template.New("").Funcs(templateFunctions()).Parse(isSepaTmpl)
	if err != nil {
		return err
	}

	var data = struct {
		PackageName string
		Countries   []registry.Country
	}{
		PackageName: generatedPackage,
		Countries:   countries,
	}

	return tmpl.ExecuteTemplate(w, "", data)
}

func GenerateGenerate(w io.Writer, countries []registry.Country) error {
	tmpl, err := template.New("").Funcs(templateFunctions()).Parse(generateTmpl)
	if err != nil {
		return err
	}

	var functions []validateCountry
	for _, country := range countries {
		functions = append(functions, validateCountry{
			Code: country.Code,
			Fn:   generateFunctionName(country.Name),
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

func GenerateGetBBAN(w io.Writer, countries []registry.Country) error {
	tmpl, err := template.New("").Funcs(templateFunctions()).Parse(getBbanTmpl)
	if err != nil {
		return err
	}

	var functions []validateCountry
	for _, country := range countries {
		functions = append(functions, validateCountry{
			Code: country.Code,
			Fn:   getBBANFunctionName(country.Name),
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

func GenerateConstants(w io.Writer, countries []registry.Country) error {
	tmpl, err := template.New("").Funcs(templateFunctions()).Parse(countryCodeConstantsTmpl)
	if err != nil {
		return err
	}

	var data = struct {
		PackageName string
		Countries   []registry.Country
	}{
		PackageName: generatedPackage,
		Countries:   countries,
	}

	return tmpl.ExecuteTemplate(w, "", data)
}
