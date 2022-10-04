package sqlplus

import (
	"context"
	"database/sql/driver"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ driver.Conn = (*mockConn)(nil)

type mockConn struct {
}

func (c *mockConn) Prepare(_ string) (driver.Stmt, error) {
	return &mockStmt{}, nil
}

func (c *mockConn) Close() error {
	return nil
}

func (c *mockConn) Begin() (driver.Tx, error) {
	return &mockTx{}, nil
}

// -----------------

var _ driver.Queryer = (*mockConnQueryer)(nil)

type mockConnQueryer struct {
	driver.Conn
}

func (m *mockConnQueryer) Query(_ string, _ []driver.Value) (driver.Rows, error) {
	return nil, nil
}

// -----------------

var _ driver.QueryerContext = (*mockConnQueryerContext)(nil)

type mockConnQueryerContext struct {
	driver.Conn
}

func (m *mockConnExecer) Exec(_ string, _ []driver.Value) (driver.Result, error) {
	return nil, nil
}

// -----------------

var _ driver.Execer = (*mockConnExecer)(nil)

type mockConnExecer struct {
	driver.Conn
}

func (m *mockConnQueryerContext) QueryContext(_ context.Context, _ string, _ []driver.NamedValue) (driver.Rows, error) {
	return nil, nil
}

// -----------------
var _ driver.ExecerContext = (*mockConnExecerContext)(nil)

type mockConnExecerContext struct {
	driver.Conn
}

func (m *mockConnExecerContext) ExecContext(_ context.Context, _ string, _ []driver.NamedValue) (driver.Result, error) {
	return nil, nil
}

// -----------------
var _ driver.ConnBeginTx = (*mockConnBeginTx)(nil)

type mockConnBeginTx struct {
	driver.Conn
}

func (m *mockConnBeginTx) BeginTx(_ context.Context, _ driver.TxOptions) (driver.Tx, error) {
	return &mockTx{}, nil
}

var _ driver.ConnPrepareContext = (*mockConnPrepareContext)(nil)

type mockConnPrepareContext struct {
	*mockConn
}

func (m *mockConnPrepareContext) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	return &mockStmt{}, nil
}

// -----------------

func createMockConn() (*conn, *mockHook) {
	m := &mockHook{}
	hooks := NewMultiHook(m)
	c := &conn{
		Conn:     &mockConn{},
		ConnHook: hooks,
	}
	return c, m
}

func Test_conn_BeginTx(t *testing.T) {
	t.Run("Begin", func(t *testing.T) {
		c, m := createMockConn()

		tt, err := c.BeginTx(context.Background(), driver.TxOptions{
			Isolation: 0,
			ReadOnly:  false,
		})
		assert.NoError(t, err)
		assert.NoError(t, tt.Commit())
		assert.Contains(t, m.String(), "BeforeBeginTx")
		assert.Contains(t, m.String(), "AfterBeginTx")
		assert.Contains(t, m.String(), "BeforeCommit")
		assert.Contains(t, m.String(), "AfterCommit")

		m.Reset()
		tt, err = c.BeginTx(context.Background(), driver.TxOptions{
			Isolation: 0,
			ReadOnly:  false,
		})

		assert.NoError(t, err)
		assert.NoError(t, tt.Rollback())
		assert.Contains(t, m.String(), "BeforeBeginTx")
		assert.Contains(t, m.String(), "AfterBeginTx")
		assert.Contains(t, m.String(), "BeforeRollback")
		assert.Contains(t, m.String(), "AfterRollback")
	})

	t.Run("BeginTx", func(t *testing.T) {
		m := &mockHook{}
		c := &conn{
			Conn:     &mockConnBeginTx{Conn: &mockConn{}},
			ConnHook: m,
		}
		tt, err := c.BeginTx(context.Background(), driver.TxOptions{
			Isolation: 0,
			ReadOnly:  false,
		})
		assert.NoError(t, err)
		assert.NoError(t, tt.Commit())
		assert.Contains(t, m.String(), "BeforeBeginTx")
		assert.Contains(t, m.String(), "AfterBeginTx")
		assert.Contains(t, m.String(), "BeforeCommit")
		assert.Contains(t, m.String(), "AfterCommit")

		m.Reset()
		tt, err = c.BeginTx(context.Background(), driver.TxOptions{
			Isolation: 0,
			ReadOnly:  false,
		})

		assert.NoError(t, err)
		assert.NoError(t, tt.Rollback())
		assert.Contains(t, m.String(), "BeforeBeginTx")
		assert.Contains(t, m.String(), "AfterBeginTx")
		assert.Contains(t, m.String(), "BeforeRollback")
		assert.Contains(t, m.String(), "AfterRollback")
	})
}

