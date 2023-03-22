package fake

import (
	"fmt"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/go-faker/faker/v4"
)

func College() *model.College {
	college := new(model.College)
	faker.FakeData(college)

	if err := college.Validate(); err != nil {
		panic(fmt.Errorf("error while generating fake college: %w", err))
	}

	return college
}
