package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/go-faker/faker/v4"
)

func UpdateStudentInput() *service.UpdateInput {
	input := new(service.UpdateInput)
	faker.FakeData(input)

	return input
}
