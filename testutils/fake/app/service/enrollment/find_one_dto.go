package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	"github.com/go-faker/faker/v4"
)

func FindOneCourseEnrollmentInput() *service.FindOneCourseEnrollmentInput {
	input := new(service.FindOneCourseEnrollmentInput)
	faker.FakeData(input)

	return input
}
