package example

import (
	"os"
	"runtime"

	"github.com/kit-x/database/migrator"
)

func main() {
	dsn := os.Getenv("TEST_DB_DSN")
	if dsn == "" {
		panic("should set TEST_DB_DSN")
	}

	migrationPath := findPath()
	migrator.Migrate(dsn, migrationPath)
}

func findPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
}
