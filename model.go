package main

type LoanApprovalAI struct {
	// Weights for different factors used in the decision making process
	incomeWeight      float64
	creditScoreWeight float64
	loanAmountWeight  float64
	debtRatioWeight   float64
	employmentWeight  float64
	approvalTreshold  float64
}

type Applicant struct {
	income         float64 // Annual income in thousands, so 50 == $50,000.00
	creditScore    float64 // Credit score  normalized to 0-1, from a typical 300-500
	loanAmount     float64 // Loan amount in thousands
	debtToIncome   float64 // Debt to income ratio 0-1, already normalized
	yearsEmployed  float64
	protectedClass bool // Whether or not applicant belongs to soe protected class
}

// ApproveLoan determines if the applicants should be approved for a loan

func (ai *LoanApprovalAI) ApproveLoan(applicant Applicant) bool {
	loanToIncomeRatio := 0.0

	if applicant.income > 0 {
		loanToIncomeRatio := applicant.loanAmount / applicant.income

	}
	score := applicant.income*ai.incomeWeight +
		applicant.creditScore*ai.incomeWeight -
		loanToIncomeRatio*ai.loanAmountWeight -
		applicant.debtToIncome*ai.debtRatioWeight +
		applicant.yearsEmployed*ai.employmentWeight

	return score > ai.approvalTreshold

}

// Some means of determining fairness

// Check verifies if the AI Model satisfies the fairness property

// Evaluate risk

// Loading CSV file

// Verify model in general. Checks to see if the model satisfies some property
