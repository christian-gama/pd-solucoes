package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"
	"github.com/go-faker/faker/v4"
)

func DeleteSubjectInput() *service.DeleteSubjectInput {
	input := new(service.DeleteSubjectInput)
	faker.FakeData(input)

	return input
}
