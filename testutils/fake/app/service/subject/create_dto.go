package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"
	"github.com/go-faker/faker/v4"
)

func CreateSubjectInput() *service.CreateSubjectInput {
	input := new(service.CreateSubjectInput)
	faker.FakeData(input)

	return input
}
