package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
	"github.com/go-faker/faker/v4"
)

func UpdateCourseInput() *service.UpdateCourseInput {
	input := new(service.UpdateCourseInput)
	faker.FakeData(input)

	return input
}
