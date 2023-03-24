package seed

import (
	"context"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
	"gorm.io/gorm"
)

func CourseSubjectData() []*service.CreateCourseSubjectInput {
	output := make([]*service.CreateCourseSubjectInput, 0)

	for i := range CourseData() {
		for j := range SubjectData() {
			output = append(output, &service.CreateCourseSubjectInput{
				CourseID:  uint(i + 1),
				SubjectID: uint(j + 1),
			})
		}
	}

	return output
}

func CourseSubject(ctx context.Context, db *gorm.DB) {
	repo := persistence.NewCourseSubject(db.Session(&gorm.Session{NewDB: true}))
	s := service.NewCreateCourseSubject(repo, service.NewFindOneCourseSubject(repo))

	for _, input := range CourseSubjectData() {
		if _, err := s.Handle(ctx, input); err != nil {
			panic(err)
		}
	}
}
