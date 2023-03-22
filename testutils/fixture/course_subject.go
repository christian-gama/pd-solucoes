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

type CourseSubjectDeps struct {
	CourseSubject *model.CourseSubject
	Course        *model.Course
	Subject       *model.Subject
}

func CreateCourseSubject(db *gorm.DB, deps *CourseSubjectDeps) *CourseSubjectDeps {
	if deps == nil {
		deps = &CourseSubjectDeps{}
	}

	course := deps.Course
	if course == nil {
		courseDeps := CreateCourse(db, nil)
		course = courseDeps.Course
	}

	subject := deps.Subject
	if subject == nil {
		subjectDeps := CreateSubject(db, nil)
		subject = subjectDeps.Subject
	}

	courseSubject := deps.CourseSubject
	if courseSubject == nil {
		courseSubject = fake.CourseSubject()
		courseSubject.CourseID = course.ID
		courseSubject.SubjectID = subject.ID

		courseSubject, err := persistence.NewCourseSubject(db).
			Create(context.Background(), repo.CreateCourseSubjectParams{
				CourseSubject: courseSubject,
			})
		if err != nil {
			panic(fmt.Errorf("could not create courseSubject: %w", err))
		}

		deps.CourseSubject = courseSubject
	}

	return deps
}
