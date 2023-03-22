package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/go-faker/faker/v4"
)

func CreateTeacherInput() *service.CreateTeacherInput {
	input := new(service.CreateTeacherInput)
	faker.FakeData(input)

	return input
}
