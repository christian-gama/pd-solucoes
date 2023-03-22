package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/go-faker/faker/v4"
)

func FindOneCollegeInput() *service.FindOneCollegeInput {
	input := new(service.FindOneCollegeInput)
	faker.FakeData(input)

	return input
}
