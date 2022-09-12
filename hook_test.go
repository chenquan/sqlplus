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

func (m *mockHook) AfterExecContext(ctx context.Context, query string, args []driver.NamedValue, r driver.Result, err error) (context.Context, driver.Result, error) {
	m.Write("AfterExecContext")
	return ctx, nil, err
}

func (m *mockHook) BeforeBeginTx(ctx context.Context, opts driver.TxOptions, err error) (context.Context, driver.TxOptions, error) {
	m.Write("BeforeBeginTx")
	return ctx, opts, err
}

func (m *mockHook) AfterBeginTx(ctx context.Context, opts driver.TxOptions, dd driver.Tx, err error) (context.Context, driver.Tx, error) {
	m.Write("AfterBeginTx")
	return ctx, dd, err

}

func (m *mockHook) BeforeQueryContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, string, []driver.NamedValue, error) {
	m.Write("BeforeQueryContext")
	return ctx, query, args, err
}

func (m *mockHook) AfterQueryContext(ctx context.Context, query string, args []driver.NamedValue, rows driver.Rows, err error) (context.Context, driver.Rows, error) {
	m.Write("AfterQueryContext")
	return ctx, rows, err
}

func (m *mockHook) BeforePrepareContext(ctx context.Context, query string, err error) (context.Context, string, error) {
	m.Write("BeforePrepareContext")
	return ctx, query, err
}

func (m *mockHook) AfterPrepareContext(ctx context.Context, query string, s driver.Stmt, err error) (context.Context, driver.Stmt, error) {
	m.Write("AfterPrepareContext")
	return ctx, s, err
}

func (m *mockHook) BeforeCommit(ctx context.Context, err error) (context.Context, error) {
	m.Write("BeforeCommit")
	return ctx, err
}

func (m *mockHook) AfterCommit(ctx context.Context, err error) (context.Context, error) {
	m.Write("AfterCommit")
	return ctx, err
}

func (m *mockHook) BeforeRollback(ctx context.Context, err error) (context.Context, error) {
	m.Write("BeforeRollback")
	return ctx, err
}

func (m *mockHook) AfterRollback(ctx context.Context, err error) (context.Context, error) {
	m.Write("AfterRollback")
	return ctx, err

}

func (m *mockHook) BeforeStmtQueryContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, string, []driver.NamedValue, error) {
	m.Write("BeforeStmtQueryContext")
	return ctx, query, args, err
}

func (m *mockHook) AfterStmtQueryContext(ctx context.Context, query string, args []driver.NamedValue, rows driver.Rows, err error) (context.Context, driver.Rows, error) {
	m.Write("AfterStmtQueryContext")
	return ctx, rows, err
}

func (m *mockHook) BeforeStmtExecContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, string, []driver.NamedValue, error) {
	m.Write("BeforeStmtExecContext")
	return ctx, query, args, err
}

func (m *mockHook) AfterStmtExecContext(ctx context.Context, query string, args []driver.NamedValue, r driver.Result, err error) (context.Context, driver.Result, error) {
	m.Write("AfterStmtExecContext")
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
