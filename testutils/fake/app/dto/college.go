package fake

import (
	"github.com/christian-gama/pd-solucoes/internal/app/dto"
	"github.com/go-faker/faker/v4"
)

func CreateCollegeInput() *dto.CreateCollegeInput {
	dto := new(dto.CreateCollegeInput)
	faker.FakeData(dto)

	return dto
}
