package sqlplus

import (
	"context"
	"database/sql/driver"
)

type txContextKey struct{}
type tx struct {
	driver.Tx
	TxHook
	txContext context.Context
}

func TxContextFromContext(ctx context.Context) context.Context {
	value := ctx.Value(txContextKey{})
	if value != nil {
		return value.(context.Context)
	}

	return nil
}

func (t *tx) Commit() (err error) {
	ctx := context.WithValue(context.Background(), txContextKey{}, t.txContext)
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
	ctx := context.WithValue(context.Background(), txContextKey{}, t.txContext)
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
