# Banking

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

// Or a specific iban directly
err := iban.ValidateUnitedKingdomIBAN("SOME GB IBAN")
```

#### Generation
```go
account := iban.Generate("GB")

// Or a specific iban directly
account := iban.GenerateUnitedKingdomIBAN()
```

#### BBAN

```go 
bban, _ := GetBBAN("GB29NWBK60161331926819")

// result
BBAN {
    BBAN:             "NWBK60161331926819",
    BankCode:         "NWBK",
    BranchCode:       "601613",
    NationalChecksum: "",
    AccountNumber:    "31926819",
}

bban.String()
"NWBK 6016 1331 9268 19"
```

## ISO-7064

Mod-97-10 implemented.

## Roadmap

* generator with specific values
* country bban verifier

## License

GNU General Public License v3.0 or later

See [LICENSE](LICENSE) to see the full text.
