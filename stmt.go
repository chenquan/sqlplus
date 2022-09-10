package sqlplus

import (
	"context"
	"database/sql/driver"
)

var (
	_ driver.Stmt             = (*stmt)(nil)
	_ driver.StmtExecContext  = (*stmt)(nil)
	_ driver.StmtQueryContext = (*stmt)(nil)
)

type stmt struct {
	driver.Stmt
	query string
	StmtHook
}

func (s *stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (rows driver.Rows, err error) {
	ctx, err = s.BeforeStmtQueryContext(ctx, s.query, args)
	defer func() {
		_, err = s.AfterStmtQueryContext(ctx, s.query, args, err)
	}()
	if err != nil {
		return nil, err
	}

	switch ss := s.Stmt.(type) {
	case driver.StmtQueryContext:
		return ss.QueryContext(ctx, args)
	case interface {
		Query(args []driver.Value) (driver.Rows, error)
	}:
		value, err := namedValueToValue(args)
		if err != nil {
			return nil, err
		}

		return ss.Query(value)
	}

	return nil, errNoInterfaceImplementation
}

func (s *stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (r driver.Result, err error) {
	ctx, err = s.BeforeStmtExecContext(ctx, s.query, args)
	defer func() {
		_, err = s.AfterStmtExecContext(ctx, s.query, args, err)
	}()
	if err != nil {
		return nil, err
	}

	switch ss := s.Stmt.(type) {
	case driver.StmtExecContext:
		return ss.ExecContext(ctx, args)
	case interface {
		Exec(args []driver.Value) (driver.Result, error)
	}:
		value, err := namedValueToValue(args)
		if err != nil {
			return nil, err
		}

		return ss.Exec(value)
	}

	return nil, errNoInterfaceImplementation
}
