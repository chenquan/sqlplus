package sqlplus

import (
	"context"
	"database/sql/driver"
)

var _ driver.Connector = (*mockConnector)(nil)

type mockConnector struct{}

func (m *mockConnector) Connect(_ context.Context) (driver.Conn, error) {
	return &mockConn{}, nil
}

func (m *mockConnector) Driver() driver.Driver {
	return &mockDriver{}
}
