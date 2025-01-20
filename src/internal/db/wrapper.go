package db

import (
	"context"
	"database/sql"
)

type QuerierWithTx interface {
	Querier // Embeds sqlc-generated interface

	BeginTx(ctx context.Context) (QuerierWithTx, error)
	Commit() error
	Rollback() error
}

// DBTx wraps sql.DB and sql.Tx to implement QuerierWithTx
type DBTx struct {
	db *sql.DB
	tx *sql.Tx
	*Queries
}

// NewDB creates a new instance of DBTx with a sql.DB connection
func NewDB(db *sql.DB) *DBTx {
	return &DBTx{
		db:      db,
		Queries: New(db), // New() is from sqlc
	}
}

// BeginTx starts a new transaction
func (d *DBTx) BeginTx(ctx context.Context) (QuerierWithTx, error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &DBTx{
		db:      d.db,
		tx:      tx,
		Queries: New(tx), // Use sqlc queries with the transaction
	}, nil
}

// Commit commits the transaction
func (d *DBTx) Commit() error {
	if d.tx == nil {
		return nil
	}
	return d.tx.Commit()
}

// Rollback rolls back the transaction
func (d *DBTx) Rollback() error {
	if d.tx == nil {
		return nil
	}
	return d.tx.Rollback()
}
