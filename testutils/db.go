package testutils

import (
	gosql "database/sql"
	"errors"
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/infra/sql"
	"gorm.io/gorm"
)

func Transaction(t *testing.T, fn func(tx *gorm.DB)) {
	db := sql.MakePostgres()

	tx := func(tx *gorm.DB) error {
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				t.Fatalf("test panicked: %v", r)
			}
		}()

		fn(tx)

		return errors.New("it will rollback automatically on error")
	}

	db.Transaction(tx, &gosql.TxOptions{Isolation: gosql.LevelSerializable})
}
