package iban

import "github.com/jacoelho/banking/internal/ascii"

type ibanRuleKind uint8

const (
	ibanRuleStatic ibanRuleKind = iota
	ibanRuleDigit
	ibanRuleUpperCase
	ibanRuleAlphaNumeric
)

const (
	countryIndexBits       = 7
	countryIndexMask uint8 = (1 << countryIndexBits) - 1
	countrySEPAFlag  uint8 = 1 << countryIndexBits
	ibanBBANOffset         = 4
)

type ibanRule struct {
	start  uint8
	length uint8
	kind   ibanRuleKind
	value  string
}

type ibanRuleViolation struct {
	offset        int
	length        int
	expected      CharClass
	expectedValue string
	actual        string
}

type bbanComponent struct {
	start uint8
	end   uint8
}

type countrySpec struct {
	code          string
	length        int
	rules         []ibanRule
	bankCode      bbanComponent
	branchCode    bbanComponent
	accountNumber bbanComponent
}

// countryCodeIndex stores packed entries: the low countryIndexBits bits are the
// 1-based countrySpecs index, and the high bit reports SEPA membership. If the
// registry grows past countryIndexMask countries, widen the entry type.
func lookupCountry(code string) (*countrySpec, bool) {
	slot, ok := countryIndexSlot(code)
	if !ok {
		return nil, false
	}
	index := countryCodeIndex[slot] & countryIndexMask
	if index == 0 {
		return nil, false
	}
	return &countrySpecs[int(index)-1], true
}

func validCountryCode(code string) bool {
	if len(code) != 2 {
		return false
	}
	_, ok := countryIndexSlot(code)
	return ok
}

// countryIndexSlot requires a two-byte country code. Public callers validate
// length before slicing or calling lookupCountry.
func countryIndexSlot(code string) (int, bool) {
	if code[0] < 'A' || code[0] > 'Z' || code[1] < 'A' || code[1] > 'Z' {
		return 0, false
	}
	return int(code[0]-'A')*26 + int(code[1]-'A'), true
}

func (c *countrySpec) validate(iban string) error {
	if len(iban) != c.length {
		return invalidIBANLength(c.length, len(iban))
	}
	for _, rule := range c.rules {
		start := int(rule.start)
		end := start + int(rule.length)
		subject := iban[start:end]
		violation, ok := rule.validate(subject, 0)
		if ok {
			continue
		}

		if violation.expectedValue != "" {
			return invalidIBANValue(start, violation.expectedValue, subject)
		}
		return validationCharacterError(rule, subject, violation.expected)
	}
	return nil
}

func validationCharacterError(rule ibanRule, actual string, expected CharClass) error {
	return invalidIBANCharacters(int(rule.start), int(rule.length), expected, actual)
}

func (r ibanRule) validate(subject string, ruleOffset int) (ibanRuleViolation, bool) {
	violation := ibanRuleViolation{
		length: len(subject),
		actual: subject,
	}

	switch r.kind {
	case ibanRuleStatic:
		expected := r.value[ruleOffset : ruleOffset+len(subject)]
		if subject == expected {
			return ibanRuleViolation{}, true
		}
		violation.expectedValue = expected
	case ibanRuleDigit:
		if ascii.IsDigit(subject) {
			return ibanRuleViolation{}, true
		}
		violation.expected = CharClassDigit
	case ibanRuleUpperCase:
		if ascii.IsUpperCase(subject) {
			return ibanRuleViolation{}, true
		}
		violation.expected = CharClassUpperAlpha
	case ibanRuleAlphaNumeric:
		if ascii.IsUpperAlphaNumeric(subject) {
			return ibanRuleViolation{}, true
		}
		violation.expected = CharClassUpperAlphaNumeric
	}

	return violation, false
}

func (c *countrySpec) generate() string {
	buf := make([]byte, 0, c.length)
	buf = append(buf, c.code...)
	buf = append(buf, '0', '0')
	for _, rule := range c.rules[2:] {
		switch rule.kind {
		case ibanRuleStatic:
			buf = append(buf, rule.value...)
		case ibanRuleDigit:
			buf = ascii.AppendRandomDigits(buf, rule.length)
		case ibanRuleUpperCase:
			buf = ascii.AppendRandomUpperCaseLetters(buf, rule.length)
		case ibanRuleAlphaNumeric:
			buf = ascii.AppendRandomAlphaNumeric(buf, rule.length)
		}
	}
	return replaceChecksumBytes(buf)
}

func (c *countrySpec) bban(iban string) BBAN {
	return BBAN{
		BBAN:          iban[4:c.length],
		BankCode:      c.bankCode.slice(iban),
		BranchCode:    c.branchCode.slice(iban),
		AccountNumber: c.accountNumber.slice(iban),
	}
}

func (c bbanComponent) slice(iban string) string {
	if !c.present() {
		return ""
	}
	return iban[int(c.start)+ibanBBANOffset : int(c.end)+ibanBBANOffset]
}

func (c bbanComponent) present() bool {
	return c.end > c.start
}
