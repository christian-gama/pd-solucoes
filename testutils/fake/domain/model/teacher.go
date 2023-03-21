package fake

import (
	"fmt"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/go-faker/faker/v4"
)

func Teacher() *model.Teacher {
	teacher := new(model.Teacher)
	faker.FakeData(teacher)

	if err := teacher.Validate(); err != nil {
		panic(fmt.Errorf("error while generating fake teacher: %w", err))
	}

	return teacher
}
