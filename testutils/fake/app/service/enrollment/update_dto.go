package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	"github.com/go-faker/faker/v4"
)

func UpdateCourseEnrollmentInput() *service.UpdateCourseEnrollmentInput {
	input := new(service.UpdateCourseEnrollmentInput)
	faker.FakeData(input)

	return input
}
