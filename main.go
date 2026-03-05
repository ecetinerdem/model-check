package main

import "fmt"

func main() {
	// Set up a CSV file path
	csvFilePath := "loan_applicants.csv"

	// Load applicant data from CSV
	applicants, err := LoadApplicantsFromCSV(csvFilePath)

	if err != nil {
		fmt.Printf("Error loading applicants: %v\n", err)
	}

	// Define some properties we want to check (fairness and risk)
	fairnessPropery := &FairnessProperty{
		maxDisparity: 0.05, // At most 5 percent disparity in approval rates
	}
	// Create some test models

	// Test each model configuration against both properties

}
