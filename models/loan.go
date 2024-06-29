package models

import (
	"time"
)

type Loan struct {
	ID            int
	Principal     float64
	InterestRate  float64
	TermWeeks     int
	WeeklyPayment float64
	Outstanding   float64
	Payments      []Payment
	CreatedAt     time.Time
}

type Payment struct {
	Amount float64
	PaidAt time.Time
}

type CreateLoanRequest struct {
	Principal    float64 `json:"principal"`
	InterestRate float64 `json:"interest_rate"`
	TermWeeks    int     `json:"term_weeks"`
}
