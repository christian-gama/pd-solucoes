package fake

import (
	"fmt"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/go-faker/faker/v4"
)

func CourseEnrollment() *model.CourseEnrollment {
	courseEnrollment := new(model.CourseEnrollment)
	faker.FakeData(courseEnrollment)

	courseEnrollment.CourseSubject = CourseSubject()

	courseEnrollment.Student = Student()
	courseEnrollment.StudentID = courseEnrollment.Student.ID

	courseEnrollment.CourseSubject.Course = Course()
	courseEnrollment.CourseSubject.CourseID = courseEnrollment.CourseSubject.Course.ID

	courseEnrollment.CourseSubject.Subject = Subject()
	courseEnrollment.CourseSubject.SubjectID = courseEnrollment.CourseSubject.Subject.ID
	courseEnrollment.CourseSubjectID = courseEnrollment.CourseSubject.ID

	if err := courseEnrollment.Validate(); err != nil {
		panic(fmt.Errorf("error while generating fake courseEnrollment: %w", err))
	}

	return courseEnrollment
}
