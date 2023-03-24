package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/go-faker/faker/v4"
)

func FindOneCourseSubjectInput() *service.FindOneCourseSubjectInput {
	input := new(service.FindOneCourseSubjectInput)
	faker.FakeData(input)

	return input
}
