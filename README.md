Loan Approval AI Verifier
A Go tool for testing and verifying loan approval AI models against fairness and risk properties using real applicant data.

Overview
This project loads applicant data from a CSV file, runs it through configurable loan approval models, and checks whether each model satisfies defined behavioral properties — specifically around fairness across protected classes and risk management.

Project Structure
FileDescriptionmain.goEntry point — sets up models, properties, and runs verificationmodel.goLoanApprovalAI struct and ApproveLoan scoring logicloader.goParses applicant data from CSV with flexible column detectionverify.goVerifyModel runner — prints results and counter-examplesfairness.goFairnessProperty — checks approval rate disparity across groupsrisk.goRiskProperty — checks approval rate for high-risk applicantsloan_applicants.csvSample dataset of 20 applicants


How It Works
Scoring — Each applicant is scored using a weighted formula:
score = income×incomeWeight + creditScore×creditScoreWeight
      - loanToIncomeRatio×loanAmountWeight
      - debtToIncome×debtRatioWeight
      + yearsEmployed×employmentWeight
The applicant is approved if score > approvalThreshold.
Properties checked:

Fairness — approval rate disparity between protected and non-protected groups must not exceed maxDisparity (default 5%)
Risk — approval rate among high-risk applicants (low credit score + high debt ratio) must not exceed maxHighRiskApprovalRate (default 10%)

CSV Format
```
Income,Credit_Score,Loan_Amount,Debt_to_Income,Years_Employed,Protected_Class
75000,720,250000,0.35,5,true
```
Column names are matched flexibly (case-insensitive, partial match).

Running
```bash
go run .
```
