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
	courseEnrollment.CourseSubject.Course = Course()
	courseEnrollment.CourseSubject.Course.College = College()
	courseEnrollment.CourseSubject.Subject.Teacher = Teacher()
	courseEnrollment.CourseSubject.Subject = Subject()

	courseEnrollment.CourseSubject.Course.CollegeID = courseEnrollment.CourseSubject.Course.College.ID
	courseEnrollment.CourseSubject.CourseID = courseEnrollment.CourseSubject.Course.ID
	courseEnrollment.CourseSubject.SubjectID = courseEnrollment.CourseSubject.Subject.ID
	courseEnrollment.CourseSubject.Subject.TeacherID = courseEnrollment.CourseSubject.Subject.Teacher.ID
	courseEnrollment.StudentID = courseEnrollment.Student.ID
	courseEnrollment.CourseSubjectID = courseEnrollment.CourseSubject.ID

	if err := courseEnrollment.Validate(); err != nil {
		panic(fmt.Errorf("error while generating fake courseEnrollment: %w", err))
	}

	return courseEnrollment
}
