package migrator

import (
	"fmt"
	"os"
	"sort"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	x   *xorm.Engine
	dsn string
)

var teardown = func(t *testing.T) func(t *testing.T) {
	return func(t *testing.T) {
		if err := cleanDB(x); err != nil {
			t.Error(err)
		}
	}
}

func TestMain(m *testing.M) {
	dsn = os.Getenv("TEST_DB_DSN")
	if dsn == "" {
		panic("should set TEST_DB_DSN")
	}

	db, err := initDB(dsn)
	if err != nil {
		panic(fmt.Errorf("init test db with err: %v", err))
	}

	x = db
	m.Run()
}

func TestMigrate(t *testing.T) {
	defer teardown(t)

	tables, _ := x.DBMetas()
	if len(tables) != 0 {
		t.Error("test database should be empty")
		t.FailNow()
	}

	Migrate(dsn, "./fixtures")

	tables, err := x.DBMetas()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if len(tables) != 2 {
		t.Error("test database should contain only two table")
		t.FailNow()
	}

	tableNames := []string{}
	for i := range tables {
		tableNames = append(tableNames, tables[i].Name)
	}

	sort.Strings(tableNames)
	if tableNames[0] != "migrations" {
		t.Error("test database should contain migrations table")
	}
	if tableNames[1] != "users" {
		t.Error("test database should contain users table")
	}
}

func initDB(dsn string) (*xorm.Engine, error) {
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		return nil, err
	}

	db, err := xorm.NewEngine("mysql", dsn)
	if _, err := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", cfg.DBName)); err != nil {
		return nil, err
	}

	if _, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s;", cfg.DBName)); err != nil {
		return nil, err
	}

	if err := cleanDB(db); err != nil {
		return nil, err
	}

	return db, nil
}

func cleanDB(db *xorm.Engine) error {
	tables, _ := db.DBMetas()
	sess := db.NewSession()
	defer sess.Close()

	for _, table := range tables {
		if _, err := sess.Exec("set foreign_key_checks = 0"); err != nil {
			return fmt.Errorf("failed to disable foreign key checks")
		}
		if _, err := sess.Exec("drop table " + table.Name + " ;"); err != nil {
			return fmt.Errorf("failed to delete table: %v, err: %v", table.Name, err)
		}
		if _, err := sess.Exec("set foreign_key_checks = 1"); err != nil {
			return fmt.Errorf("failed to disable foreign key checks")
		}
	}

	return nil
}
