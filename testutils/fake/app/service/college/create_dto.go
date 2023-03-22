package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/go-faker/faker/v4"
)

func CreateCollegeInput() *service.CreateCollegeInput {
	input := new(service.CreateCollegeInput)
	faker.FakeData(input)

	return input
}
