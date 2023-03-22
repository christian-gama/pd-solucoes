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

type TeacherDeps struct {
	Teacher *model.Teacher
}

func CreateTeacher(db *gorm.DB, deps *TeacherDeps) *TeacherDeps {
	if deps == nil {
		deps = &TeacherDeps{}
	}

	teacher := deps.Teacher
	if teacher == nil {
		teacher = fake.Teacher()

		teacher, err := persistence.NewTeacher(db).
			Create(context.Background(), repo.CreateTeacherParams{
				Teacher: teacher,
			})
		if err != nil {
			panic(fmt.Errorf("could not create teacher: %w", err))
		}

		deps.Teacher = teacher
	}

	return deps
}
