package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/go-faker/faker/v4"
)

func FindOneStudentInput() *service.FindOneInput {
	input := new(service.FindOneInput)
	faker.FakeData(input)

	return input
}
