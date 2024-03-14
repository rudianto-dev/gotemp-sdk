package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type IDatabase interface {
	GetMasterConnection() (*sqlx.DB, error)

	CreateDBTransaction(ctx context.Context) (tx *sqlx.Tx, err error)
	CommitDBTransaction(ctx context.Context, tx *sqlx.Tx) (err error)
	Write(ctx context.Context, tx *sqlx.Tx, query string, params map[string]interface{}) (err error)
	Read(ctx context.Context, query string, params map[string]interface{}) (rows *sqlx.Rows, err error)
}
