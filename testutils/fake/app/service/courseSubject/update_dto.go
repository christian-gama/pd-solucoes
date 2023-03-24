package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/go-faker/faker/v4"
)

func UpdateCourseSubjectInput() *service.UpdateCourseSubjectInput {
	input := new(service.UpdateCourseSubjectInput)
	faker.FakeData(input)

	return input
}
