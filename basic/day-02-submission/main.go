package main

import (
	"fmt"

	"github.com/dargoz/simplebank/domain"
	"github.com/dargoz/simplebank/impl"
	"github.com/dargoz/simplebank/model"
)

func main() {

	var sourceAccount = &model.Account{
		ID:            1,
		Name:          "Source Account",
		Balance:       1000.0,
		Status:        "Active",
		AccountNumber: "123456789",
	}
	var destinationAccount = &model.Account{
		ID:            2,
		Name:          "Destination Account",
		Balance:       500.0,
		Status:        "Active",
		AccountNumber: "987654321",
	}

	var deposit domain.Transaction = impl.Deposit{Amount: 100.0}

	var withdrawal domain.Transaction = impl.Withdrawal{Amount: 100.0}

	deposit.Apply(destinationAccount)
	fmt.Println(deposit.Description())
	fmt.Printf("After deposit, destination account balance: %.2f\n", destinationAccount.Balance)

	withdrawal.Apply(sourceAccount)
	fmt.Println(withdrawal.Description())
	fmt.Printf("After withdrawal, source account balance: %.2f\n", sourceAccount.Balance)

}
