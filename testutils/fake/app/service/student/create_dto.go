package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/go-faker/faker/v4"
)

func CreateStudentInput() *service.CreateInput {
	input := new(service.CreateInput)
	faker.FakeData(input)

	return input
}
