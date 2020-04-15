package iban

type BBAN struct {
	BBAN             string
	BankCode         string
	BranchCode       string
	NationalChecksum string
	AccountNumber    string
}
