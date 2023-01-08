package sqlplus

import (
	"context"
	"database/sql/driver"
)

type (
	Hook interface {
		ConnectorHook
		ConnHook
		TxHook
		StmtHook
	}
	ConnHook interface {
		BeforeExecContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, string, []driver.NamedValue, error)
		AfterExecContext(ctx context.Context, query string, args []driver.NamedValue, dr driver.Result, err error) (context.Context, driver.Result, error)

		BeforeBeginTx(ctx context.Context, opts driver.TxOptions, err error) (context.Context, driver.TxOptions, error)
		AfterBeginTx(ctx context.Context, opts driver.TxOptions, dt driver.Tx, err error) (context.Context, driver.Tx, error)

		BeforeQueryContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, string, []driver.NamedValue, error)
		AfterQueryContext(ctx context.Context, query string, args []driver.NamedValue, rows driver.Rows, err error) (context.Context, driver.Rows, error)

		BeforePrepareContext(ctx context.Context, query string, err error) (context.Context, string, error)
		AfterPrepareContext(ctx context.Context, query string, ds driver.Stmt, err error) (context.Context, driver.Stmt, error)

		BeforeClose(ctx context.Context, err error) (context.Context, error)
		AfterClose(ctx context.Context, err error) (context.Context, error)
	}
	ConnectorHook interface {
		BeforeConnect(ctx context.Context, err error) (context.Context, error)
		AfterConnect(ctx context.Context, dc driver.Conn, err error) (context.Context, driver.Conn, error)
	}
	TxHook interface {
		BeforeCommit(ctx context.Context, err error) (context.Context, error)
		AfterCommit(ctx context.Context, err error) (context.Context, error)

		BeforeRollback(ctx context.Context, err error) (context.Context, error)
		AfterRollback(ctx context.Context, err error) (context.Context, error)
	}
	StmtHook interface {
		BeforeStmtQueryContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, []driver.NamedValue, error)
		AfterStmtQueryContext(ctx context.Context, query string, args []driver.NamedValue, rows driver.Rows, err error) (context.Context, driver.Rows, error)

		BeforeStmtExecContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, []driver.NamedValue, error)
		AfterStmtExecContext(ctx context.Context, query string, args []driver.NamedValue, r driver.Result, err error) (context.Context, driver.Result, error)
	}
)
