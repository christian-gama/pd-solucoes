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

type CourseEnrollmentDeps struct {
	CourseEnrollment *model.CourseEnrollment
	Student          *model.Student
	CourseSubject    *model.CourseSubject
}

func CreateCourseEnrollment(db *gorm.DB, deps *CourseEnrollmentDeps) *CourseEnrollmentDeps {
	if deps == nil {
		deps = &CourseEnrollmentDeps{}
	}

	student := deps.Student
	if student == nil {
		studentDeps := CreateStudent(db, nil)
		student = studentDeps.Student
		deps.Student = student
	}

	courseSubject := deps.CourseSubject
	if courseSubject == nil {
		courseSubjectDeps := CreateCourseSubject(db, nil)
		courseSubject = courseSubjectDeps.CourseSubject
		deps.CourseSubject = courseSubject
	}

	courseEnrollment := deps.CourseEnrollment
	if courseEnrollment == nil {
		courseEnrollment = fake.CourseEnrollment()
		courseEnrollment.StudentID = student.ID
		courseEnrollment.CourseSubjectID = courseSubject.ID
		courseEnrollment.Student = student
		courseEnrollment.CourseSubject = courseSubject

		courseEnrollment, err := persistence.NewCourseEnrollment(db).
			Create(context.Background(), repo.CreateCourseEnrollmentParams{
				CourseEnrollment: courseEnrollment,
			})
		if err != nil {
			panic(fmt.Errorf("could not create courseEnrollment: %w", err))
		}

		deps.CourseEnrollment = courseEnrollment
	}

	return deps
}
