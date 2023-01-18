# sqlplus

[![Godoc](https://img.shields.io/badge/godoc-reference-brightgreen)](https://pkg.go.dev/github.com/chenquan/sqlplus)
[![Go Report Card](https://goreportcard.com/badge/github.com/chenquan/sqlplus)](https://goreportcard.com/report/github.com/chenquan/sqlplus)
[![Release](https://img.shields.io/github/v/release/chenquan/sqlplus.svg?style=flat-square)](https://github.com/chenquan/sqlplus)
[![codecov](https://codecov.io/gh/chenquan/sqlplus/branch/master/graph/badge.svg?token=74phc5KVI7)](https://codecov.io/gh/chenquan/sqlplus)
[![Download](https://goproxy.cn/stats/github.com/chenquan/sqlplus/badges/download-count.svg)](https://github.com/chenquan/sqlplus)
[![GitHub](https://img.shields.io/github/license/chenquan/sqlplus)](https://github.com/chenquan/sqlplus/blob/master/LICENSE)

A sql enhancement tool library based on `database/sql/driver`

> Unstable, use with caution in production environment.

# 😜installation

```shell
go get -u github.com/chenquan/sqlplus
```

# 👏how to use

Implement the `sqlplus.Hook` interface and wrap it with `sqlplus.New(d driver.Driver, hook Hook) driver.Driver`

# 👐ecosystem

- [sqltrace](https://github.com/chenquan/sqltrace): A low-code intrusion library that provides SQL tracing capabilities, suitable for any
  relational database (Sqlite3, MySQL, Oracle, SQL Server, PostgreSQL, TiDB, TDengine, etc.) and ORM libraries for various
  relational database (gorm, xorm, sqlx, etc.)
- [sqlbreaker](https://github.com/chenquan/sqlbreaker): A low-code intrusion library that provides SQL breaker capabilities, suitable for any
  relational database (Sqlite3, MySQL, Oracle, SQL Server, PostgreSQL, TiDB, TDengine, etc.) and ORM libraries for various
  relational database (gorm, xorm, sqlx, etc.)
