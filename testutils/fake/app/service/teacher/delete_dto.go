package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/go-faker/faker/v4"
)

func DeleteTeacherInput() *service.DeleteInput {
	input := new(service.DeleteInput)
	faker.FakeData(input)

	return input
}
