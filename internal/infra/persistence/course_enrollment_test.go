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

type CourseEnrollmentSuite struct {
	suite.SuiteWithConn
	CourseEnrollment func(db *gorm.DB) repo.CourseEnrollment
}

func TestCourseEnrollmentSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(CourseEnrollmentSuite))
}

func (s *CourseEnrollmentSuite) SetupTest() {
	s.CourseEnrollment = func(db *gorm.DB) repo.CourseEnrollment {
		return persistence.NewCourseEnrollment(db)
	}
}

func (s *CourseEnrollmentSuite) TestCreate() {
	type Sut struct {
		Sut    func(context.Context, repo.CreateCourseEnrollmentParams) (*model.CourseEnrollment, error)
		Ctx    context.Context
		Params repo.CreateCourseEnrollmentParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		courseEnrollment := fake.CourseEnrollment()
		params := repo.CreateCourseEnrollmentParams{
			CourseEnrollment: courseEnrollment,
		}

		sut := s.CourseEnrollment(db).Create

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should create a new courseEnrollment", func(db *gorm.DB) {
		sut := makeSut(db)

		courseSubjectDeps := fixture.CreateCourseSubject(db, nil)
		studentDeps := fixture.CreateStudent(db, nil)

		sut.Params.CourseEnrollment.CourseSubjectID = courseSubjectDeps.CourseSubject.ID
		sut.Params.CourseEnrollment.StudentID = studentDeps.Student.ID

		model, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.NotZero(model.ID, "should have an ID")
	})

	s.Run("should return an error when the courseEnrollment already exists", func(db *gorm.DB) {
		sut := makeSut(db)

		courseSubjectDeps := fixture.CreateCourseSubject(db, nil)
		studentDeps := fixture.CreateStudent(db, nil)

		sut.Params.CourseEnrollment.CourseSubjectID = courseSubjectDeps.CourseSubject.ID
		sut.Params.CourseEnrollment.StudentID = studentDeps.Student.ID

		_, err := sut.Sut(sut.Ctx, sut.Params)
		s.NoError(err)

		_, err = sut.Sut(sut.Ctx, sut.Params)
		s.Error(err)
	})
}

func (s *CourseEnrollmentSuite) TestDelete() {
	type Sut struct {
		Sut    func(ctx context.Context, params repo.DeleteCourseEnrollmentParams) error
		Ctx    context.Context
		Params repo.DeleteCourseEnrollmentParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.DeleteCourseEnrollmentParams{}
		sut := s.CourseEnrollment(db).Delete

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should delete a courseEnrollment", func(db *gorm.DB) {
		sut := makeSut(db)

		courseEnrollmentDeps := fixture.CreateCourseEnrollment(db, nil)

		sut.Params.ID = courseEnrollmentDeps.CourseEnrollment.ID

		err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
	})

	s.Run("should delete nothing if the courseEnrollment does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Params.ID = 404_404_404

		err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
	})
}

func (s *CourseEnrollmentSuite) TestFindOne() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.FindOneCourseEnrollmentParams,
			preload ...string,
		) (*model.CourseEnrollment, error)
		Ctx    context.Context
		Params repo.FindOneCourseEnrollmentParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.FindOneCourseEnrollmentParams{
			ID: 0,
		}
		sut := s.CourseEnrollment(db).FindOne

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should find a courseEnrollment", func(db *gorm.DB) {
		sut := makeSut(db)

		courseEnrollmentDeps := fixture.CreateCourseEnrollment(db, nil)

		sut.Params.ID = courseEnrollmentDeps.CourseEnrollment.ID

		courseEnrollment, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(
			courseEnrollment.ID,
			courseEnrollmentDeps.CourseEnrollment.ID,
			"should have the same ID",
		)
	})

	s.Run("should return an error if the courseEnrollment does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Params.ID = 404_404_404

		_, err := sut.Sut(sut.Ctx, sut.Params)

		s.Error(err)
	})
}

