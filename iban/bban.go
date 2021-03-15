package iban

// BBAN is short for Basic Bank Account Number
// It represents a country-specific bank account number.
type BBAN struct {
	BBAN             string
	BankCode         string
	BranchCode       string
	NationalChecksum string
	AccountNumber    string
}
