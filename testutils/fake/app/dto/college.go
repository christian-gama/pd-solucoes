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

func UpdateCollegeInput() *dto.UpdateCollegeInput {
	dto := new(dto.UpdateCollegeInput)
	faker.FakeData(dto)

	return dto
}

func FindOneCollegeInput() *dto.FindOneCollegeInput {
	dto := new(dto.FindOneCollegeInput)
	faker.FakeData(dto)

	return dto
}

func FindAllCollegesInput() *dto.FindAllCollegesInput {
	dto := new(dto.FindAllCollegesInput)
	faker.FakeData(dto)

	return dto
}
