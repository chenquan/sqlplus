package sqlplus

import (
	"context"
	"database/sql/driver"
	"strings"
)

var _ Hook = (*mockHook)(nil)

type mockHook struct {
	Args []string
}

func (m *mockHook) BeforeClose(ctx context.Context, err error) (context.Context, error) {
	m.Write("BeforeClose")
	return ctx, err
}

func (m *mockHook) AfterClose(ctx context.Context, err error) (context.Context, error) {
	m.Write("AfterClose")
	return ctx, err
}

func (m *mockHook) BeforeConnect(ctx context.Context, err error) (context.Context, error) {
	m.Write("BeforeConnect")
	return ctx, err
}

func (m *mockHook) AfterConnect(ctx context.Context, dc driver.Conn, err error) (context.Context, driver.Conn, error) {
	m.Write("AfterConnect")
	return ctx, dc, err
}

func (m *mockHook) BeforeExecContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, string, []driver.NamedValue, error) {
	m.Write("BeforeExecContext")
	return ctx, query, args, err
}

func (m *mockHook) AfterExecContext(ctx context.Context, _ string, _ []driver.NamedValue, _ driver.Result, err error) (context.Context, driver.Result, error) {
	m.Write("AfterExecContext")
	return ctx, nil, err
}

func (m *mockHook) BeforeBeginTx(ctx context.Context, opts driver.TxOptions, err error) (context.Context, driver.TxOptions, error) {
	m.Write("BeforeBeginTx")
	return ctx, opts, err
}

func (m *mockHook) AfterBeginTx(ctx context.Context, _ driver.TxOptions, dt driver.Tx, err error) (context.Context, driver.Tx, error) {
	m.Write("AfterBeginTx")
	return ctx, dt, err
}

func (m *mockHook) BeforeQueryContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, string, []driver.NamedValue, error) {
	m.Write("BeforeQueryContext")
	return ctx, query, args, err
}

func (m *mockHook) AfterQueryContext(ctx context.Context, _ string, _ []driver.NamedValue, rows driver.Rows, err error) (context.Context, driver.Rows, error) {
	m.Write("AfterQueryContext")
	return ctx, rows, err
}

func (m *mockHook) BeforePrepareContext(ctx context.Context, query string, err error) (context.Context, string, error) {
	m.Write("BeforePrepareContext")
	return ctx, query, err
}

func (m *mockHook) AfterPrepareContext(ctx context.Context, _ string, ds driver.Stmt, err error) (context.Context, driver.Stmt, error) {
	m.Write("AfterPrepareContext")
	return ctx, ds, err
}

func (m *mockHook) BeforeCommit(ctx context.Context, err error) (context.Context, error) {
	m.Write("BeforeCommit")

	prepareContext := PrepareContextFromContext(ctx)
	if prepareContext != nil {
		panic("prepareContext is not nil")
	}

	return ctx, err
}

func (m *mockHook) AfterCommit(ctx context.Context, err error) (context.Context, error) {
	m.Write("AfterCommit")

	prepareContext := PrepareContextFromContext(ctx)
	if prepareContext != nil {
		panic("prepareContext is not nil")
	}

	return ctx, err
}

func (m *mockHook) BeforeRollback(ctx context.Context, err error) (context.Context, error) {
	m.Write("BeforeRollback")

	prepareContext := PrepareContextFromContext(ctx)
	if prepareContext != nil {
		panic("prepareContext is not nil")
	}

	return ctx, err
}

func (m *mockHook) AfterRollback(ctx context.Context, err error) (context.Context, error) {
	m.Write("AfterRollback")

	prepareContext := PrepareContextFromContext(ctx)
	if prepareContext != nil {
		panic("prepareContext is not nil")
	}

	return ctx, err
}

func (m *mockHook) BeforeStmtQueryContext(ctx context.Context, _ string, args []driver.NamedValue, err error) (context.Context, []driver.NamedValue, error) {
	m.Write("BeforeStmtQueryContext")
	prepareContext := PrepareContextFromContext(ctx)
	if prepareContext == nil {
		panic("prepareContext is nil")
	}

	return ctx, args, err
}

func (m *mockHook) AfterStmtQueryContext(ctx context.Context, _ string, _ []driver.NamedValue, rows driver.Rows, err error) (context.Context, driver.Rows, error) {
	m.Write("AfterStmtQueryContext")
	prepareContext := PrepareContextFromContext(ctx)
	if prepareContext == nil {
		panic("prepareContext is nil")
	}

	return ctx, rows, err
}

func (m *mockHook) BeforeStmtExecContext(ctx context.Context, _ string, args []driver.NamedValue, err error) (context.Context, []driver.NamedValue, error) {
	m.Write("BeforeStmtExecContext")
	prepareContext := PrepareContextFromContext(ctx)
	if prepareContext == nil {
		panic("prepareContext is nil")
	}

	return ctx, args, err
}

func (m *mockHook) AfterStmtExecContext(ctx context.Context, _ string, _ []driver.NamedValue, r driver.Result, err error) (context.Context, driver.Result, error) {
	m.Write("AfterStmtExecContext")
	prepareContext := PrepareContextFromContext(ctx)
	if prepareContext == nil {
		panic("prepareContext is nil")
	}

	return ctx, r, err
}

func (m *mockHook) Reset() {
	m.Args = nil
}

func (m *mockHook) Write(s string) {
	m.Args = append(m.Args, s)
}

func (m *mockHook) String() string {
	return strings.Join(m.Args, "|")
}
