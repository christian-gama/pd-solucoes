package persistence_test

import (
	"context"
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	queryingPort "github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
	"github.com/christian-gama/pd-solucoes/internal/infra/querying"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	"github.com/christian-gama/pd-solucoes/testutils/fixture"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"gorm.io/gorm"
)

type CourseSuite struct {
	suite.SuiteWithConn
	Course func(db *gorm.DB) repo.Course
}

func TestCourseSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(CourseSuite))
}

func (s *CourseSuite) SetupTest() {
	s.Course = func(db *gorm.DB) repo.Course {
		return persistence.NewCourse(db)
	}
}

func (s *CourseSuite) TestCreate() {
	type Sut struct {
		Sut    func(ctx context.Context, params repo.CreateCourseParams) (*model.Course, error)
		Ctx    context.Context
		Params repo.CreateCourseParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		course := fake.Course()
		params := repo.CreateCourseParams{
			Course: course,
		}

		sut := s.Course(db).Create

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should create a new course", func(db *gorm.DB) {
		sut := makeSut(db)

		collegeDeps := fixture.CreateCollege(db, nil)

		sut.Params.Course.CollegeID = collegeDeps.College.ID

		model, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.NotZero(model.ID, "should have an ID")
	})

	s.Run("should return an error when the course already exists", func(db *gorm.DB) {
		sut := makeSut(db)

		collegeDeps := fixture.CreateCollege(db, nil)

		sut.Params.Course.CollegeID = collegeDeps.College.ID

		_, err := sut.Sut(sut.Ctx, sut.Params)
		s.NoError(err)

		_, err = sut.Sut(sut.Ctx, sut.Params)
		s.Error(err)
	})
}

func (s *CourseSuite) TestDelete() {
	type Sut struct {
		Sut    func(ctx context.Context, params repo.DeleteCourseParams) error
		Ctx    context.Context
		Params repo.DeleteCourseParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.DeleteCourseParams{}
		sut := s.Course(db).Delete

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should delete a course", func(db *gorm.DB) {
		sut := makeSut(db)

		courseDeps := fixture.CreateCourse(db, nil)

		sut.Params.ID = courseDeps.Course.ID

		err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
	})

	s.Run("should delete nothing if the course does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Params.ID = 404_404_404

		err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
	})
}

func (s *CourseSuite) TestFindOne() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.FindOneCourseParams,
			preload ...string,
		) (*model.Course, error)
		Ctx    context.Context
		Params repo.FindOneCourseParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.FindOneCourseParams{
			Filterer: querying.Filter{},
			ID:       0,
		}
		sut := s.Course(db).FindOne

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should find a course", func(db *gorm.DB) {
		sut := makeSut(db)

		courseDeps := fixture.CreateCourse(db, nil)

		sut.Params.ID = courseDeps.Course.ID

		course, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(course.ID, courseDeps.Course.ID, "should have the same ID")
	})

	s.Run("should return an error if the course does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Params.ID = 404_404_404

		_, err := sut.Sut(sut.Ctx, sut.Params)

		s.Error(err)
	})

	s.Run("should return a course with filtered fields", func(db *gorm.DB) {
		sut := makeSut(db)

		courseDeps := fixture.CreateCourse(db, nil)

		sut.Params.ID = courseDeps.Course.ID
		sut.Params.Filterer = querying.AddFilter("name", querying.EqOperator, courseDeps.Course.Name)

		course, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(course.ID, courseDeps.Course.ID, "should have the same ID")
		s.Equal(course.Name, courseDeps.Course.Name, "should have the same name")
	})

	s.Run("should not return a course if filter does not met the condition", func(db *gorm.DB) {
		sut := makeSut(db)

		courseDeps := fixture.CreateCourse(db, nil)

		sut.Params.ID = courseDeps.Course.ID
		sut.Params.Filterer = querying.AddFilter("name", querying.EqOperator, "wrong name")

		course, err := sut.Sut(sut.Ctx, sut.Params)

		s.Error(err)
		s.Nil(course, "should not return a course")
	})
}

