package sqlplus

import (
	"context"
	"database/sql/driver"
)

type Hook interface {
	ConnectorHook
	ConnHook
	TxHook
	StmtHook
}

type ConnHook interface {
	BeforeExecContext(ctx context.Context, query string, args []driver.NamedValue) (context.Context, error)
	AfterExecContext(ctx context.Context, query string, args []driver.NamedValue, r driver.Result, err error) (context.Context, driver.Result, error)
	BeforeBeginTx(ctx context.Context, opts driver.TxOptions) (context.Context, error)
	AfterBeginTx(ctx context.Context, opts driver.TxOptions, dd driver.Tx, err error) (context.Context, driver.Tx, error)
	BeforeQueryContext(ctx context.Context, query string, args []driver.NamedValue) (context.Context, error)
	AfterQueryContext(ctx context.Context, query string, args []driver.NamedValue, rows driver.Rows, err error) (context.Context, driver.Rows, error)
	BeforePrepareContext(ctx context.Context, query string) (context.Context, error)
	AfterPrepareContext(ctx context.Context, query string, s driver.Stmt, err error) (context.Context, driver.Stmt, error)
}

type ConnectorHook interface {
	BeforeConnect(ctx context.Context) (context.Context, error)
	AfterConnect(ctx context.Context, dc driver.Conn, err error) (context.Context, driver.Conn, error)
}

type TxHook interface {
	BeforeCommit(ctx context.Context) (context.Context, error)
	AfterCommit(ctx context.Context, err error) (context.Context, error)
	BeforeRollback(ctx context.Context) (context.Context, error)
	AfterRollback(ctx context.Context, err error) (context.Context, error)
}

type StmtHook interface {
	BeforeStmtQueryContext(ctx context.Context, query string, args []driver.NamedValue) (context.Context, error)
	AfterStmtQueryContext(ctx context.Context, query string, args []driver.NamedValue, rows driver.Rows, err error) (context.Context, driver.Rows, error)
	BeforeStmtExecContext(ctx context.Context, query string, args []driver.NamedValue) (context.Context, error)
	AfterStmtExecContext(ctx context.Context, query string, args []driver.NamedValue, r driver.Result, err error) (context.Context, driver.Result, error)
}
