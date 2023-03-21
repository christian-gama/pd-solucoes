package faker

import (
	"reflect"

	"github.com/christian-gama/pd-solucoes/testutils/faker/providers"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
)

func InitializeProviders() {

	_ = faker.AddProvider("cpf", func(v reflect.Value) (any, error) {
		return providers.Cpf(), nil
	})

	_ = faker.AddProvider("cnpj", func(v reflect.Value) (any, error) {
		return providers.Cnpj(), nil
	})

	_ = faker.AddProvider("uint", func(v reflect.Value) (any, error) {
		return providers.Uint(), nil
	})

	_ = faker.AddProvider("[]uint", func(v reflect.Value) (any, error) {
		return providers.UintSlice(), nil
	})

	_ = faker.AddProvider("time_now", func(v reflect.Value) (any, error) {
		return providers.TimeNow(), nil
	})

	_ = faker.AddProvider("time_zero", func(v reflect.Value) (any, error) {
		return providers.TimeZero(), nil
	})

}

func Setup() {
	options.SetRandomMapAndSliceMinSize(2)
	options.SetRandomMapAndSliceMaxSize(5)
}
