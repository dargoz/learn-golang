package impl

import (
	"errors"
	"fmt"

	"github.com/dargoz/simplebank/model"
)

type Deposit struct {
	Amount float64
}

func (d Deposit) Apply(acc *model.Account) error {
	if d.Amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	acc.Balance += d.Amount

	acc.Status = "Active"

	return nil
}
func (d Deposit) Description() string {
	return fmt.Sprintf("Deposit of %.2f", d.Amount)
}
