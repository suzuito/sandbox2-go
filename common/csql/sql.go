package csql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/suzuito/sandbox2-go/common/terrors"
)

type SQLError struct {
	Query         string
	Args          []any
	OriginalError error
}

func (t *SQLError) Error() string {
	return fmt.Sprintf("Query: [%s], Args: %+v, Error: %s", t.Query, t.Args, t.OriginalError)
}

func (t *SQLError) Unwrap() error {
	return t.OriginalError
}

type TxOrDB interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

func SqlIn[T []E, E any](fieldName string, cols T) string {
	s := make([]byte, len(cols)*2-1)
	for i := range cols {
		j := i * 2
		s[j] = '?'
		if i >= len(cols)-1 {
			break
		}
		s[j+1] = ','
	}
	return fmt.Sprintf("`%s` IN (%s)", fieldName, string(s))
}

func ToAnySlice[T []E, E any](cols T) []any {
	ret := []any{}
	for _, e := range cols {
		ret = append(ret, e)
	}
	return ret
}

func QueryContext(
	ctx context.Context,
	txOrDB TxOrDB,
	query string,
	args ...any,
) (*sql.Rows, error) {
	rows, err := txOrDB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, &SQLError{
			Query:         query,
			Args:          args,
			OriginalError: err,
		}
	}
	return rows, nil
}

func QueryRowContext(
	ctx context.Context,
	txOrDB TxOrDB,
	query string,
	args ...any,
) *sql.Row {
	row := txOrDB.QueryRowContext(ctx, query, args...)
	return row
}

func ExecContext(
	ctx context.Context,
	txOrDB TxOrDB,
	query string,
	args ...any,
) (sql.Result, error) {
	result, err := txOrDB.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, &SQLError{
			Query:         query,
			Args:          args,
			OriginalError: err,
		}
	}
	return result, nil
}

func WithTransaction(ctx context.Context, db *sql.DB, f func(tx TxOrDB) error) error {
	tx, err := db.Begin()
	if err != nil {
		return terrors.Wrap(err)
	}
	if err := f(tx); err != nil {
		if err := tx.Rollback(); err != nil {
			return terrors.Wrap(err)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
