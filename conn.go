package sqlplus

import (
	"context"
	"database/sql/driver"
	"errors"
)

var (
	_ driver.ConnPrepareContext = (*conn)(nil)
	_ driver.QueryerContext     = (*conn)(nil)
	_ driver.ConnBeginTx        = (*conn)(nil)
	_ driver.ExecerContext      = (*conn)(nil)
)

type conn struct {
	driver.Conn
	ConnHook
}

func (c *conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (result driver.Result, err error) {
	ctx, query, args, err = c.BeforeExecContext(ctx, query, args, nil)
	defer func() {
		_, result, err = c.AfterExecContext(ctx, query, args, result, err)
	}()
	if err != nil {
		return nil, err
	}

	switch cc := c.Conn.(type) {
	case driver.ExecerContext:
		return cc.ExecContext(ctx, query, args)
	case driver.Execer:
		value, err := namedValueToValue(args)
		if err != nil {
			return nil, err
		}

		return cc.Exec(query, value)
	default:
		ts, err := c.PrepareContext(ctx, query)
		if err != nil {
			return nil, err
		}

		return ts.(*stmt).ExecContext(ctx, args)
	}
}

func (c *conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (rows driver.Rows, err error) {
	ctx, query, args, err = c.BeforeQueryContext(ctx, query, args, nil)
	defer func() {
		_, rows, err = c.AfterQueryContext(ctx, query, args, rows, err)
	}()
	if err != nil {
		return
	}

	switch cc := c.Conn.(type) {
	case driver.QueryerContext:
		return cc.QueryContext(ctx, query, args)
	case driver.Queryer:
		value, err := namedValueToValue(args)
		if err != nil {
			return nil, err
		}

		return cc.Query(query, value)
	default:
		var st driver.Stmt
		st, err = c.PrepareContext(ctx, query)
		if err != nil {
			return nil, err
		}

		return st.(*stmt).QueryContext(ctx, args)
	}
}

func (c *conn) PrepareContext(ctx context.Context, query string) (s driver.Stmt, err error) {
	ctx, query, err = c.BeforePrepareContext(ctx, query, nil)
	defer func() {
		_, s, err = c.AfterPrepareContext(ctx, query, s, err)
	}()
	if err != nil {
		return nil, err
	}

	var st driver.Stmt
	if cc, ok := c.Conn.(driver.ConnPrepareContext); ok {
		st, err = cc.PrepareContext(ctx, query)
	} else {
		st, err = c.Conn.Prepare(query)
	}

	if err != nil {
		return st, err
	}

	return &stmt{Stmt: st, StmtHook: c.ConnHook.(StmtHook), query: query, prepareContext: ctx}, nil
}

func (c *conn) BeginTx(ctx context.Context, opts driver.TxOptions) (dd driver.Tx, err error) {
	ctx, opts, err = c.BeforeBeginTx(ctx, opts, nil)
	defer func() {
		_, dd, err = c.AfterBeginTx(ctx, opts, dd, err)
	}()
	if err != nil {
		return nil, err
	}

	var t driver.Tx
	switch cc := c.Conn.(type) {
	case driver.ConnBeginTx:
		t, err = cc.BeginTx(ctx, opts)
	default:
		t, err = c.Begin()
	}
	if err != nil {
		return nil, err
	}

	return &tx{Tx: t, TxHook: c.ConnHook.(TxHook), txContext: ctx}, nil
}

func namedValueToValue(named []driver.NamedValue) ([]driver.Value, error) {
	dargs := make([]driver.Value, len(named))
	for n, param := range named {
		if len(param.Name) > 0 {
			return nil, errors.New("sql: driver does not support the use of Named Parameters")
		}

		dargs[n] = param.Value
	}

	return dargs, nil
}
