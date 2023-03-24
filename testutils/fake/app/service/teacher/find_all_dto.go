package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/go-faker/faker/v4"
)

func FindAllTeachersInput() *service.FindAllInput {
	input := new(service.FindAllInput)
	faker.FakeData(input)

	return input
}
