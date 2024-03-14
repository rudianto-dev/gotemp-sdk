package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// CreateDBTransaction creates new database transaction
func (s *DB) CreateDBTransaction(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = s.Master.BeginTxx(ctx, nil)
	if err != nil {
		logging := map[string]interface{}{
			"scope":  s.Scope,
			"action": "CreateDBTransaction",
			"error":  err,
			"stage":  "DB",
		}
		s.Logger.ErrorContextWithField(ctx, logging)
		return nil, err
	}

	return tx, nil
}

// CommitDBTransaction commits database transaction
func (s *DB) CommitDBTransaction(ctx context.Context, tx *sqlx.Tx) (err error) {
	err = tx.Commit()
	if err != nil {
		logging := map[string]interface{}{
			"scope":  s.Scope,
			"action": "CommitDBTransaction",
			"error":  err,
			"stage":  "DB",
		}
		s.Logger.ErrorContextWithField(ctx, logging)
		return err
	}

	return
}

// Write database transaction
func (s *DB) Write(ctx context.Context, tx *sqlx.Tx, query string, params map[string]interface{}) (err error) {
	// Create new sql transaction if no sql transaction exist
	sqlTx := tx
	if tx == nil {
		sqlTx, err = s.CreateDBTransaction(ctx)
		if err != nil {
			return
		}
	}
	var args []interface{}
	query, args, err = sqlx.Named(query, params)
	if err != nil {
		return
	}
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	_, err = sqlTx.ExecContext(ctx, query, args...)
	if s.Debug {
		logging := map[string]interface{}{
			"scope":        s.Scope,
			"action":       "write",
			"query":        query,
			"query_params": args,
			"error":        err,
			"stage":        "DB",
		}
		s.Logger.InfoContextWithField(ctx, logging)
	}
	if err != nil {
		if err == context.Canceled {
			err = nil
			return
		}
		logging := map[string]interface{}{
			"scope":        s.Scope,
			"action":       "write",
			"query":        query,
			"query_params": args,
			"error":        err,
			"stage":        "DB",
		}
		s.Logger.ErrorContextWithField(ctx, logging)
		return
	}

	// Commit sql transaction immediately if no sql transaction exist
	if tx == nil {
		err = s.CommitDBTransaction(ctx, sqlTx)
		if err != nil {
			return
		}
	}

	return
}

// Read database data
func (s *DB) Read(ctx context.Context, query string, params map[string]interface{}) (rows *sqlx.Rows, err error) {
	var args []interface{}
	query, args, err = sqlx.Named(query, params)
	if err != nil {
		return
	}
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	rows, err = s.Slave.QueryxContext(ctx, query, args...)
	if s.Debug {
		logging := map[string]interface{}{
			"scope":        s.Scope,
			"action":       "read",
			"query":        query,
			"query_params": args,
			"error":        err,
			"stage":        "DB",
		}
		s.Logger.InfoContextWithField(ctx, logging)
	}
	if err != nil {
		if err == context.Canceled {
			err = nil
			return
		}
		logging := map[string]interface{}{
			"scope":        s.Scope,
			"action":       "read",
			"query":        query,
			"query_params": args,
			"error":        err,
			"stage":        "DB",
		}
		s.Logger.ErrorContextWithField(ctx, logging)
		return
	}

	return
}
