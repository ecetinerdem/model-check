package main

import "fmt"

func VerifyModel(model *LoanApprovalAI, property Property, applicants []Applicant) {
	propertyName := property.Name()
	fmt.Printf("Verifying %s...\n", propertyName)

	// Run the property check on the model and applicants
	satisfied, counterExamples := property.Check(model, applicants)

	// Print out vericication stats
	if satisfied {
		fmt.Printf("\u2713 %s is satisfied\n", propertyName)
	} else {
		fmt.Printf("\u2717 %s is violated. Found %d problematic cases. \n",
			propertyName, len(counterExamples))

		// PRint up to three examples for clarity
		for i := range min(3, len(counterExamples)) {
			a := counterExamples[i]
			fmt.Printf("   Example: %d: Income: $%.1fk, Credit Score: %.2f, Debt Ratio: %.2f, "+"Protected: %v, Decision: %v\n", i+1, a.income, a.creditScore, a.debtToIncome, a.protectedClass, model.ApproveLoan(a))
		}

		if len(counterExamples) > 3 {
			fmt.Printf("   ... and %d more\n", len(counterExamples)-3)
		}

	}

}
