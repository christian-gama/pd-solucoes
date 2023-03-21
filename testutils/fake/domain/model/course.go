package fake

import (
	"fmt"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/go-faker/faker/v4"
)

func Course() *model.Course {
	course := new(model.Course)
	faker.FakeData(course)

	if err := course.Validate(); err != nil {
		panic(fmt.Errorf("error while generating fake course: %w", err))
	}

	return course
}