func (s *CourseEnrollmentSuite) TestFindAll() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.FindAllCourseEnrollmentParams,
			preload ...string,
		) (*queryingPort.PaginationOutput[*model.CourseEnrollment], error)
		Ctx    context.Context
		Params repo.FindAllCourseEnrollmentParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.FindAllCourseEnrollmentParams{
			Paginator: &querying.Pagination{},
			Sorter:    querying.Sort{},
			Filterer:  querying.Filter{},
		}
		sut := s.CourseEnrollment(db).FindAll

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should find courseEnrollments", func(db *gorm.DB) {
		sut := makeSut(db)

		length := 3
		for i := 0; i < length; i++ {
			fixture.CreateCourseEnrollment(db, nil)
		}

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.NotZero(result.Results[0].ID, "should have a valid id")
		s.Equal(length, result.Total, "should return %d total", length)
		s.Len(result.Results, length, "should return %d results", length)
	})

	s.Run("should return the correct courseEnrollments using filter", func(db *gorm.DB) {
		sut := makeSut(db)

		courseEnrollmentDeps := fixture.CreateCourseEnrollment(db, nil)
		length := 3
		for i := 0; i < length; i++ {
			fixture.CreateCourseEnrollment(db, nil)
		}

		sut.Params.Filterer = sut.Params.Filterer.Add(
			"courseSubjectID",
			querying.EqOperator,
			courseEnrollmentDeps.CourseEnrollment.CourseSubjectID,
		)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(
			result.Results[0].ID,
			courseEnrollmentDeps.CourseEnrollment.ID,
			"should have the same ID",
		)
		s.Equal(1, result.Total, "should return only one courseEnrollment")
		s.Len(result.Results, 1, "should return only one courseEnrollment")
	})

	s.Run("should return the correct courseEnrollments using sorter as desc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.CreateCourseEnrollment(db, nil)
		}

		sut.Params.Sorter = sut.Params.Sorter.Add("id", true)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Greater(int(result.Results[1].ID), int(result.Results[2].ID), "should have the correct order")
	})

	s.Run("should return the correct courseEnrollments using sorter as asc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.CreateCourseEnrollment(db, nil)
		}

		sut.Params.Sorter = sut.Params.Sorter.Add("id", false)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Greater(int(result.Results[2].ID), int(result.Results[1].ID), "should have the correct order")
	})

	s.Run("should return the correct courseEnrollments using pagination", func(db *gorm.DB) {
		sut := makeSut(db)

		courseEnrollments := make([]*model.CourseEnrollment, 0)
		for i := 0; i < 3; i++ {
			courseEnrollmentDeps := fixture.CreateCourseEnrollment(db, nil)
			courseEnrollments = append(courseEnrollments, courseEnrollmentDeps.CourseEnrollment)
		}

		sut.Params.Paginator = sut.Params.Paginator.SetLimit(1).SetPage(1)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(3, result.Total, "should return the correct total")
		s.Len(result.Results, 1, "should return the correct number of courseEnrollments")
		s.Equal(
			int(courseEnrollments[0].ID),
			int(result.Results[0].ID),
			"should return the correct courseEnrollment",
		)
	})
}

func (s *CourseEnrollmentSuite) TestUpdate() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.UpdateCourseEnrollmentParams,
		) (*model.CourseEnrollment, error)
		Ctx    context.Context
		Params repo.UpdateCourseEnrollmentParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.UpdateCourseEnrollmentParams{
			CourseEnrollment: fake.CourseEnrollment(),
		}
		sut := s.CourseEnrollment(db).Update

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should update a courseEnrollment", func(db *gorm.DB) {
		sut := makeSut(db)

		courseEnrollmentDeps := fixture.CreateCourseEnrollment(db, nil)
		courseEnrollmentDeps2 := fixture.CreateCourseEnrollment(db, nil)

		*sut.Params.CourseEnrollment = *courseEnrollmentDeps.CourseEnrollment
		sut.Params.CourseEnrollment.CourseSubjectID = courseEnrollmentDeps2.CourseEnrollment.CourseSubjectID

		newCourseEnrollment, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(
			courseEnrollmentDeps.CourseEnrollment.ID,
			newCourseEnrollment.ID,
			"should return the correct ID",
		)
		s.NotEqual(
			courseEnrollmentDeps.CourseEnrollment.CourseSubjectID,
			newCourseEnrollment.CourseSubjectID,
			"should return the correct CourseID",
		)
	})

	s.Run(
		"should return an error if tries to update a non existent courseEnrollment",
		func(db *gorm.DB) {
			sut := makeSut(db)

			fixture.CreateCourseEnrollment(db, nil)

			sut.Params.CourseEnrollment.ID = 404_404_404

			newCourseEnrollment, err := sut.Sut(sut.Ctx, sut.Params)

			s.Error(err)
			s.Nil(newCourseEnrollment, "should return nil")
		},
	)
}
