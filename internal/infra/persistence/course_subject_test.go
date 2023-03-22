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

type CourseSubjectSuite struct {
	suite.SuiteWithConn
	CourseSubject func(db *gorm.DB) repo.CourseSubject
}

func TestCourseSubjectSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(CourseSubjectSuite))
}

func (s *CourseSubjectSuite) SetupTest() {
	s.CourseSubject = func(db *gorm.DB) repo.CourseSubject {
		return persistence.NewCourseSubject(db)
	}
}

func (s *CourseSubjectSuite) TestCreate() {
	type Sut struct {
		Sut    func(context.Context, repo.CreateCourseSubjectParams) (*model.CourseSubject, error)
		Ctx    context.Context
		Params repo.CreateCourseSubjectParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		courseSubject := fake.CourseSubject()
		params := repo.CreateCourseSubjectParams{
			CourseSubject: courseSubject,
		}

		sut := s.CourseSubject(db).Create

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should create a new courseSubject", func(db *gorm.DB) {
		sut := makeSut(db)

		courseDeps := fixture.CreateCourse(db, nil)
		subjectDeps := fixture.CreateSubject(db, nil)

		sut.Params.CourseSubject.CourseID = courseDeps.Course.ID
		sut.Params.CourseSubject.SubjectID = subjectDeps.Subject.ID

		model, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.NotZero(model.ID, "should have an ID")
	})

	s.Run("should return an error when the courseSubject already exists", func(db *gorm.DB) {
		sut := makeSut(db)

		courseDeps := fixture.CreateCourse(db, nil)
		subjectDeps := fixture.CreateSubject(db, nil)

		sut.Params.CourseSubject.CourseID = courseDeps.Course.ID
		sut.Params.CourseSubject.SubjectID = subjectDeps.Subject.ID

		_, err := sut.Sut(sut.Ctx, sut.Params)
		s.NoError(err)

		_, err = sut.Sut(sut.Ctx, sut.Params)
		s.Error(err)
	})
}

func (s *CourseSubjectSuite) TestDelete() {
	type Sut struct {
		Sut    func(ctx context.Context, params repo.DeleteCourseSubjectParams) error
		Ctx    context.Context
		Params repo.DeleteCourseSubjectParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.DeleteCourseSubjectParams{}
		sut := s.CourseSubject(db).Delete

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should delete a courseSubject", func(db *gorm.DB) {
		sut := makeSut(db)

		courseSubjectDeps := fixture.CreateCourseSubject(db, nil)

		sut.Params.ID = courseSubjectDeps.CourseSubject.ID

		err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
	})

	s.Run("should delete nothing if the courseSubject does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Params.ID = 404_404_404

		err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
	})
}

func (s *CourseSubjectSuite) TestFindOne() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.FindOneCourseSubjectParams,
			preload ...string,
		) (*model.CourseSubject, error)
		Ctx    context.Context
		Params repo.FindOneCourseSubjectParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.FindOneCourseSubjectParams{
			Filterer: querying.Filter{},
			ID:       0,
		}
		sut := s.CourseSubject(db).FindOne

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should find a courseSubject", func(db *gorm.DB) {
		sut := makeSut(db)

		courseSubjectDeps := fixture.CreateCourseSubject(db, nil)

		sut.Params.ID = courseSubjectDeps.CourseSubject.ID

		courseSubject, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(courseSubject.ID, courseSubjectDeps.CourseSubject.ID, "should have the same ID")
	})

	s.Run("should return an error if the courseSubject does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Params.ID = 404_404_404

		_, err := sut.Sut(sut.Ctx, sut.Params)

		s.Error(err)
	})

	s.Run("should return a courseSubject with filtered fields", func(db *gorm.DB) {
		sut := makeSut(db)

		courseSubjectDeps := fixture.CreateCourseSubject(db, nil)

		sut.Params.ID = courseSubjectDeps.CourseSubject.ID
		sut.Params.Filterer = querying.AddFilter(
			"subjectID",
			querying.EqOperator,
			courseSubjectDeps.CourseSubject.SubjectID,
		)

		courseSubject, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(courseSubject.ID, courseSubjectDeps.CourseSubject.ID, "should have the same ID")
		s.Equal(
			courseSubject.SubjectID,
			courseSubjectDeps.CourseSubject.SubjectID,
			"should have the same subjectID",
		)
	})

	s.Run(
		"should not return a courseSubject if filter does not met the condition",
		func(db *gorm.DB) {
			sut := makeSut(db)

			courseSubjectDeps := fixture.CreateCourseSubject(db, nil)

			sut.Params.ID = courseSubjectDeps.CourseSubject.ID
			sut.Params.Filterer = querying.AddFilter("name", querying.EqOperator, "wrong name")

			courseSubject, err := sut.Sut(sut.Ctx, sut.Params)

			s.Error(err)
			s.Nil(courseSubject, "should not return a courseSubject")
		},
	)
}

