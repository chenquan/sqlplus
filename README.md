# sqlplus

A sql enhancement tool library based on `database/sql/driver`

# installation

```shell
go get -u github.com/chenquan/sqlplus
```

# usage

Implement the `sqlplus.Hook` interface and wrap it with `sqlplus.New(d driver.Driver, hook Hook) driver.Driver`

# example

- [sqltrace](https://github.com/chenquan/sqltrace): A sql tracing library,  suitable for any relational database such as Sqlite3, MySQL, Oracle, SQL Server, PostgreSQL, TiDB, etc.