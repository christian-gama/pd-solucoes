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

type SubjectDeps struct {
	Subject *model.Subject
	Teacher *model.Teacher
}

func CreateSubject(db *gorm.DB, deps *SubjectDeps) *SubjectDeps {
	if deps == nil {
		deps = &SubjectDeps{}
	}

	teacher := deps.Teacher
	if teacher == nil {
		teacherDeps := CreateTeacher(db, nil)
		teacher = teacherDeps.Teacher
	}

	subject := deps.Subject
	if subject == nil {
		subject = fake.Subject()
		subject.TeacherID = teacher.ID

		subject, err := persistence.NewSubject(db).
			Create(context.Background(), repo.CreateSubjectParams{
				Subject: subject,
			})
		if err != nil {
			panic(fmt.Errorf("could not create subject: %w", err))
		}

		deps.Subject = subject
	}

	return deps
}
