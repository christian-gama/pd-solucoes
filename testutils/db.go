package testutils

import (
	gosql "database/sql"
	"errors"
	"fmt"

	"github.com/christian-gama/pd-solucoes/internal/infra/sql"
	"gorm.io/gorm"
)

func Transaction(fn func(tx *gorm.DB)) {
	db := sql.MakePostgres()

	tx := func(tx *gorm.DB) error {
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				panic(fmt.Errorf("Rolling back transaction because of panic: %v", r))
			}
		}()

		fn(tx)

		return errors.New("it will rollback automatically on error")
	}

	db.Transaction(tx, &gosql.TxOptions{Isolation: gosql.LevelSerializable})
}
