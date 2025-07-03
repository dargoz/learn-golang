package db

import (
	"context"
	"database/sql"

	internal "github.com/dargoz/day04/data/local/internal/db"
)

type Store interface {
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
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
