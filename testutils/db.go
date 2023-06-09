package testutils

import (
	gosql "database/sql"
	"errors"

	"github.com/christian-gama/pd-solucoes/internal/infra/sql"
	"gorm.io/gorm"
)

func Transaction(
	failFn func(failureMessage string, msgAndArgs ...interface{}) bool,
	fn func(tx *gorm.DB),
) {
	db := sql.MakePostgres()

	tx := func(tx *gorm.DB) error {
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				failFn("test panicked", "panic: %v", r)
			}
		}()

		fn(tx)

		return errors.New("it will rollback automatically on error")
	}

	db.Transaction(tx, &gosql.TxOptions{Isolation: gosql.LevelSerializable})
}
