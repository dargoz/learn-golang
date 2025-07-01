package impl

import (
	"errors"
	"fmt"

	"github.com/dargoz/simplebank/model"
)

type Withdrawal struct {
	Amount float64
}

func (w Withdrawal) Apply(acc *model.Account) error {
	if w.Amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	if acc.Balance < w.Amount {
		return errors.New("insufficient funds for withdrawal")
	}
	acc.Balance -= w.Amount

	if acc.Balance == 0 {
		acc.Status = "Inactive"
	}

	return nil
}

func (w Withdrawal) Description() string {
	return fmt.Sprintf("Withdrawal of %.2f", w.Amount)
}
