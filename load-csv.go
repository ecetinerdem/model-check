package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func LoadApplicantsFromCSV(csvFilePath string) ([]Applicant, error) {
	file, err := os.Open(csvFilePath)

	if err != nil {
		return nil, fmt.Errorf("error opening file: %v\n", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the header row
	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading header row: %v\n", err)
	}

	columnIndecies := map[string]int{
		"income":         -1,
		"creditScore":    -1,
		"loanAmount":     -1,
		"debtToIncome":   -1,
		"yearsEmployed":  -1,
		"protectedClass": -1,
	}

	// Find the column indicies
	for i, column := range header {

		col := strings.ToLower(strings.TrimSpace(column))

		switch {
		case strings.Contains(col, "income") && !strings.Contains(col, "debt"):
			columnIndecies["income"] = i
		case strings.Contains(col, "credit"):
			columnIndecies["creditScore"] = i
		case strings.Contains(col, "loan"):
			columnIndecies["loanAmount"] = i
		case strings.Contains(col, "debt"):
			columnIndecies["debtToIncome"] = i
		case strings.Contains(col, "employ"):
			columnIndecies["yearsEmployed"] = i
		case strings.Contains(col, "protect"):
			columnIndecies["protectedClass"] = i
		}
	}
	// Verify all columns are found

	for field, idx := range columnIndecies {
		if idx == -1 {
			return nil, fmt.Errorf("required column %s not found in CSV\n", field)
		}
	}

	// Read applicant records

	var applicants []Applicant
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %v\n", err)
	}

	for i, record := range records {
		// Parse values
		income, err := parseFloat(record[columnIndecies["income"]])
		if err != nil {
			return nil, fmt.Errorf("invalid income at line %d: %v", i+2, err)
		}

		// Convert income to thousands

		income = income / 1000

		creditScore, err := parseFloat(record[columnIndecies["creditScore"]])

		if err != nil {
			return nil, fmt.Errorf("invalid credit score at line %d: %v", i+2, err)
		}

		// Normalize credit score if 300 - 500 range
		if creditScore > 1 {
			creditScore = (creditScore - 300) / 500
		}

		loanAmount, err := parseFloat(record[columnIndecies["loanAmount"]])

		if err != nil {
			return nil, fmt.Errorf("invalid loan amount at line %d: %v", i+2, err)
		}

		// Convert loan amount to thousands
		loanAmount = loanAmount / 1000

		debtToIncome, err := parseFloat(record[columnIndecies["debtToIncome"]])

		if err != nil {
			return nil, fmt.Errorf("invalid debt ratioat line %d: %v", i+2, err)
		}

		// Normalize debt ratio if it is given as percentage
		if debtToIncome > 1 {
			debtToIncome = debtToIncome / 100
		}

		yearsEmployed, err := parseFloat(record[columnIndecies["yearsEmployed"]])

		if err != nil {
			return nil, fmt.Errorf("invalid years employed at line %d: %v", i+2, err)
		}

		protectedClass, err := parseBool(record[columnIndecies["protectedClass"]])
		if err != nil {
			return nil, fmt.Errorf("invalid protected class  at line %d: %v", i+2, err)
		}

		applicants = append(applicants, Applicant{
			income:         income,
			creditScore:    creditScore,
			loanAmount:     loanAmount,
			debtToIncome:   debtToIncome,
			yearsEmployed:  yearsEmployed,
			protectedClass: protectedClass,
		})

	}
	fmt.Printf("Successfully loaded %d applicants from CSV\n", len(applicants))

	return applicants, nil
}
