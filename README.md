# Banking

Banking related library.

## Install

```bash
go get -u github.com/jacoelho/banking
```

## IBAN

Supports IBAN validation based on [swift rules](https://www.swift.com/sites/default/files/resources/iban_registry.txt).

### Usage

```go
err := iban.Validate("SOME IBAN")

// Or a specific iban directly
err := iban.ValidateUnitedKingdomIBAN("SOME GB IBAN")
```

## ISO-7064

Mod-97-10 implemented.

## Performance

One library goal is to be performant, any reflection, allocations must be careful measured.

## Roadmap

* generators
  * generator with specific values
* country bban verifier

## License

GNU General Public License v3.0 or later

See [LICENSE](LICENSE) to see the full text.
