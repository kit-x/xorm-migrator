# xorm-migrator [![Build Status](https://travis-ci.org/kit-x/xorm-migrator.svg?branch=master)](https://travis-ci.org/kit-x/xorm-migrator) [![Go Report Card](https://goreportcard.com/badge/github.com/kit-x/xorm-migrator)](https://goreportcard.com/report/github.com/kit-x/xorm-migrator) [![GoDoc](https://godoc.org/github.com/kit-x/xorm-migrator?status.svg)](https://godoc.org/github.com/kit-x/xorm-migrator) [![codecov](https://codecov.io/gh/kit-x/xorm-migrator/branch/master/graph/badge.svg)](https://codecov.io/gh/kit-x/xorm-migrator)

基于 xorm 提供一些比较好用的数据库迁移功能

```
.
└── migrator 提供基本的迁移升级功能(暂时没有回滚操作)
```

## Usage

- [migrate](example/main.go)

## How to start dev

we should install those tools for dev

- Export `TEST_DB_DSN=root@tcp(mysql:3306)/xm_test?charset=utf8&parseTime=true&loc=Local`
