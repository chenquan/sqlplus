package sqlplus

import (
	"context"
	"database/sql/driver"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ driver.Driver = (*mockDriver)(nil)
var _ driver.DriverContext = (*mockDriverCtx)(nil)
var _ driver.Driver = (*mockDriverCtx)(nil)

type (
	mockDriver    struct{}
	mockDriverCtx struct{}
)

func (d *mockDriver) Open(_ string) (driver.Conn, error) {
	return &mockConn{}, nil
}

// -----------------

func (m *mockDriverCtx) OpenConnector(_ string) (driver.Connector, error) {
	return &mockConnector{}, nil
}

func (m *mockDriverCtx) Open(_ string) (driver.Conn, error) {
	return &mockConn{}, nil
}

func TestNew(t *testing.T) {
	t.Run("mockDriver", func(t *testing.T) {
		d := New(&mockDriver{}, &mockHook{})
		conn, err := d.Open("any")
		assert.NoError(t, err)
		assert.NotNil(t, conn)
	})

	t.Run("mockDriverCtx", func(t *testing.T) {
		d := New(&mockDriverCtx{}, NewMultiHook(&mockHook{}))
		connect, err := d.Open("any")
		assert.NoError(t, err)
		assert.NotNil(t, connect)

		driverContext := d.(driver.DriverContext)
		openConnector, err := driverContext.OpenConnector("any")
		assert.NoError(t, err)
		connect, err = openConnector.Connect(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, connect)
		assert.NotNil(t, openConnector.Driver())
	})

}
