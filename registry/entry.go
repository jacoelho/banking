package registry

type BBAN struct {
	Structure              string `json:"structure"`
	Length                 string `json:"length"`
	BankIdentifierPosition string `json:"bank_identifier_position"`
	BankIdentifierLength   string `json:"bank_identifier_length"`
	BankIdentifierExample  string `json:"bank_identifier_example"`
	Example                string `json:"example"`
}

type IBAN struct {
	Structure               string `json:"structure"`
	Length                  string `json:"length"`
	ElectronicFormatExample string `json:"electronic_format_example"`
	PrintFormatExample      string `json:"print_format_example"`
}

type Entry struct {
	CountryName                  string `json:"country_name"`
	CountryCode                  string `json:"country_code"`
	DomesticAccountNumberExample string `json:"domestic_account_number_example"`
	BBAN                         BBAN   `json:"bban"`
	IBAN                         IBAN   `json:"iban"`
}
