package fake

import (
	"fmt"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/go-faker/faker/v4"
)

func Student() *model.Student {
	student := new(model.Student)
	faker.FakeData(student)

	if err := student.Validate(); err != nil {
		panic(fmt.Errorf("error while generating fake student: %w", err))
	}

	return student
}
