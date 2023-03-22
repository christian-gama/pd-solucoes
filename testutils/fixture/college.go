package fixture

import (
	"context"
	"fmt"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	"gorm.io/gorm"
)

type CollegeDeps struct {
	College *model.College
}

func CreateCollege(db *gorm.DB, deps *CollegeDeps) *CollegeDeps {
	if deps == nil {
		deps = &CollegeDeps{}
	}

	college := deps.College
	if college == nil {
		college = fake.College()

		college, err := persistence.NewCollege(db).
			Create(context.Background(), repo.CreateCollegeParams{
				College: college,
			})
		if err != nil {
			panic(fmt.Errorf("could not create college: %w", err))
		}

		deps.College = college
	}

	return deps
}
