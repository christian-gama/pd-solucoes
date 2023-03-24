package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/go-faker/faker/v4"
)

func UpdateTeacherInput() *service.UpdateInput {
	input := new(service.UpdateInput)
	faker.FakeData(input)

	return input
}
