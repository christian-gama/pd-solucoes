package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"
	"github.com/go-faker/faker/v4"
)

func UpdateSubjectInput() *service.UpdateSubjectInput {
	input := new(service.UpdateSubjectInput)
	faker.FakeData(input)

	return input
}
