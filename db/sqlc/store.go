package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store can do both queries and transcations
type Store interface {
	Querier
}

// SQLStore provides all funcs for SQL queries and transactions
type SQLStore struct {
	// queries only supports queries not transactions,
	// so we use it in store struct and add more functionality.
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
