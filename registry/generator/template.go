package generator

const validateCountryTmpl = `// Code generated by banking/registry; DO NOT EDIT.

package {{ .PackageName }}

import (
    "fmt"
    "github.com/jacoelho/banking/pool"
    "github.com/jacoelho/banking/ascii"
)

// {{ .FunctionValidate }} validates {{ .Country.Name }} IBAN
func {{ .FunctionValidate }}(iban string) error {
    if len(iban) != {{ .Length }} {
        return fmt.Errorf("unexpected length, want: {{ .Length }}: %w", ErrValidation)
    }
    {{ range .Rules }}
    if subject := iban[{{ .StartPos }}:{{ .EndPos }}]; {{ rule . "subject" }} {
        return fmt.Errorf("{{ .String }}, found %s: %w", subject, ErrValidation)
    }
    {{ end }}
	if c := checksum(iban); c != iban[2:4] {
		return fmt.Errorf("incorrect checksum: %w", ErrValidation)
	}

    return nil
}

// {{ .FunctionGenerate }} generates {{ .Country.Name }} IBAN
func {{ .FunctionGenerate }}() (string, error) {
	sb := pool.BytesPool.Get()
	defer sb.Free()

	{{ range .Rules }}
    {{ generator . "sb" -}}
    {{ end }}

	return ReplaceChecksum(sb.String())
}

// {{ .FunctionBBAN }} retrieves BBAN structure from {{ .Country.Name }} IBAN
func {{ .FunctionBBAN }}(iban string) (BBAN, error) {
	if len(iban) != {{ .Length }} {
        return BBAN{}, fmt.Errorf("unexpected length, want: {{ .Length }}: %w", ErrValidation)
    }

	return BBAN {
		BBAN: iban[4:{{ .Length }}],
		BankCode: {{ bban .Country.BankCode "iban" }},
		BranchCode:  {{ bban .Country.BranchCode "iban" }},
		NationalChecksum:  {{ bban .Country.NationalChecksum "iban" }},
		AccountNumber: {{ bban .Country.AccountNumber "iban" }},
	}, nil
}
`

const validateTmpl = `// Code generated by banking/registry; DO NOT EDIT.

package {{ .PackageName }}

import (
    "fmt"
)

// Validate an IBAN
func Validate(iban string) error {
    if len(iban) < 2 {
        return fmt.Errorf("unexpected iban length: %w", ErrValidation)
    }

	code := iban[0:2]
	switch code {
    {{- range .Functions }}
    case {{ ToLower .Code }}:
		return {{ .Fn }}(iban)
    {{- end }}

	default:
		return fmt.Errorf("%s is not supported: %w", code, ErrValidation)
	}
}
`

const isSepaTmpl = `// Code generated by banking/registry; DO NOT EDIT.

package {{ .PackageName }}

import (
    "fmt"
)

// IsSEPACountryCode returns if a country code is a SEPA member
func IsSEPACountryCode(countryCode string) (bool, error) {
    if len(countryCode) != 2 {
        return false, fmt.Errorf("unexpected country code length: %w", ErrValidation)
    }

	switch countryCode {
	{{- range .Countries }}
	// Country Code {{ .Name }}
	case {{ ToLower .Code }}:
		return {{ .IsSEPA }}, nil
    {{- end }}
	default:
		return false, fmt.Errorf("%s is not supported: %w", countryCode, ErrValidation)
	}
}

// IsSEPA returns if an iban country is a SEPA member
func IsSEPA(iban string) (bool, error) {
    if len(iban) < 2 {
        return false, fmt.Errorf("unexpected iban length: %w", ErrValidation)
    }

	return IsSEPACountryCode(iban[0:2])
}
`

const getBbanTmpl = `// Code generated by banking/registry; DO NOT EDIT.

package {{ .PackageName }}

import (
    "fmt"
)

// GetBBAN retrieves BBAN from an iban
func GetBBAN(iban string) (BBAN, error) {
    if len(iban) < 2 {
        return BBAN{}, fmt.Errorf("unexpected iban length: %w", ErrValidation)
    }

	code := iban[0:2]
	switch code {
    {{- range .Functions }}
    case {{ ToLower .Code }}:
		return {{ .Fn }}(iban)
    {{- end }}

	default:
		return BBAN{}, fmt.Errorf("%s is not supported: %w", code, ErrValidation)
	}
}
`

const generateTmpl = `// Code generated by banking/registry; DO NOT EDIT.

package {{ .PackageName }}

import (
    "fmt"
)

// Generate IBAN based on ISO 3166-1 country code
func Generate(countryCode string) (string, error) {
	switch countryCode {
    {{- range .Functions }}
    case {{ ToLower .Code }}:
		return {{ .Fn }}()
    {{- end }}

	default:
		return "", fmt.Errorf("%s is not supported: %w", countryCode, ErrValidation)
	}
}
`

const countryCodeConstantsTmpl = `// Code generated by banking/registry; DO NOT EDIT.

package {{ .PackageName }}

const (
    {{- range .Countries }}
	// Country Code {{ .Name }}
    {{ ToLower .Code }} = "{{ .Code }}"
    {{- end }}
)
`
