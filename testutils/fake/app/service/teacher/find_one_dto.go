package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/go-faker/faker/v4"
)

func FindOneTeacherInput() *service.FindOneInput {
	input := new(service.FindOneInput)
	faker.FakeData(input)

	return input
}
