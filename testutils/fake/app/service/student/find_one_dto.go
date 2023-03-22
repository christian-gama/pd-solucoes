package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/go-faker/faker/v4"
)

func FindOneStudentInput() *service.FindOneStudentInput {
	input := new(service.FindOneStudentInput)
	faker.FakeData(input)

	return input
}
