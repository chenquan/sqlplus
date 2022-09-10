package sqlhook

import (
	"context"
	"database/sql/driver"
)

type tx struct {
	driver.Tx
	TxHook
}

func (t *tx) Commit() (err error) {
	ctx := context.Background()
	ctx, err = t.BeforeCommit(ctx)
	defer func() {
		_, err = t.AfterCommit(ctx, err)
	}()
	if err != nil {
		return
	}

	err = t.Tx.Commit()
	if err != nil {
		return
	}

	return
}

func (t *tx) Rollback() (err error) {
	ctx := context.Background()
	ctx, err = t.BeforeRollback(ctx)
	defer func() {
		_, err = t.AfterRollback(ctx, err)
	}()
	if err != nil {
		return
	}

	err = t.Tx.Rollback()
	if err != nil {
		return
	}

	return
}
