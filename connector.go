package sqlplus

import (
	"context"
	"database/sql/driver"
)

var _ driver.Connector = (*connector)(nil)

type connector struct {
	driver.Connector
	ConnectorHook
}

func (c *connector) Connect(ctx context.Context) (dc driver.Conn, err error) {
	ctx, err = c.BeforeConnect(ctx)
	defer func() {
		_, dc, err = c.AfterConnect(ctx, dc, err)
	}()
	if err != nil {
		return nil, err
	}

	cc, err := c.Connector.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &conn{Conn: cc, ConnHook: c.ConnectorHook.(ConnHook)}, nil
}

func (c *connector) Driver() driver.Driver {
	return &Driver{Driver: c.Connector.Driver(), Hook: c.ConnectorHook.(Hook)}
}
