package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
	"github.com/go-faker/faker/v4"
)

func CreateCourseInput() *service.CreateCourseInput {
	input := new(service.CreateCourseInput)
	faker.FakeData(input)

	return input
}
