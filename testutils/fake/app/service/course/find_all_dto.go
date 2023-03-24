package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
	"github.com/go-faker/faker/v4"
)

func FindAllCoursesInput() *service.FindAllInput {
	input := new(service.FindAllInput)
	faker.FakeData(input)

	return input
}
