package iban

// BBAN is short for Basic Bank Account Number
// It represents a country-specific bank account number.
type BBAN struct {
	BBAN          string
	BankCode      string
	BranchCode    string
	AccountNumber string
}

// BBANParts contains BBAN fields that constrain IBAN generation.
// Empty fields are generated from country rules.
type BBANParts struct {
	BankCode      string
	BranchCode    string
	AccountNumber string
}
