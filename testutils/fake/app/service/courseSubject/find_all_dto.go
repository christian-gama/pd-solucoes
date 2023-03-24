package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/go-faker/faker/v4"
)

func FindAllCourseSubjectsInput() *service.FindAllCourseSubjectsInput {
	input := new(service.FindAllCourseSubjectsInput)
	faker.FakeData(input)

	return input
}
