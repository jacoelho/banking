package generator

const validateTmpl = `// Code generated DO NOT EDIT.

package {{ .PackageName }}

import (
    "fmt"
    "github.com/jacoelho/banking/ascii"
)

func Validate{{ .CountryName }}IBAN(iban string) error {
    if len(iban) != {{ .Length }} {
        return fmt.Errorf("unexpected length, want: {{ .Length }}: %w", ErrValidation)
    }
    {{ range .Rules }}
    if subject := iban[{{ .StartPos }}:{{ .EndPos }}]; {{ rule . "subject" }} {
        return fmt.Errorf("{{ .String }}, found %s: %w", subject, ErrValidation)
    }
    {{ end }}
    return nil
}
`
