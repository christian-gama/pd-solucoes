package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/go-faker/faker/v4"
)

func FindAllStudentsInput() *service.FindAllStudentsInput {
	input := new(service.FindAllStudentsInput)
	faker.FakeData(input)

	return input
}