func (s *CourseSuite) TestFindAll() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.FindAllCourseParams,
			preload ...string,
		) (*queryingPort.PaginationOutput[*model.Course], error)
		Ctx    context.Context
		Params repo.FindAllCourseParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.FindAllCourseParams{
			Paginator: &querying.Pagination{},
			Sorter:    querying.Sort{},
			Filterer:  querying.Filter{},
		}
		sut := s.Course(db).FindAll

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should find courses", func(db *gorm.DB) {
		sut := makeSut(db)

		length := 3
		for i := 0; i < length; i++ {
			fixture.CreateCourse(db, nil)
		}

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.NotZero(result.Results[0].ID, "should have a valid id")
		s.Equal(length, result.Total, "should return %d total", length)
		s.Len(result.Results, length, "should return %d results", length)
	})

	s.Run("should return the correct courses using filter", func(db *gorm.DB) {
		sut := makeSut(db)

		courseDeps := fixture.CreateCourse(db, nil)
		length := 3
		for i := 0; i < length; i++ {
			fixture.CreateCourse(db, nil)
		}

		sut.Params.Filterer = sut.Params.Filterer.Add(
			"name",
			querying.EqOperator,
			courseDeps.Course.Name,
		)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(result.Results[0].ID, courseDeps.Course.ID, "should have the same ID")
		s.Equal(1, result.Total, "should return only one course")
		s.Len(result.Results, 1, "should return only one course")
	})

	s.Run("should return the correct courses using sorter as desc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.CreateCourse(db, nil)
		}

		sut.Params.Sorter = sut.Params.Sorter.Add("id", true)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Greater(int(result.Results[1].ID), int(result.Results[2].ID), "should have the correct order")
	})

	s.Run("should return the correct courses using sorter as asc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.CreateCourse(db, nil)
		}

		sut.Params.Sorter = sut.Params.Sorter.Add("id", false)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Greater(int(result.Results[2].ID), int(result.Results[1].ID), "should have the correct order")
	})

	s.Run("should return the correct courses using pagination", func(db *gorm.DB) {
		sut := makeSut(db)

		courses := make([]*model.Course, 0)
		for i := 0; i < 3; i++ {
			courseDeps := fixture.CreateCourse(db, nil)
			courses = append(courses, courseDeps.Course)
		}

		sut.Params.Paginator = sut.Params.Paginator.SetLimit(1).SetPage(1)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(3, result.Total, "should return the correct total")
		s.Len(result.Results, 1, "should return the correct number of courses")
		s.Equal(int(courses[0].ID), int(result.Results[0].ID), "should return the correct course")
	})
}

func (s *CourseSuite) TestUpdate() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.UpdateCourseParams,
		) (*model.Course, error)
		Ctx    context.Context
		Params repo.UpdateCourseParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.UpdateCourseParams{
			Course: fake.Course(),
		}
		sut := s.Course(db).Update

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should update a course", func(db *gorm.DB) {
		sut := makeSut(db)

		courseDeps := fixture.CreateCourse(db, nil)

		*sut.Params.Course = *courseDeps.Course
		sut.Params.Course.Name = "new name"

		newCourse, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(courseDeps.Course.ID, newCourse.ID, "should return the correct ID")
		s.NotEqual(courseDeps.Course.Name, newCourse.Name, "should return the correct name")
	})

	s.Run("should return an error if tries to update a non existent course", func(db *gorm.DB) {
		sut := makeSut(db)

		fixture.CreateCourse(db, nil)

		sut.Params.Course.Name = "new name"
		sut.Params.Course.ID = 404_404_404

		newCourse, err := sut.Sut(sut.Ctx, sut.Params)

		s.Error(err)
		s.Nil(newCourse, "should return nil")
	})
}
