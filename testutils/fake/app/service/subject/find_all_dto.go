package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"
	"github.com/go-faker/faker/v4"
)

func FindAllSubjectsInput() *service.FindAllSubjectsInput {
	input := new(service.FindAllSubjectsInput)
	faker.FakeData(input)

	return input
}
