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

type (
	stmt struct {
		driver.Stmt
		query string
		StmtHook
		prepareContext context.Context
	}
	prepareContextKey struct{}
	stmtKey           struct{}
)

func PrepareContextFromContext(ctx context.Context) context.Context {
	value := ctx.Value(prepareContextKey{})
	if value != nil {
		return value.(context.Context)
	}

	return nil
}

func StmtFromContext(ctx context.Context) interface {
	driver.Stmt
	driver.StmtExecContext
	driver.StmtQueryContext
} {
	value := ctx.Value(stmtKey{})
	if value != nil {
		return nil
	}

	return value.(interface {
		driver.Stmt
		driver.StmtExecContext
		driver.StmtQueryContext
	})
}

// -----------------

func (s *stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (rows driver.Rows, err error) {
	query := s.query
	ctx = s.newStmtContext(ctx)
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
	query := s.query
	ctx = s.newStmtContext(ctx)
	ctx, args, err = s.BeforeStmtExecContext(ctx, query, args, nil)
	defer func() {
		_, r, err = s.AfterStmtExecContext(ctx, query, args, r, err)
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

func (s *stmt) newStmtContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, prepareContextKey{}, s.prepareContext)
	return context.WithValue(ctx, stmtKey{}, s)
}
