package migrator

import (
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var (
	maxOpenConns    = 64
	maxIdleConns    = 32
	connMaxLifeTime = 8 * time.Hour
	logLevel        = core.LOG_DEBUG
)

// NewDatabase return xorm engine
// with some default params
func NewDatabase(dsn string) (*xorm.Engine, error) {
	db, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifeTime)
	db.SetLogLevel(logLevel)

	// DB stores UTC time, timezone of server is Asia/Shanghai
	// xorm use server's timezone by default, which would cause a bug, have to set
	// it to UTC explicitly
	db.DatabaseTZ, err = time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	db.TZLocation, err = time.LoadLocation("Local")
	if err != nil {
		return nil, err
	}

	return db, err
}
