package fixture

import (
	"context"
	"fmt"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	"gorm.io/gorm"
)

type CourseDeps struct {
	Course  *model.Course
	College *model.College
}

func CreateCourse(db *gorm.DB, deps *CourseDeps) *CourseDeps {
	if deps == nil {
		deps = &CourseDeps{}
	}

	college := deps.College
	if college == nil {
		collegeDeps := CreateCollege(db, nil)
		college = collegeDeps.College
		deps.College = college
	}

	course := deps.Course
	if course == nil {
		course = fake.Course()
		course.CollegeID = college.ID
		course.College = college

		course, err := persistence.NewCourse(db).
			Create(context.Background(), repo.CreateCourseParams{
				Course: course,
			})
		if err != nil {
			panic(fmt.Errorf("could not create course: %w", err))
		}

		deps.Course = course
	}

	return deps
}
