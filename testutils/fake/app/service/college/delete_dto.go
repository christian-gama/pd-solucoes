package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/go-faker/faker/v4"
)

func DeleteCollegeInput() *service.DeleteCollegeInput {
	input := new(service.DeleteCollegeInput)
	faker.FakeData(input)

	return input
}
