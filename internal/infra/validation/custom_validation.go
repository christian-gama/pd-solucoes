package validation

import (
	"errors"
	"reflect"
	"strings"

	"github.com/christian-gama/pd-solucoes/internal/infra/validation/validators"
	"github.com/go-playground/validator/v10"
)

func validateCnpj(fl validator.FieldLevel) bool {
	if fl.Field().Type().Kind() != reflect.String {
		return false
	}

	return validators.Cnpj(fl.Field().String())
}

func validateCPF(fl validator.FieldLevel) bool {
	if fl.Field().Type().Kind() != reflect.String {
		return false
	}

	return validators.Cpf(fl.Field().String())
}

func validateFilter(fl validator.FieldLevel) bool {
	if fl.Field().Type().Kind() != reflect.String {
		return false
	}

	params := strings.Split(fl.Param(), " ")
	if len(params) == 0 {
		panic(errors.New("the filter tag must have parameters"))
	}

	return validators.Filter(fl.Field().String(), params)
}

func validateSort(fl validator.FieldLevel) bool {
	if fl.Field().Type().Kind() != reflect.String {
		return false
	}

	params := strings.Split(fl.Param(), " ")
	if len(params) == 0 {
		panic(errors.New("the sort tag must have parameters"))
	}

	return validators.Sort(fl.Field().String(), params)
}
