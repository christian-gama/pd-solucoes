package fake

import (
	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
	"github.com/go-faker/faker/v4"
)

func FindOneCourseInput() *service.FindOneCourseInput {
	input := new(service.FindOneCourseInput)
	faker.FakeData(input)

	return input
}
