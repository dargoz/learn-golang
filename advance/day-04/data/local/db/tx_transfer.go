package db

import (
	"context"

	internal "github.com/dargoz/day04/data/local/internal/db"
)

type TransferTxParams struct {
	FromAccountID int64
	ToAccountID   int64
	Amount        int64
}

type TransferTxResult struct {
	Transfer    internal.Transfer
	FromAccount internal.Account
	ToAccount   internal.Account
}

// TransferTx implements the Store interface.
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *internal.Queries) error {
		var err error

		// 1. Buat transfer record
		result.Transfer, err = q.CreateTransfer(ctx, internal.CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// 2. Update saldo dari pengirim
		result.FromAccount, err = q.AddAccountBalance(ctx, internal.AddAccountBalanceParams{
			ID:      arg.FromAccountID,
			Balance: -arg.Amount,
		})
		if err != nil {
			return err
		}

		// 3. Update saldo ke penerima
		result.ToAccount, err = q.AddAccountBalance(ctx, internal.AddAccountBalanceParams{
			ID:      arg.ToAccountID,
			Balance: arg.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

func (store *SQLStore) GetTransferByID(ctx context.Context, id int64) (TransferTxResult, error) {
	var result TransferTxResult

	error := store.execTx(ctx, func(q *internal.Queries) error {
		var err error

		result.Transfer, err = q.GetTransfer(ctx, id)
		if err != nil {
			return err
		}

		result.FromAccount, err = q.GetAccount(ctx, result.Transfer.FromAccountID)
		if err != nil {
			return err
		}

		result.ToAccount, err = q.GetAccount(ctx, result.Transfer.ToAccountID)
		if err != nil {
			return err
		}

		return nil
	})

	if error != nil {
		return result, error
	}
	return result, nil
}
