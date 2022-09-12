# sqlplus

A sql enhancement tool library based on `database/sql/driver`

# installation

```shell
go get -u github.com/chenquan/sqlplus
```

# usage
Implement the `sqlplus.Hook` interface and wrap it with `sqlplus.New(d driver.Driver, hook Hook) driver.Driver`

# example

- [https://github.com/chenquan/sqltrace](sqltrace): A sql link tracking library, 
suitable for any relational database such as MySQL, oracle, SQL Server, PostgreSQL,TiDB etc.