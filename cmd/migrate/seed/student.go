package seed

import (
	"context"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
	"gorm.io/gorm"
)

func StudentData() []*service.CreateInput {
	return []*service.CreateInput{
		{
			Name: "Enzo Antonio Fogaça",
			Cpf:  "84093587434",
		},
		{
			Name: "Murilo Vinicius Lopes",
			Cpf:  "32077994185",
		},
		{
			Name: "Matheus Miguel Lopes",
			Cpf:  "80332383806",
		},
		{
			Name: "Laura Vera Brito",
			Cpf:  "66695194270",
		},
		{
			Name: "Pietra Isis Rebeca Ribeiro",
			Cpf:  "20439488907",
		},
		{
			Name: "Amanda Renata Francisca Sales",
			Cpf:  "16933473856",
		},
		{
			Name: "Sarah Yasmin Souza",
			Cpf:  "13416395093",
		},
		{
			Name: "Benedita Aparecida Ayla Galvão",
			Cpf:  "32687559334",
		},
		{
			Name: "Raquel Cláudia Clarice Rezende",
			Cpf:  "79175057581",
		},
		{
			Name: "Paulo Augusto Rocha",
			Cpf:  "92239767588",
		},
		{
			Name: "Melissa Marlene Vitória Pires",
			Cpf:  "35395992405",
		},
		{
			Name: "Bento Severino Rodrigo Cardoso",
			Cpf:  "55224030471",
		},
		{
			Name: "Regina Rita Débora Oliveira",
			Cpf:  "58632090380",
		},
		{
			Name: "Manoel Ruan Viana",
			Cpf:  "88802130086",
		},
		{
			Name: "Milena Sarah Fernandes",
			Cpf:  "44383362443",
		},
		{
			Name: "Isabel Joana da Conceição",
			Cpf:  "30302517219",
		},
		{
			Name: "Raimundo Giovanni da Cunha",
			Cpf:  "60077829646",
		},
		{
			Name: "Renan Márcio Gabriel da Luz",
			Cpf:  "76418799771",
		},
		{
			Name: "Eduarda Rita Nair Drumond",
			Cpf:  "31872211658",
		},
		{
			Name: "Simone Caroline Moraes",
			Cpf:  "21767625332",
		},
		{
			Name: "Mariana Raquel Teixeira",
			Cpf:  "75428049871",
		},
		{
			Name: "Cauê Eduardo Eduardo Silva",
			Cpf:  "86823045116",
		},
		{
			Name: "Isabela Lúcia Fátima das Neves",
			Cpf:  "84081888280",
		},
		{
			Name: "Mariane Emanuelly Giovanna Duarte",
			Cpf:  "74875180004",
		},
		{
			Name: "Lavínia Analu Aparício",
			Cpf:  "46509532948",
		},
		{
			Name: "Francisco César Gomes",
			Cpf:  "18569968248",
		},
	}
}

func Student(ctx context.Context, db *gorm.DB) {
	repo := persistence.NewStudent(db.Session(&gorm.Session{NewDB: true}))
	s := service.NewCreateStudent(repo)

	for _, input := range StudentData() {
		if _, err := s.Handle(ctx, input); err != nil {
			panic(err)
		}
	}
}
