package sqlplus

import (
	"database/sql/driver"
)

var (
	_ driver.Driver        = (*wrappedDriver)(nil)
	_ driver.DriverContext = (*wrappedDriverCtx)(nil)
)

type wrappedDriverCtx struct {
	driver.Driver
	Hook
}

type wrappedDriver struct {
	driver.Driver
	Hook
}

func New(d driver.Driver, hook Hook) driver.Driver {
	if _, ok := d.(driver.DriverContext); ok {
		return &wrappedDriverCtx{Driver: d, Hook: hook}
	}

	return &wrappedDriver{Driver: d, Hook: hook}
}

func (d *wrappedDriver) Open(name string) (driver.Conn, error) {
	c, err := d.Driver.Open(name)
	if err != nil {
		return nil, err
	}

	return &conn{Conn: c, ConnHook: d.Hook}, nil
}

// -----------------

func (d *wrappedDriverCtx) OpenConnector(name string) (driver.Connector, error) {
	if dd, ok := d.Driver.(driver.DriverContext); ok {
		openConnector, err := dd.OpenConnector(name)
		if err != nil {
			return nil, err
		}

		return &connector{Connector: openConnector, ConnectorHook: d.Hook}, nil
	}

	return nil, errNoInterfaceImplementation
}
