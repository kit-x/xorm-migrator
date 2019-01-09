# xorm-migrator

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
