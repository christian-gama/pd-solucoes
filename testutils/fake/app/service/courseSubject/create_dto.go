package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/go-faker/faker/v4"
)

func CreateCourseSubjectInput() *service.CreateCourseSubjectInput {
	input := new(service.CreateCourseSubjectInput)
	faker.FakeData(input)

	return input
}
