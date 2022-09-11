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
	query := s.query
	ctx, args, err = s.BeforeStmtQueryContext(ctx, query, args, nil)
	defer func() {
		_, rows, err = s.AfterStmtQueryContext(ctx, query, args, rows, err)
	}()
	if err != nil {
		return nil, err
	}

	switch ss := s.Stmt.(type) {
	case driver.StmtQueryContext:
		return ss.QueryContext(ctx, args)
	default:
		value, err := namedValueToValue(args)
		if err != nil {
			return nil, err
		}

		return s.Query(value)
	}

}

func (s *stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (r driver.Result, err error) {

	ctx, args, err = s.BeforeStmtExecContext(ctx, s.query, args, nil)
	defer func() {
		_, r, err = s.AfterStmtExecContext(ctx, s.query, args, r, err)
	}()
	if err != nil {
		return nil, err
	}

	switch ss := s.Stmt.(type) {
	case driver.StmtExecContext:
		return ss.ExecContext(ctx, args)
	default:
		value, err := namedValueToValue(args)
		if err != nil {
			return nil, err
		}

		return s.Exec(value)
	}
}
