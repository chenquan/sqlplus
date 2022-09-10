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

func (m *mockHook) Reset() {
	m.Args = nil
}

func (m *mockHook) Write(s string) {
	m.Args = append(m.Args, s)
}

func (m *mockHook) String() string {
	return strings.Join(m.Args, "|")
}

func (m *mockHook) BeforeConnect(ctx context.Context) (context.Context, error) {
	m.Write("BeforeConnect")
	return ctx, nil
}

func (m *mockHook) AfterConnect(ctx context.Context, err error) (context.Context, error) {
	m.Write("AfterConnect")
	return ctx, nil
}

func (m *mockHook) BeforeExecContext(ctx context.Context, query string, args []driver.NamedValue) (context.Context, error) {
	m.Write("BeforeExecContext")
	return ctx, nil
}

func (m *mockHook) AfterExecContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, error) {
	m.Write("AfterExecContext")
	return ctx, nil
}

func (m *mockHook) BeforeBeginTx(ctx context.Context, opts driver.TxOptions) (context.Context, error) {
	m.Write("BeforeBeginTx")
	return ctx, nil
}

func (m *mockHook) AfterBeginTx(ctx context.Context, opts driver.TxOptions, err error) (context.Context, error) {
	m.Write("AfterBeginTx")
	return ctx, nil
}

func (m *mockHook) BeforeQueryContext(ctx context.Context, query string, args []driver.NamedValue) (context.Context, error) {
	m.Write("BeforeQueryContext")
	return ctx, nil
}

func (m *mockHook) AfterQueryContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, error) {
	m.Write("AfterQueryContext")
	return ctx, nil
}

func (m *mockHook) BeforePrepareContext(ctx context.Context, query string) (context.Context, error) {
	m.Write("BeforePrepareContext")
	return ctx, nil
}

func (m *mockHook) AfterPrepareContext(ctx context.Context, query string, err error) (context.Context, error) {
	m.Write("AfterPrepareContext")
	return ctx, nil
}

func (m *mockHook) BeforeCommit(ctx context.Context) (context.Context, error) {
	m.Write("BeforeCommit")
	return ctx, nil
}

func (m *mockHook) AfterCommit(ctx context.Context, err error) (context.Context, error) {
	m.Write("AfterCommit")
	return ctx, nil
}

func (m *mockHook) BeforeRollback(ctx context.Context) (context.Context, error) {
	m.Write("BeforeRollback")
	return ctx, nil
}

func (m *mockHook) AfterRollback(ctx context.Context, err error) (context.Context, error) {
	m.Write("AfterRollback")
	return ctx, nil
}

func (m *mockHook) BeforeStmtQueryContext(ctx context.Context, query string, args []driver.NamedValue) (context.Context, error) {
	m.Write("BeforeStmtQueryContext")
	return ctx, nil
}

func (m *mockHook) AfterStmtQueryContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, error) {
	m.Write("AfterStmtQueryContext")
	return ctx, nil
}

func (m *mockHook) BeforeStmtExecContext(ctx context.Context, query string, args []driver.NamedValue) (context.Context, error) {
	m.Write("BeforeStmtExecContext")
	return ctx, nil
}

func (m *mockHook) AfterStmtExecContext(ctx context.Context, query string, args []driver.NamedValue, err error) (context.Context, error) {
	m.Write("AfterStmtExecContext")
	return ctx, nil
}
