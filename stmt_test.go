package sqlhook

import (
	"database/sql/driver"
)

type mockStmt struct {
}

func (s *mockStmt) Close() error {
	return nil
}

func (s *mockStmt) NumInput() int {
	return 0
}

func (s *mockStmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, nil
}

func (s *mockStmt) Query(args []driver.Value) (driver.Rows, error) {
	return nil, nil
}
