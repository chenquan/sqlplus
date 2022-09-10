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
	AfterExecContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, error)
	BeforeBeginTx(ctx context.Context, opts driver.TxOptions) (context.Context, error)
	AfterBeginTx(ctx context.Context, opts driver.TxOptions, err error) (context.Context, error)
	BeforeQueryContext(ctx context.Context, query string, args []driver.NamedValue) (context.Context, error)
	AfterQueryContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, error)
	BeforePrepareContext(ctx context.Context, query string) (context.Context, error)
	AfterPrepareContext(ctx context.Context, query string, err error) (context.Context, error)
}

type ConnectorHook interface {
	BeforeConnect(ctx context.Context) (context.Context, error)
	AfterConnect(ctx context.Context, err error) (context.Context, error)
}

type TxHook interface {
	BeforeCommit(ctx context.Context) (context.Context, error)
	AfterCommit(ctx context.Context, err error) (context.Context, error)
	BeforeRollback(ctx context.Context) (context.Context, error)
	AfterRollback(ctx context.Context, err error) (context.Context, error)
}

type StmtHook interface {
	BeforeStmtQueryContext(ctx context.Context, query string, args []driver.NamedValue) (context.Context, error)
	AfterStmtQueryContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, error)
	BeforeStmtExecContext(ctx context.Context, query string, args []driver.NamedValue) (context.Context, error)
	AfterStmtExecContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, error)
}
