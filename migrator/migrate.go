package migrator

import (
	_ "github.com/go-sql-driver/mysql"
)

// Migrate ...
func Migrate(dsn, path string) {
	m, err := newMigrator(dsn, path)
	if err != nil {
		panic(err)
	}

	if err := m.migrate(); err != nil {
		panic(err)
	}
}

func newMigrator(dsn, path string) (*migrator, error) {
	engine, err := NewDatabase(dsn)
	if err != nil {
		return nil, err
	}

	return &migrator{engine: engine, path: path}, nil
}
