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

type StudentDeps struct {
	Student *model.Student
}

func CreateStudent(db *gorm.DB, deps *StudentDeps) *StudentDeps {
	if deps == nil {
		deps = &StudentDeps{}
	}

	student := deps.Student
	if student == nil {
		student = fake.Student()

		student, err := persistence.NewStudent(db).
			Create(context.Background(), repo.CreateStudentParams{
				Student: student,
			})
		if err != nil {
			panic(fmt.Errorf("could not create student: %w", err))
		}

		deps.Student = student
	}

	return deps
}
