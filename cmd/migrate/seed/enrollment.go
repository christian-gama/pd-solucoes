package seed

import (
	"context"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
	"gorm.io/gorm"
)

func CourseEnrollmentData() []*service.CreateCourseEnrollmentInput {
	output := make([]*service.CreateCourseEnrollmentInput, 0)

	for i := range CourseSubjectData() {
		for j := range StudentData() {
			output = append(output, &service.CreateCourseEnrollmentInput{
				StudentID:       uint(j + 1),
				CourseSubjectID: uint(i + 1),
			})
		}
	}

	return output
}

func CourseEnrollment(ctx context.Context, db *gorm.DB) {
	repo := persistence.NewCourseEnrollment(db.Session(&gorm.Session{NewDB: true}))
	s := service.NewCreateCourseEnrollment(repo, service.NewFindOneCourseEnrollment(repo))

	for _, input := range CourseEnrollmentData() {
		if _, err := s.Handle(ctx, input); err != nil {
			panic(err)
		}
	}
}
