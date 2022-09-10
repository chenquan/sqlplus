package sqlhook

import (
	"database/sql/driver"
)

var (
	_ driver.Driver        = (*Driver)(nil)
	_ driver.DriverContext = (*DriverCtx)(nil)
)

type DriverCtx struct {
	driver.Driver
	Hook
}

type Driver struct {
	driver.Driver
	Hook
}

func New(d driver.Driver, hook Hook) driver.Driver {
	if _, ok := d.(driver.DriverContext); ok {
		return &DriverCtx{Driver: d, Hook: hook}
	}

	return &Driver{Driver: d, Hook: hook}
}

func (d Driver) Open(name string) (driver.Conn, error) {
	c, err := d.Driver.Open(name)
	if err != nil {
		return nil, err
	}

	return &conn{Conn: c, ConnHook: d.Hook}, nil
}

// -----------------

func (d DriverCtx) OpenConnector(name string) (driver.Connector, error) {
	if dd, ok := d.Driver.(driver.DriverContext); ok {
		openConnector, err := dd.OpenConnector(name)
		if err != nil {
			return nil, err
		}

		return &connector{Connector: openConnector, ConnectorHook: d.Hook}, nil
	}

	return nil, errNoInterfaceImplementation
}
