package migrator

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/xorm/migrate"
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

type migrator struct {
	path   string
	engine *xorm.Engine
}

func (m *migrator) migrate() error {
	migrations, err := m.findMigrations()
	if err != nil {
		return err
	}

	mg := migrate.New(m.engine, migrate.DefaultOptions, migrations)
	if err := mg.Migrate(); err != nil {
		return err
	}

	log.Println("migrate ok.")
	return nil
}

func (m *migrator) findMigrations() (ms []*migrate.Migration, err error) {
	fns, err := filepath.Glob(filepath.Join(m.path, "*.sql"))
	if err != nil {
		return
	}

	var id, sql string
	for _, fn := range fns {
		id = strings.TrimSuffix(filepath.Base(fn), filepath.Ext(fn))
		if buf, err := ioutil.ReadFile(fn); err == nil {
			sql = string(buf)
			ms = append(ms, migration(id, sql))
		}
	}
	return
}

func migration(id, sql string) *migrate.Migration {
	return &migrate.Migration{
		ID: id,
		Migrate: func(x *xorm.Engine) error {
			fmt.Println("migrate: ", id)
			_, err := x.Exec(sql)
			return err
		},
	}
}
