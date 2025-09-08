// Package registry provides IBAN (International Bank Account Number) validation
// and generation capabilities based on country-specific rules.
//
// The package parses IBAN structure definitions according to ISO 13616 standard
// and generates validation code for different countries.
//
// Main components:
//   - parser: Parses ISO 13616 IBAN structure definitions
//   - lexer: Tokenizes IBAN structure strings
//   - generator: Generates Go code for IBAN validation and generation
//   - rule: Defines validation rules for IBAN components
//
// Example IBAN structure: "GB2!n4!a6!n8!n" means:
//   - GB: Static country code
//   - 2!n: 2 digits (check digits)
//   - 4!a: 4 uppercase letters (bank code)
//   - 6!n: 6 digits (sort code)
//   - 8!n: 8 digits (account number)
package registry
