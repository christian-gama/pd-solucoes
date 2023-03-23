package sql

import (
	"strings"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

func preload(db *gorm.DB, name []string) *gorm.DB {
	for _, n := range name {
		splittedByDot := strings.Split(n, ".")
		output := ""

		for i, s := range splittedByDot {
			if i != 0 {
				output += "."
			}

			output += strcase.ToCamel(s)
		}

		db = db.Preload(output)
	}

	return db
}

func PreloadScope(name []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return preload(db, name)
	}
}
