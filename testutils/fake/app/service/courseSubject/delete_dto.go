package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/go-faker/faker/v4"
)

func DeleteCourseSubjectInput() *service.DeleleInput {
	input := new(service.DeleleInput)
	faker.FakeData(input)

	return input
}
