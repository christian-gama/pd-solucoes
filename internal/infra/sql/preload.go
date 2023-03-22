package sql

import (
	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

func preload(db *gorm.DB, name []string) *gorm.DB {
	for _, n := range name {
		db = db.Preload(strcase.ToCamel(n))
	}

	return db
}

func PreloadScope(name []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return preload(db, name)
	}
}
