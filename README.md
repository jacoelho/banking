# Banking

[![Go Reference](https://pkg.go.dev/badge/github.com/jacoelho/banking.svg)](https://pkg.go.dev/github.com/jacoelho/banking)

Go library for IBAN validation, IBAN generation, BBAN extraction, SEPA membership checks, and BIC parsing.

## Install

```bash
go get github.com/jacoelho/banking
```

## IBAN

IBAN validation uses country rules generated from the [Swift ISO 13616 IBAN Registry](https://www.swift.com/standards/data-standards/iban-international-bank-account-number) and ISO/IEC 7064 MOD 97-10 check digits.

Validation is structural. It does not verify that an account exists or can receive transactions.

[Test IBANs online](https://www.jacoelho.com/banking/)

### Usage

```go
package main

import (
	"fmt"

	"github.com/jacoelho/banking/iban"
)

func main() {
	if err := iban.Validate("GB29NWBK60161331926819"); err != nil {
		panic(err)
	}

	corrected, err := iban.ReplaceChecksum("GB99NWBK60161331926819")
	if err != nil {
		panic(err)
	}
	fmt.Println(corrected)

	generated, err := iban.Generate("GB")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(generated), generated[:2])

	generatedWithBBAN, err := iban.GenerateWithBBAN("GB", iban.BBANParts{
		BankCode:      "NWBK",
		BranchCode:    "601613",
		AccountNumber: "31926819",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(generatedWithBBAN)

	fmt.Println(iban.PaperFormat("GB29NWBK60161331926819"))

	bban, err := iban.GetBBAN("GB29NWBK60161331926819")
	if err != nil {
		panic(err)
	}
	fmt.Println(bban.BBAN, bban.BankCode, bban.BranchCode, bban.AccountNumber)

	isSEPA, err := iban.IsSEPA("GB29NWBK60161331926819")
	if err != nil {
		panic(err)
	}
	fmt.Println(isSEPA)

	isSEPACountry, err := iban.IsSEPACountryCode("GB")
	if err != nil {
		panic(err)
	}
	fmt.Println(isSEPACountry)
}
```

Output:

```text
GB29NWBK60161331926819
22 GB
GB29NWBK60161331926819
GB29 NWBK 6016 1331 9268 19
NWBK60161331926819 NWBK 601613 31926819
true
true
```

### Validation Errors

Validation returns structured errors that work with `errors.Is` and `errors.As`.

```go
package main

import (
	"errors"
	"fmt"

	"github.com/jacoelho/banking/iban"
)

func main() {
	err := iban.Validate("GB99INVALID")

	if errors.Is(err, iban.ErrInvalidIBAN) {
		fmt.Println("invalid IBAN")
	}

	var validationErr *iban.ValidationError
	if errors.As(err, &validationErr) && validationErr.Reason == iban.ReasonInvalidLength {
		fmt.Printf("invalid length: expected %d, got %d\n",
			validationErr.ExpectedLength,
			validationErr.ActualLength)
	}
}
```

Output:

```text
invalid IBAN
invalid length: expected 22, got 11
```

Validation reasons:

- `iban.ReasonInvalidLength`
- `iban.ReasonInvalidChecksum`
- `iban.ReasonInvalidCharacters`
- `iban.ReasonUnsupportedCountry`

Country-code APIs such as `Generate` and `IsSEPACountryCode` return `*iban.CountryCodeError`.

## BIC

```go
package main

import (
	"fmt"

	"github.com/jacoelho/banking/bic"
)

func main() {
	code, err := bic.Parse("ABCDBEB1XXX")
	if err != nil {
		panic(err)
	}

	fmt.Println(code.IsValid())
	fmt.Println(code.BIC8())
}
```

Output:

```text
true
ABCDBEB1
```

## Generate Registry Data

To regenerate IBAN validation code from an external Swift TXT registry file:

```bash
make update-registry SOURCE_REGISTRY=/path/to/iban-registry.txt
```

The full registry source is not committed to this repository. Generated Go files under `iban/` are the committed registry artifact.

## License

MIT License

See [LICENSE](LICENSE) to see the full text.
