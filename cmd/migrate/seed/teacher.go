package seed

import (
	"context"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
	"gorm.io/gorm"
)

func TeacherData() []*service.CreateTeacherInput {
	return []*service.CreateTeacherInput{
		{
			Name:   "Luana Alice Bianca Ramos",
			Degree: "Engenharia de Software",
		},
		{
			Name:   "Olivia Lara Daiane da Paz",
			Degree: "Matemática",
		},
		{
			Name:   "Oliver Gabriel Breno da Costa",
			Degree: "Física",
		},
		{
			Name:   "Enzo Miguel da Rosa",
			Degree: "Ciências da Computação",
		},
		{
			Name:   "Fabiana Isabelly Débora Nogueira",
			Degree: "Análise e Desenvolvimento de Sistemas",
		},
		{
			Name:   "Victor Emanuel Dias",
			Degree: "Engenharia da Computação",
		},
		{
			Name:   "Sabrina Carla Corte Real",
			Degree: "Sistemas de Informação",
		},
	}
}

func Teacher(ctx context.Context, db *gorm.DB) {
	repo := persistence.NewTeacher(db.Session(&gorm.Session{NewDB: true}))
	s := service.NewCreateTeacher(repo)

	for _, input := range TeacherData() {
		if _, err := s.Handle(ctx, input); err != nil {
			panic(err)
		}
	}
}
