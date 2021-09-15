# Banking

[![GoDoc](https://godoc.org/github.com/jacoelho/banking?status.svg)](https://pkg.go.dev/github.com/jacoelho/banking?tab=overview)

Banking related library.

## Install

```bash
go get -u github.com/jacoelho/banking
```

## IBAN

Supports IBAN validation based on [swift rules](https://www.swift.com/sites/default/files/resources/iban_registry.txt).

### Usage

#### Validation
```go
err := iban.Validate("SOME IBAN")
```

#### Replace check digits
```go
result, err := iban.ReplaceChecksum("GB99NWBK60161331926819")
// Output: GB29NWBK60161331926819
```

#### Generation
```go
account := iban.Generate("GB")
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

## ISO-7064

Mod-97-10 implemented.

## Generate

```
cd registry
go install -v ./...
cd ..
./bin/generator  -registry-file ./docs/registry.yml
```

## Roadmap

* generator with specific values
* country bban verifier

## License

MIT License

See [LICENSE](LICENSE) to see the full text.
