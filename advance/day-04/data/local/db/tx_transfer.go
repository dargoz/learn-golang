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
	Transfer internal.Transfer
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*internal.Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := internal.New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit()
}

func (store *SQLStore) AddAccountBalance(ctx context.Context, arg internal.AddAccountBalanceParams) (internal.Account, error) {
	row := store.db.QueryRowContext(ctx, internal., arg.ID, arg.Balance)
	var i internal.Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

// TransferTx implements the Store interface.
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	// TODO: implement the transaction logic here
	var result TransferTxResult

	err := store.execTx(ctx, func(q *internal.Queries) error {
		// 1. Check if from account exists and is locked
		fromAccount, err := q.GetAccountForUpdate(ctx, arg.FromAccountID)
		if err != nil {

		}
		if fromAccount.Balance < arg.Amount {
			return internal.ErrInsufficientFunds
		}
		// execute create transfer
		transfer, err := q.CreateTransfer(ctx, internal.CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
	})
	return TransferTxResult{}, nil
}
