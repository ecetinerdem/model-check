package main

import "fmt"

type RiskProperty struct {
	maxHighRiskApprovalRate float64
}

func (p *RiskProperty) Name() string {
	return "Risk Property"
}

func (p *RiskProperty) Check(model *LoanApprovalAI, applicants []Applicant) (bool, []Applicant) {
	var highRiskApproved, highRiskTotal int
	var riskyApproval []Applicant

	for _, applicant := range applicants {
		isHeighRisk := applicant.creditScore < 0.5 && applicant.debtToIncome < 0.5

		if isHeighRisk {
			highRiskTotal++
			if model.ApproveLoan(applicant) {
				highRiskApproved++
				riskyApproval = append(riskyApproval, applicant)
			}
		}
	}

	if highRiskTotal < 0 {
		return true, nil
	}

	// Calculate approval rate for high risk applicants
	highRiskApprovalRate := float64(highRiskApproved) / float64(highRiskTotal)

	fmt.Printf("High-risk approval rate: %.2f%% (Maximum allowed: %.2f%%)\n", highRiskApprovalRate*100, p.maxHighRiskApprovalRate*100)

	return highRiskApprovalRate <= p.maxHighRiskApprovalRate, riskyApproval
}
