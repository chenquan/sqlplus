package sqlplus

import (
	"context"
	"database/sql/driver"
)

type multiHook struct {
	hooks []Hook
}

func NewMultiHook(hooks ...Hook) Hook {
	return &multiHook{hooks: hooks}
}

func (h *multiHook) BeforeConnect(ctx context.Context, err error) (context.Context, error) {
	for _, hook := range h.hooks {
		ctx, err = hook.BeforeConnect(ctx, err)
	}

	return ctx, err
}

func (h *multiHook) AfterConnect(ctx context.Context, dc driver.Conn, err error) (context.Context, driver.Conn, error) {
	for _, hook := range h.hooks {
		ctx, dc, err = hook.AfterConnect(ctx, dc, err)
	}

	return ctx, dc, err
}

func (h *multiHook) BeforeExecContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, string, []driver.NamedValue, error) {
	for _, hook := range h.hooks {
		ctx, query, args, err = hook.BeforeExecContext(ctx, query, args, err)
	}

	return ctx, query, args, err
}

func (h *multiHook) AfterExecContext(ctx context.Context, query string, args []driver.NamedValue, r driver.Result, err error) (context.Context, driver.Result, error) {
	for _, hook := range h.hooks {
		ctx, r, err = hook.AfterExecContext(ctx, query, args, r, err)
	}

	return ctx, r, err
}

func (h *multiHook) BeforeBeginTx(ctx context.Context, opts driver.TxOptions, err error) (context.Context, driver.TxOptions, error) {
	for _, hook := range h.hooks {
		ctx, opts, err = hook.BeforeBeginTx(ctx, opts, err)
	}

	return ctx, opts, err
}

func (h *multiHook) AfterBeginTx(ctx context.Context, opts driver.TxOptions, dd driver.Tx, err error) (context.Context, driver.Tx, error) {
	for _, hook := range h.hooks {
		ctx, dd, err = hook.AfterBeginTx(ctx, opts, dd, err)
	}

	return ctx, dd, err
}

func (h *multiHook) BeforeQueryContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, string, []driver.NamedValue, error) {
	for _, hook := range h.hooks {
		ctx, query, args, err = hook.BeforeQueryContext(ctx, query, args, err)
	}

	return ctx, query, args, err
}

func (h *multiHook) AfterQueryContext(ctx context.Context, query string, args []driver.NamedValue, rows driver.Rows, err error) (context.Context, driver.Rows, error) {
	for _, hook := range h.hooks {
		ctx, rows, err = hook.AfterQueryContext(ctx, query, args, rows, err)

	}

	return ctx, rows, err
}

func (h *multiHook) BeforePrepareContext(ctx context.Context, query string, err error) (context.Context, string, error) {
	for _, hook := range h.hooks {
		ctx, query, err = hook.BeforePrepareContext(ctx, query, err)
	}

	return ctx, query, err
}

func (h *multiHook) AfterPrepareContext(ctx context.Context, query string, s driver.Stmt, err error) (context.Context, driver.Stmt, error) {
	for _, hook := range h.hooks {
		ctx, s, err = hook.AfterPrepareContext(ctx, query, s, err)
	}

	return ctx, s, err
}

func (h *multiHook) BeforeCommit(ctx context.Context, err error) (context.Context, error) {
	for _, hook := range h.hooks {
		ctx, err = hook.BeforeCommit(ctx, err)
	}

	return ctx, err
}

func (h *multiHook) AfterCommit(ctx context.Context, err error) (context.Context, error) {
	for _, hook := range h.hooks {
		ctx, err = hook.AfterCommit(ctx, err)
	}

	return ctx, err
}

func (h *multiHook) BeforeRollback(ctx context.Context, err error) (context.Context, error) {
	for _, hook := range h.hooks {
		ctx, err = hook.BeforeRollback(ctx, err)
	}

	return ctx, err
}

func (h *multiHook) AfterRollback(ctx context.Context, err error) (context.Context, error) {
	for _, hook := range h.hooks {
		ctx, err = hook.AfterRollback(ctx, err)
	}

	return ctx, err
}

func (h *multiHook) BeforeStmtQueryContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, []driver.NamedValue, error) {
	for _, hook := range h.hooks {
		ctx, args, err = hook.BeforeStmtQueryContext(ctx, query, args, err)
	}

	return ctx, args, err
}

func (h *multiHook) AfterStmtQueryContext(ctx context.Context, query string, args []driver.NamedValue, rows driver.Rows, err error) (context.Context, driver.Rows, error) {
	for _, hook := range h.hooks {
		ctx, rows, err = hook.AfterStmtQueryContext(ctx, query, args, rows, err)
	}

	return ctx, rows, err
}

func (h *multiHook) BeforeStmtExecContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, []driver.NamedValue, error) {
	for _, hook := range h.hooks {
		ctx, args, err = hook.BeforeStmtExecContext(ctx, query, args, err)
	}

	return ctx, args, err
}

func (h *multiHook) AfterStmtExecContext(ctx context.Context, query string, args []driver.NamedValue, r driver.Result, err error) (context.Context, driver.Result, error) {
	for _, hook := range h.hooks {
		ctx, r, err = hook.AfterStmtExecContext(ctx, query, args, r, err)
	}

	return ctx, r, err
}
