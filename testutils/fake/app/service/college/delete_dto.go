package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/go-faker/faker/v4"
)

func DeleteCollegeInput() *service.DeleteInput {
	input := new(service.DeleteInput)
	faker.FakeData(input)

	return input
}
