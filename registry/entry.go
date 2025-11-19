package registry

// BBAN represents the Basic Bank Account Number structure
type BBAN struct {
	Structure              string `json:"structure"`
	Length                 int    `json:"length"`
	BankIdentifierPosition string `json:"bank_identifier_position"`
	BankIdentifierLength   string `json:"bank_identifier_length"`
	BankIdentifierExample  string `json:"bank_identifier_example"`
	Example                string `json:"example"`
}

// IBAN represents the International Bank Account Number structure
type IBAN struct {
	Structure               string `json:"structure"`
	Length                  int    `json:"length"`
	ElectronicFormatExample string `json:"electronic_format_example"`
	PrintFormatExample      string `json:"print_format_example"`
}

// Entry represents a complete IBAN registry entry for a country
type Entry struct {
	CountryName                  string `json:"country_name"`
	CountryCode                  string `json:"country_code"`
	DomesticAccountNumberExample string `json:"domestic_account_number_example"`
	BBAN                         BBAN   `json:"bban"`
	IBAN                         IBAN   `json:"iban"`
}

// Country represents a country's IBAN configuration
type Country struct {
	Code          string `yaml:"code"`
	Name          string `yaml:"name"`
	IBAN          string `yaml:"IBAN"`
	BBAN          string `yaml:"BBAN"`
	BankCode      string `yaml:"bank_code"`
	BranchCode    string `yaml:"branch_code"`
	AccountNumber string `yaml:"account_number"`
	IsSEPA        bool   `yaml:"sepa"`
}

// Countries represents a collection of country IBAN configurations
type Countries struct {
	Countries []Country `yaml:"countries"`
}
