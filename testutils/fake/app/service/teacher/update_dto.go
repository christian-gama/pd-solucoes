package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/go-faker/faker/v4"
)

func UpdateTeacherInput() *service.UpdateTeacherInput {
	input := new(service.UpdateTeacherInput)
	faker.FakeData(input)

	return input
}
