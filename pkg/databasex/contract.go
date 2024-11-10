package databasex

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Adapter interface {
	Ping() error
	//InTransaction() bool
	Close() error
	Query(ctx context.Context, dst any, query string, args ...any) error
	QueryRow(ctx context.Context, dst any, query string, args ...any) error
	QueryX(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowX(ctx context.Context, query string, args ...any) *sql.Row
	Exec(ctx context.Context, query string, args ...any) (_ int64, err error)
	Transact(ctx context.Context, iso sql.IsolationLevel, txFunc func(*DB) error) (err error)
	PrepareNamedContext(ctx context.Context, query string) (*NamedStmt, error)
	PreparedNameContextForRead(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	PreparedNameContextForWrite(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	DBRead() Adapter
	DBWrite() Adapter
	BeginTx(ctx context.Context, iso sql.IsolationLevel) Adapter
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error

	Rebind(ctx context.Context, q string) string
	ParseSQLError(err error) error
}

type AdapterPrepare interface {
	QueryxContext(ctx context.Context, arg interface{}) (*sqlx.Rows, error)
	QueryRowxContext(ctx context.Context, arg interface{}) *sqlx.Row
	ExecContext(ctx context.Context, arg interface{}) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, arg interface{}) error
}
