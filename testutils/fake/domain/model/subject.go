package fake

import (
	"fmt"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/go-faker/faker/v4"
)

func Subject() *model.Subject {
	subject := new(model.Subject)
	faker.FakeData(subject)

	if err := subject.Validate(); err != nil {
		panic(fmt.Errorf("error while generating fake subject: %w", err))
	}

	return subject
}
