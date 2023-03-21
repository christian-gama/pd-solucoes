package fake

import (
	"fmt"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/go-faker/faker/v4"
)

func College() *model.College {
	college := new(model.College)
	faker.FakeData(college)

	college.Courses = make([]*model.Course, 3)
	for i := range college.Courses {
		college.Courses[i] = Course()
	}

	if err := college.Validate(); err != nil {
		panic(fmt.Errorf("error while generating fake college: %w", err))
	}

	return college
}
