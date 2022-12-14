package sqlplus

import (
	"context"
	"database/sql/driver"
)

type (
	tx struct {
		driver.Tx
		TxHook
		txContext context.Context
	}
)

func (t *tx) Commit() (err error) {
	ctx := t.txContext
	ctx, err = t.BeforeCommit(ctx, nil)
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
	ctx := t.txContext
	ctx, err = t.BeforeRollback(ctx, nil)
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
