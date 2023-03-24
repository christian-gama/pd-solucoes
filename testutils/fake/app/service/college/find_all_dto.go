package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/go-faker/faker/v4"
)

func FindAllCollegesInput() *service.FindAllInput {
	input := new(service.FindAllInput)
	faker.FakeData(input)

	return input
}
