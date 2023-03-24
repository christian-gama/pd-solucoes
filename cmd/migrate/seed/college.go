package seed

import (
	"context"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
	"gorm.io/gorm"
)

func CollegeData() []*service.CreateInput {
	return []*service.CreateInput{
		{
			Name: "Universidade Federal de Minas Gerais",
			Cnpj: "05124078523150",
		},
	}
}

func College(ctx context.Context, db *gorm.DB) {
	repo := persistence.NewCollege(db.Session(&gorm.Session{NewDB: true}))
	s := service.NewCreateCollege(repo)

	for _, input := range CollegeData() {
		if _, err := s.Handle(ctx, input); err != nil {
			panic(err)
		}
	}
}