func (s *CourseSubjectSuite) TestFindAll() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.FindAllCourseSubjectParams,
			preload ...string,
		) (*queryingPort.PaginationOutput[*model.CourseSubject], error)
		Ctx    context.Context
		Params repo.FindAllCourseSubjectParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.FindAllCourseSubjectParams{
			Paginator: &querying.Pagination{},
			Sorter:    querying.Sort{},
			Filterer:  querying.Filter{},
		}
		sut := s.CourseSubject(db).FindAll

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should find courseSubjects", func(db *gorm.DB) {
		sut := makeSut(db)

		length := 3
		for i := 0; i < length; i++ {
			fixture.CreateCourseSubject(db, nil)
		}

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.NotZero(result.Results[0].ID, "should have a valid id")
		s.Equal(length, result.Total, "should return %d total", length)
		s.Len(result.Results, length, "should return %d results", length)
	})

	s.Run("should return the correct courseSubjects using filter", func(db *gorm.DB) {
		sut := makeSut(db)

		courseSubjectDeps := fixture.CreateCourseSubject(db, nil)
		length := 3
		for i := 0; i < length; i++ {
			fixture.CreateCourseSubject(db, nil)
		}

		sut.Params.Filterer = sut.Params.Filterer.Add(
			"courseID",
			querying.EqOperator,
			courseSubjectDeps.CourseSubject.CourseID,
		)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(result.Results[0].ID, courseSubjectDeps.CourseSubject.ID, "should have the same ID")
		s.Equal(1, result.Total, "should return only one courseSubject")
		s.Len(result.Results, 1, "should return only one courseSubject")
	})

	s.Run("should return the correct courseSubjects using sorter as desc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.CreateCourseSubject(db, nil)
		}

		sut.Params.Sorter = sut.Params.Sorter.Add("id", true)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Greater(int(result.Results[1].ID), int(result.Results[2].ID), "should have the correct order")
	})

	s.Run("should return the correct courseSubjects using sorter as asc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.CreateCourseSubject(db, nil)
		}

		sut.Params.Sorter = sut.Params.Sorter.Add("id", false)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Greater(int(result.Results[2].ID), int(result.Results[1].ID), "should have the correct order")
	})

	s.Run("should return the correct courseSubjects using pagination", func(db *gorm.DB) {
		sut := makeSut(db)

		courseSubjects := make([]*model.CourseSubject, 0)
		for i := 0; i < 3; i++ {
			courseSubjectDeps := fixture.CreateCourseSubject(db, nil)
			courseSubjects = append(courseSubjects, courseSubjectDeps.CourseSubject)
		}

		sut.Params.Paginator = sut.Params.Paginator.SetLimit(1).SetPage(1)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(3, result.Total, "should return the correct total")
		s.Len(result.Results, 1, "should return the correct number of courseSubjects")
		s.Equal(
			int(courseSubjects[0].ID),
			int(result.Results[0].ID),
			"should return the correct courseSubject",
		)
	})
}

func (s *CourseSubjectSuite) TestUpdate() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.UpdateCourseSubjectParams,
		) (*model.CourseSubject, error)
		Ctx    context.Context
		Params repo.UpdateCourseSubjectParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.UpdateCourseSubjectParams{
			CourseSubject: fake.CourseSubject(),
		}
		sut := s.CourseSubject(db).Update

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should update a courseSubject", func(db *gorm.DB) {
		sut := makeSut(db)

		courseSubjectDeps := fixture.CreateCourseSubject(db, nil)
		courseSubjectDeps2 := fixture.CreateCourseSubject(db, nil)

		*sut.Params.CourseSubject = *courseSubjectDeps.CourseSubject
		sut.Params.CourseSubject.CourseID = courseSubjectDeps2.CourseSubject.CourseID

		newCourseSubject, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(courseSubjectDeps.CourseSubject.ID, newCourseSubject.ID, "should return the correct ID")
		s.NotEqual(
			courseSubjectDeps.CourseSubject.CourseID,
			newCourseSubject.CourseID,
			"should return the correct CourseID",
		)
	})

	s.Run(
		"should return an error if tries to update a non existent courseSubject",
		func(db *gorm.DB) {
			sut := makeSut(db)

			fixture.CreateCourseSubject(db, nil)

			sut.Params.CourseSubject.ID = 404_404_404

			newCourseSubject, err := sut.Sut(sut.Ctx, sut.Params)

			s.Error(err)
			s.Nil(newCourseSubject, "should return nil")
		},
	)
}