func Test_conn_ExecContext(t *testing.T) {
	t.Run("mockConn", func(t *testing.T) {
		c, m := createMockConn()

		_, err := c.ExecContext(context.Background(), "any", nil)
		assert.NoError(t, err)
		s := m.String()
		assert.Contains(t, s, "BeforeExecContext")
		assert.Contains(t, s, "BeforePrepareContext")
		assert.Contains(t, s, "AfterPrepareContext")
		assert.Contains(t, s, "BeforeStmtExecContext")
		assert.Contains(t, s, "AfterStmtExecContext")
		assert.Contains(t, s, "AfterExecContext")
	})

	t.Run("mockConnExecer", func(t *testing.T) {
		m := &mockHook{}
		c := &conn{
			Conn:     &mockConnExecer{Conn: &mockConn{}},
			ConnHook: m,
		}
		_, err := c.ExecContext(context.Background(), "any", nil)
		assert.NoError(t, err)
		s := m.String()
		assert.Contains(t, s, "BeforeExecContext")
		assert.Contains(t, s, "AfterExecContext")
	})

	t.Run("mockConnExecerContext", func(t *testing.T) {
		m := &mockHook{}
		c := &conn{
			Conn:     &mockConnExecerContext{Conn: &mockConn{}},
			ConnHook: m,
		}
		_, err := c.ExecContext(context.Background(), "any", nil)
		assert.NoError(t, err)
		s := m.String()
		assert.Contains(t, s, "BeforeExecContext")
		assert.Contains(t, s, "AfterExecContext")
	})

}

func Test_conn_PrepareContext(t *testing.T) {
	c, m := createMockConn()

	st, err := c.PrepareContext(context.Background(), "any")
	assert.NoError(t, err)
	assert.NotNil(t, st)
	assert.Contains(t, m.String(), "BeforePrepareContext")
	assert.Contains(t, m.String(), "AfterPrepareContext")
}

func Test_conn_QueryContext(t *testing.T) {
	t.Run("mockConn", func(t *testing.T) {
		c, m := createMockConn()
		_, err := c.QueryContext(context.Background(), "any", nil)
		assert.NoError(t, err)
		assert.Contains(t, m.String(), "AfterQueryContext")
		assert.Contains(t, m.String(), "BeforeQueryContext")

	})

	t.Run("mockConnQueryer", func(t *testing.T) {
		m := &mockHook{}
		c := &conn{
			Conn:     &mockConnQueryer{Conn: &mockConn{}},
			ConnHook: m,
		}
		_, err := c.QueryContext(context.Background(), "any", nil)
		assert.NoError(t, err)
		s := m.String()
		assert.Contains(t, s, "BeforeQueryContext")
		assert.Contains(t, s, "AfterQueryContext")
	})

	t.Run("mockConnQueryerContext", func(t *testing.T) {
		m := &mockHook{}
		c := &conn{
			Conn:     &mockConnQueryerContext{Conn: &mockConn{}},
			ConnHook: m,
		}
		_, err := c.QueryContext(context.Background(), "any", nil)
		assert.NoError(t, err)
		s := m.String()
		assert.Contains(t, s, "BeforeQueryContext")
		assert.Contains(t, s, "AfterQueryContext")
	})
}

func TestMultiHook_Close(t *testing.T) {
	c, m := createMockConn()
	err := c.Close()
	assert.NoError(t, err)
	assert.Contains(t, m.String(), "AfterClose")
	assert.Contains(t, m.String(), "BeforeClose")
}

func Test_namedValueToValue(t *testing.T) {
	t.Run("no err", func(t *testing.T) {
		value, err := namedValueToValue([]driver.NamedValue{{
			Ordinal: 0,
			Value:   "A",
		}})
		assert.NoError(t, err)
		assert.EqualValues(t, value, []driver.Value{"A"})
	})

	t.Run("err", func(t *testing.T) {
		value, err := namedValueToValue([]driver.NamedValue{{
			Name:    "a",
			Ordinal: 0,
			Value:   "A",
		}})
		assert.Error(t, err)
		assert.Nil(t, value)
	})
}
