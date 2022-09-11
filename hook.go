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
		AfterExecContext(ctx context.Context, query string, args []driver.NamedValue, r driver.Result, err error) (context.Context, driver.Result, error)

		BeforeBeginTx(ctx context.Context, opts driver.TxOptions, err error) (context.Context, driver.TxOptions, error)
		AfterBeginTx(ctx context.Context, opts driver.TxOptions, dd driver.Tx, err error) (context.Context, driver.Tx, error)

		BeforeQueryContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, string, []driver.NamedValue, error)
		AfterQueryContext(ctx context.Context, query string, args []driver.NamedValue, rows driver.Rows, err error) (context.Context, driver.Rows, error)

		BeforePrepareContext(ctx context.Context, query string, err error) (context.Context, string, error)
		AfterPrepareContext(ctx context.Context, query string, s driver.Stmt, err error) (context.Context, driver.Stmt, error)
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
	Hooks struct {
		hooks []Hook
	}
)

func NewMultiHook(hooks ...Hook) Hook {
	return &Hooks{hooks: hooks}
}

func (h *Hooks) BeforeConnect(ctx context.Context, err error) (context.Context, error) {
	for _, hook := range h.hooks {
		ctx, err = hook.BeforeConnect(ctx, err)
	}

	return ctx, err
}

func (h *Hooks) AfterConnect(ctx context.Context, dc driver.Conn, err error) (context.Context, driver.Conn, error) {
	for _, hook := range h.hooks {
		ctx, dc, err = hook.AfterConnect(ctx, dc, err)
	}

	return ctx, dc, err
}

func (h *Hooks) BeforeExecContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, string, []driver.NamedValue, error) {
	for _, hook := range h.hooks {
		ctx, query, args, err = hook.BeforeExecContext(ctx, query, args, err)
	}

	return ctx, query, args, err
}

func (h *Hooks) AfterExecContext(ctx context.Context, query string, args []driver.NamedValue, r driver.Result, err error) (context.Context, driver.Result, error) {
	for _, hook := range h.hooks {
		ctx, r, err = hook.AfterExecContext(ctx, query, args, r, err)
	}

	return ctx, r, err
}

func (h *Hooks) BeforeBeginTx(ctx context.Context, opts driver.TxOptions, err error) (context.Context, driver.TxOptions, error) {
	for _, hook := range h.hooks {
		ctx, opts, err = hook.BeforeBeginTx(ctx, opts, err)
	}

	return ctx, opts, err
}

func (h *Hooks) AfterBeginTx(ctx context.Context, opts driver.TxOptions, dd driver.Tx, err error) (context.Context, driver.Tx, error) {
	for _, hook := range h.hooks {
		ctx, dd, err = hook.AfterBeginTx(ctx, opts, dd, err)
	}

	return ctx, dd, err
}

func (h *Hooks) BeforeQueryContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, string, []driver.NamedValue, error) {
	for _, hook := range h.hooks {
		ctx, query, args, err = hook.BeforeQueryContext(ctx, query, args, err)
	}

	return ctx, query, args, err
}

func (h *Hooks) AfterQueryContext(ctx context.Context, query string, args []driver.NamedValue, rows driver.Rows, err error) (context.Context, driver.Rows, error) {
	for _, hook := range h.hooks {
		ctx, rows, err = hook.AfterQueryContext(ctx, query, args, rows, err)

	}

	return ctx, rows, err
}

func (h *Hooks) BeforePrepareContext(ctx context.Context, query string, err error) (context.Context, string, error) {
	for _, hook := range h.hooks {
		ctx, query, err = hook.BeforePrepareContext(ctx, query, err)
	}

	return ctx, query, err
}

func (h *Hooks) AfterPrepareContext(ctx context.Context, query string, s driver.Stmt, err error) (context.Context, driver.Stmt, error) {
	for _, hook := range h.hooks {
		ctx, s, err = hook.AfterPrepareContext(ctx, query, s, err)
	}

	return ctx, s, err
}

func (h *Hooks) BeforeCommit(ctx context.Context, err error) (context.Context, error) {
	for _, hook := range h.hooks {
		ctx, err = hook.BeforeCommit(ctx, err)
	}

	return ctx, err
}

func (h *Hooks) AfterCommit(ctx context.Context, err error) (context.Context, error) {
	for _, hook := range h.hooks {
		ctx, err = hook.AfterCommit(ctx, err)
	}

	return ctx, err
}

func (h *Hooks) BeforeRollback(ctx context.Context, err error) (context.Context, error) {
	for _, hook := range h.hooks {
		ctx, err = hook.BeforeRollback(ctx, err)
	}

	return ctx, err
}

func (h *Hooks) AfterRollback(ctx context.Context, err error) (context.Context, error) {
	for _, hook := range h.hooks {
		ctx, err = hook.AfterRollback(ctx, err)
	}

	return ctx, err
}

func (h *Hooks) BeforeStmtQueryContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, []driver.NamedValue, error) {
	for _, hook := range h.hooks {
		ctx, args, err = hook.BeforeStmtQueryContext(ctx, query, args, err)
	}

	return ctx, args, err
}

func (h *Hooks) AfterStmtQueryContext(ctx context.Context, query string, args []driver.NamedValue, rows driver.Rows, err error) (context.Context, driver.Rows, error) {
	for _, hook := range h.hooks {
		ctx, rows, err = hook.AfterStmtQueryContext(ctx, query, args, rows, err)
	}

	return ctx, rows, err
}

func (h *Hooks) BeforeStmtExecContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, []driver.NamedValue, error) {
	for _, hook := range h.hooks {
		ctx, args, err = hook.BeforeStmtExecContext(ctx, query, args, err)
	}

	return ctx, args, err
}

func (h *Hooks) AfterStmtExecContext(ctx context.Context, query string, args []driver.NamedValue, r driver.Result, err error) (context.Context, driver.Result, error) {
	for _, hook := range h.hooks {
		ctx, r, err = hook.AfterStmtExecContext(ctx, query, args, r, err)
	}

	return ctx, r, err
}
