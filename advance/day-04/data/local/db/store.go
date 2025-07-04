package db

import (
	"context"
	"database/sql"

	internal "github.com/dargoz/day04/data/local/internal/db"
)

type Store interface {
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	GetTransferByID(ctx context.Context, id int64) (TransferTxResult, error)
	CreateAccountTx(ctx context.Context, arg CreateAccountTxParams) (CreateAccountTxResult, error)
	internal.Querier
}

// SQLStore provides all functions to execute SQL queries and transactions
// on the database.

type SQLStore struct {
	db *sql.DB
	*internal.Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: internal.New(db),
	}
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
