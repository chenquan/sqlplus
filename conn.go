package sqlhook

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

func (c *conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (dd driver.Result, err error) {
	ctx, err = c.BeforeExecContext(ctx, query, args)
	defer func() {
		_, err = c.AfterExecContext(ctx, query, args, err)
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

func (c *conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (dd driver.Rows, err error) {
	ctx, err = c.BeforeQueryContext(ctx, query, args)
	defer func() {
		_, err = c.AfterQueryContext(ctx, query, args, err)
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

func (c *conn) PrepareContext(ctx context.Context, query string) (dd driver.Stmt, err error) {
	ctx, err = c.BeforePrepareContext(ctx, query)
	defer func() {
		_, err = c.AfterPrepareContext(ctx, query, err)
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

	return &stmt{Stmt: st, StmtHook: c.ConnHook.(StmtHook), query: query}, nil
}

func (c *conn) BeginTx(ctx context.Context, opts driver.TxOptions) (dd driver.Tx, err error) {
	ctx, err = c.BeforeBeginTx(ctx, opts)
	defer func() {
		_, err = c.AfterBeginTx(ctx, opts, err)
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

	return &tx{Tx: t, TxHook: c.ConnHook.(TxHook)}, nil
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
