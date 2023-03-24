package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	"github.com/go-faker/faker/v4"
)

func FindOneCourseEnrollmentInput() *service.FindOneInput {
	input := new(service.FindOneInput)
	faker.FakeData(input)

	return input
}
