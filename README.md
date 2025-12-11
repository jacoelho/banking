# Banking

[![GoDoc](https://godoc.org/github.com/jacoelho/banking?status.svg)](https://pkg.go.dev/github.com/jacoelho/banking?tab=overview)

Banking related library.

## Install

```bash
go get -u github.com/jacoelho/banking
```

## IBAN

Supports IBAN validation based on [swift rules](https://www.swift.com/node/11971).

[Test IBANs online](https://www.jacoelho.com/banking/)

### Usage

#### Validation
```go
err := iban.Validate("SOME IBAN")
```

For simple validation error detection:
```go
if iban.IsValidationError(err) {
    // Handle validation errors
}
```

#### Error Handling

For detailed validation feedback, the library provides structured error types:

```go
import "errors"

err := iban.Validate("GB99INVALID")
if err != nil {
    var lengthErr *iban.ErrValidationLength
    if errors.As(err, &lengthErr) {
        fmt.Printf("Invalid length: expected %d, got %d\n", lengthErr.Expected, lengthErr.Actual)
    }

    var checksumErr *iban.ErrValidationChecksum
    if errors.As(err, &checksumErr) {
        fmt.Printf("Invalid checksum: expected %s, got %s\n", checksumErr.Expected, checksumErr.Actual)
    }

    var rangeErr *iban.ErrValidationRange
    if errors.As(err, &rangeErr) {
        fmt.Printf("Invalid characters at position %d (length %d): expected %s\n",
            rangeErr.Position, rangeErr.Length, rangeErr.Expected)
    }

    var staticErr *iban.ErrValidationStaticValue
    if errors.As(err, &staticErr) {
        fmt.Printf("Invalid value at position %d: expected %s, got %s\n",
            staticErr.Position, staticErr.Expected, staticErr.Actual)
    }

    var unsupportedErr *iban.ErrUnsupportedCountry
    if errors.As(err, &unsupportedErr) {
        fmt.Printf("Country code '%s' is not supported\n", unsupportedErr.CountryCode)
    }
}
```

#### Replace check digits
```go
result, err := iban.ReplaceChecksum("GB99NWBK60161331926819")
// Output: GB29NWBK60161331926819
```

#### Generation

```go
iban, err := iban.Generate("GB")
// Output: GB29NWBK60161331926819
```

#### Printing

```go
iban.PaperFormat("GB29NWBK60161331926819"))
// Output: GB29 NWBK 6016 1331 9268 19
```

#### BBAN

```go
// Get BBAN from IBAN
result, _ := iban.GetBBAN("GB29NWBK60161331926819")

fmt.Println(result.BBAN, result.BankCode, result.BranchCode, result.AccountNumber)
// Output: NWBK60161331926819 NWBK 601613 31926819
```

#### IsSEPA

```go
// IsSEPA returns if an iban country is a SEPA member
result, _ := iban.IsSEPA("GB29NWBK60161331926819")
fmt.Println(result)
// Output: true
```

## ISO-7064

Mod-97-10 implemented.

## Generate

To update the registry from a TSV file and regenerate IBAN validation code:

```bash
make update-registry REGISTRY=iban-registry-v101.txt
```

## License

MIT License

See [LICENSE](LICENSE) to see the full text.
