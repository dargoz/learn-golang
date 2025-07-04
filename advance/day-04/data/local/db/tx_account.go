package db

import (
	"context"

	internal "github.com/dargoz/day04/data/local/internal/db"
)

type CreateAccountTxParams struct {
	Owner    string
	Balance  int64
	Currency string
}

type CreateAccountTxResult struct {
	Account internal.Account
}

func (sql *SQLStore) CreateAccountTx(ctx context.Context, arg CreateAccountTxParams) (CreateAccountTxResult, error) {
	var result CreateAccountTxResult

	err := sql.execTx(ctx, func(q *internal.Queries) error {
		var err error

		result.Account, err = q.CreateAccount(ctx, internal.CreateAccountParams{
			Owner:    arg.Owner,
			Balance:  arg.Balance,
			Currency: arg.Currency,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
