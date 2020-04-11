package registry

import (
	"encoding/csv"
	"fmt"
	"io"
)

type DecodeError string

func (p DecodeError) Error() string {
	return string(p)
}

const (
	ErrLineParse   = DecodeError("failed to parse line")
	ErrInvalidFile = DecodeError("invalid file")
)

func Decode(r io.Reader) ([]Entry, error) {
	if r == nil {
		return nil, ErrInvalidFile
	}

	reader := csv.NewReader(r)
	reader.Comma = '\t'

	// skip header
	_, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, ErrInvalidFile)
	}

	var entries []Entry
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		entry, err := parseLine(line)
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func parseLine(line []string) (Entry, error) {
	if len(line) != 18 {
		return Entry{}, fmt.Errorf("%s invalid length: %w", line, ErrLineParse)
	}

	return Entry{
		CountryName:                  line[0],
		CountryCode:                  line[1],
		DomesticAccountNumberExample: line[2],
		BBAN: BBAN{
			Structure:              line[4],
			Length:                 line[5],
			BankIdentifierPosition: line[6],
			BankIdentifierLength:   line[7],
			BankIdentifierExample:  line[8],
			Example:                line[9],
		},
		IBAN: IBAN{
			Structure:               line[11],
			Length:                  line[12],
			ElectronicFormatExample: line[13],
			PrintFormatExample:      line[14],
		},
	}, nil
}