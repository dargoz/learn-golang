package entity

import "errors"

type Account struct {
	ID            int
	Name          string
	Balance       float64
	Status        string
	AccountNumber string
}

func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.Balance += amount
		a.Status = "Active"
	}
}

func (a *Account) Withdraw(amount float64) {
	if amount > 0 && amount <= a.Balance {
		a.Balance -= amount
		if a.Balance == 0 {
			a.Status = "Inactive"
		}
	}
}

func (a *Account) Transfer(amount float64, to *Account) error {
	if amount > 0 && amount <= a.Balance {
		a.Balance -= amount
		to.Balance += amount
		a.Status = "Active"
		to.Status = "Active"
	} else if amount > 0 && a.Balance < amount {
		return errors.New("insufficient balance for transfer")
	}
	return nil
}

// Helper functions for benchmarking
func ProcessBalanceByValue(a Account, amount float64) float64 {
	a.Balance += amount
	return a.Balance
}

func ProcessBalanceByReference(a *Account, amount float64) float64 {
	a.Balance += amount
	return a.Balance
}
