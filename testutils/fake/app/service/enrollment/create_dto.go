package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	"github.com/go-faker/faker/v4"
)

func CreateCourseEnrollmentInput() *service.CreateCourseEnrollmentInput {
	input := new(service.CreateCourseEnrollmentInput)
	faker.FakeData(input)

	return input
}
