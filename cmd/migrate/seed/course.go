package seed

import (
	"context"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
	"gorm.io/gorm"
)

func CourseData() []*service.CreateCourseInput {
	return []*service.CreateCourseInput{
		{
			Name:      "Engenharia de Software",
			CollegeID: 1,
		},
		{
			Name:      "Ciência da Computação",
			CollegeID: 1,
		},
	}
}

func Course(ctx context.Context, db *gorm.DB) {
	repo := persistence.NewCourse(db.Session(&gorm.Session{NewDB: true}))
	s := service.NewCreateCourse(repo)

	for _, input := range CourseData() {
		if _, err := s.Handle(ctx, input); err != nil {
			panic(err)
		}
	}
}
