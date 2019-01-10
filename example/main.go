package main

import (
	"os"

	"github.com/kit-x/xorm-migrator/migrator"
)

func main() {
	dsn := os.Getenv("TEST_DB_DSN")
	if dsn == "" {
		panic("should set TEST_DB_DSN")
	}

	migrator.Migrate(dsn, "../migrator/fixtures")
}
