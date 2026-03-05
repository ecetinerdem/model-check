package main

import (
	"fmt"
)

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

	riskProperty := &RiskProperty{
		maxHighRiskApprovalRate: 0.1, // At most 10 percent high risk applicants to be approved
	}
	// Create some test models
	models := []*LoanApprovalAI{
		{
			incomeWeight:      0.3,
			creditScoreWeight: 0.4,
			loanAmountWeight:  1.0,
			debtRatioWeight:   2.0,
			approvalTreshold:  5.0,
			employmentWeight:  0.1,
		},

		{
			incomeWeight:      0.25,
			creditScoreWeight: 0.45,
			loanAmountWeight:  1.2,
			debtRatioWeight:   2.5,
			approvalTreshold:  4.5,
			employmentWeight:  0.15,
		},

		{
			incomeWeight:      0.2,
			creditScoreWeight: 0.5,
			loanAmountWeight:  1.5,
			debtRatioWeight:   3.0,
			approvalTreshold:  4.0,
			employmentWeight:  0.2,
		},
	}

	// Test each model configuration against both properties
	descriptions := []string{
		"Loan Approval AI model with Initial Parameters",
		"Loan Approval AI model with Adjusted Parameters",
		"Loan Approval AI model with Final Parameters",
	}

	for i, model := range models {
		// PRint the current model's Parameters
		PrintModelPArams(model, descriptions[i])

		// Verify the model against the both properties
		VerifyModel(model, fairnessPropery, applicants)
		VerifyModel(model, riskProperty, applicants)
	}

}
