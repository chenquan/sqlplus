# sqlplus

[![Godoc](https://img.shields.io/badge/godoc-reference-brightgreen)](https://pkg.go.dev/github.com/chenquan/sqlplus)
[![Go Report Card](https://goreportcard.com/badge/github.com/chenquan/sqlplus)](https://goreportcard.com/report/github.com/chenquan/sqlplus)
[![Release](https://img.shields.io/github/v/release/chenquan/sqlplus.svg?style=flat-square)](https://github.com/chenquan/sqlplus)
[![codecov](https://codecov.io/gh/chenquan/sqlplus/branch/master/graph/badge.svg?token=74phc5KVI7)](https://codecov.io/gh/chenquan/sqlplus)
[![Download](https://goproxy.cn/stats/github.com/chenquan/sqlplus/badges/download-count.svg)](https://github.com/chenquan/sqlplus)
[![GitHub](https://img.shields.io/github/license/chenquan/sqlplus)](https://github.com/chenquan/sqlplus/blob/master/LICENSE)

A sql enhancement tool library based on `database/sql/driver`

# installation

```shell
go get -u github.com/chenquan/sqlplus
```

# usage

Implement the `sqlplus.Hook` interface and wrap it with `sqlplus.New(d driver.Driver, hook Hook) driver.Driver`

# example

- [sqltrace](https://github.com/chenquan/sqltrace): A sql tracing library, suitable for any relational database such as
  Sqlite3, MySQL, Oracle, SQL Server, PostgreSQL, TiDB, etc.