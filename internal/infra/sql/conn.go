package sql

import (
	"time"

	"gorm.io/gorm"
)

// dialector is a function that returns a GORM dialector.
type dialector func(dsn string) gorm.Dialector

type conn struct {
	dialector
	opt *gorm.Config
}

// NewConn creates a new instance of a GORM connection.
func NewConn(dialector dialector, opt *gorm.Config) *conn {
	return &conn{dialector: dialector, opt: opt}
}

// Open will open a new GORM connection.
func (c *conn) Open() *gorm.DB {
	db, err := gorm.Open(c.dialector(Dsn()), c.opt)
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
