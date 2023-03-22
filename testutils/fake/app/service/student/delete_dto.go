package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/go-faker/faker/v4"
)

func DeleteStudentInput() *service.DeleteStudentInput {
	input := new(service.DeleteStudentInput)
	faker.FakeData(input)

	return input
}
