package sqlhook

import (
	"database/sql/driver"
)

var _ driver.Tx = (*mockTx)(nil)

type mockTx struct {
}

func (t mockTx) Commit() error {
	return nil
}

func (t mockTx) Rollback() error {
	return nil
}
