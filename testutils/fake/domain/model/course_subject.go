package fake

import (
	"fmt"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/go-faker/faker/v4"
)

func CourseSubject() *model.CourseSubject {
	courseSubject := new(model.CourseSubject)
	faker.FakeData(courseSubject)

	courseSubject.Course = Course()
	courseSubject.CourseID = courseSubject.Course.ID

	courseSubject.Subject = Subject()
	courseSubject.SubjectID = courseSubject.Subject.ID

	if err := courseSubject.Validate(); err != nil {
		panic(fmt.Errorf("error while generating fake courseSubject: %w", err))
	}

	return courseSubject
}
