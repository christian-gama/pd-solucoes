package seed

import (
	"context"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
	"gorm.io/gorm"
)

func SubjectData() []*service.CreateSubjectInput {
	return []*service.CreateSubjectInput{
		{
			Name:      "Algoritmos e Estrutura de Dados",
			TeacherID: 1,
		},
		{
			Name:      "Programação Orientada a Objetos",
			TeacherID: 1,
		},
		{
			Name:      "Banco de Dados",
			TeacherID: 2,
		},
		{
			Name:      "Sistemas Operacionais",
			TeacherID: 2,
		},
		{
			Name:      "Cálculo Diferencial e Integral",
			TeacherID: 3,
		},
		{
			Name:      "Estrutura de Dados",
			TeacherID: 4,
		},
		{
			Name:      "Análise e Projeto de Sistemas",
			TeacherID: 5,
		},
		{
			Name:      "Gestão de Projetos",
			TeacherID: 6,
		},
		{
			Name:      "Arquitetura de Computadores",
			TeacherID: 7,
		},
	}
}

func Subject(ctx context.Context, db *gorm.DB) {
	repo := persistence.NewSubject(db.Session(&gorm.Session{NewDB: true}))
	s := service.NewCreateSubject(repo)

	for _, input := range SubjectData() {
		if _, err := s.Handle(ctx, input); err != nil {
			panic(err)
		}
	}
}
