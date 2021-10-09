package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Birthday describes the structure of an expected record in the CSV
type Birthday struct {
	Name        string
	Day         uint8
	Month       uint8
	Description string
}

func fromCSV(raw io.Reader) ([]Birthday, error) {
	reader := csv.NewReader(raw)
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true
	lines, err := reader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("an error occurred while reading the csv: %w", err)
	}

	birthdays := make([]Birthday, len(lines))
	for i, line := range lines {
		birthdays[i], err = fromCSVLine(line)
		if err != nil {
			return nil, fmt.Errorf("Invalid birthday record at line %d: %w, got: [%s]", i+1, err, strings.Join(line, ","))
		}
	}

	return birthdays, nil
}

const formatError = "expected format for CSV is: 'name','day','month','description' or 'name','day','month'"

func fromCSVLine(line []string) (Birthday, error) {
	if len(line) != 4 && len(line) != 3 {
		return Birthday{}, fmt.Errorf(formatError)
	}

	// TODO: more strict validation of dates (31 of february should not work)
	day, err := strconv.ParseUint(line[1], 10, 8)
	if err != nil || day > 31 || day == 0 {
		return Birthday{}, fmt.Errorf("invalid value for day: '%s', %s", line[1], formatError)
	}

	month, err := strconv.ParseUint(line[2], 10, 8)
	if err != nil || month > 12 || month == 0 {
		return Birthday{}, fmt.Errorf("invalid value for month: '%s', %s", line[2], formatError)
	}

	description := ""
	if len(line) == 4 {
		description = line[3]
	}

	return Birthday{
		Name: line[0], Day: uint8(day), Month: uint8(month), Description: description}, nil
}
