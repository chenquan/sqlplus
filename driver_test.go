package sqlplus

import (
	"database/sql/driver"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ driver.Driver = (*mockDriver)(nil)

type mockDriver struct {
}

func (d *mockDriver) Open(name string) (driver.Conn, error) {
	return &mockConn{}, nil
}

func TestNew(t *testing.T) {
	d := New(&mockDriver{}, &mockHook{})
	conn, err := d.Open("any")
	assert.NoError(t, err)
	assert.NotNil(t, conn)
}
