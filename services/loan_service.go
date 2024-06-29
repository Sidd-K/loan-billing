package services

import (
	"time"

	"loan-building/models"
)

var loans = make(map[int]models.Loan)
var currentID = 1

func CreateLoan(req models.CreateLoanRequest) models.Loan {
	weeklyPayment := (req.Principal * (1 + req.InterestRate)) / float64(req.TermWeeks)
	outstanding := req.Principal * (1 + req.InterestRate)

	loan := models.Loan{
		ID:            currentID,
		Principal:     req.Principal,
		InterestRate:  req.InterestRate,
		TermWeeks:     req.TermWeeks,
		WeeklyPayment: weeklyPayment,
		Outstanding:   outstanding,
		CreatedAt:     time.Now(),
	}

	loans[currentID] = loan
	currentID++
	return loan
}

func GetOutstanding(id int) (float64, bool) {
	loan, exists := loans[id]
	if !exists {
		return 0, false
	}
	return loan.Outstanding, true
}

func IsDelinquent(id int) (bool, bool) {
	loan, exists := loans[id]
	if !exists {
		return false, false
	}

	delinquent := false
	missedPayments := 0
	for i := len(loan.Payments) - 1; i >= 0; i-- {
		if time.Since(loan.Payments[i].PaidAt) > 7*24*time.Hour {
			missedPayments++
			if missedPayments >= 2 {
				delinquent = true
				break
			}
		} else {
			break
		}
	}

	return delinquent, true
}

func MakePayment(id int, payment models.Payment) (models.Loan, bool) {
	loan, exists := loans[id]
	if !exists {
		return models.Loan{}, false
	}

	if payment.Amount != loan.WeeklyPayment {
		return models.Loan{}, false
	}

	loan.Outstanding -= payment.Amount
	loan.Payments = append(loan.Payments, payment)
	loans[id] = loan

	return loan, true
}
