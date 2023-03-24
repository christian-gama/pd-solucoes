package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	"github.com/go-faker/faker/v4"
)

func FindAllCourseEnrollmentsInput() *service.FindAllCourseEnrollmentsInput {
	input := new(service.FindAllCourseEnrollmentsInput)
	faker.FakeData(input)

	return input
}
