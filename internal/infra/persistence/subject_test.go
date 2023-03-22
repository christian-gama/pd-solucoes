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

type SubjectSuite struct {
	suite.SuiteWithConn
	Subject func(db *gorm.DB) repo.Subject
}

func TestSubjectSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(SubjectSuite))
}

func (s *SubjectSuite) SetupTest() {
	s.Subject = func(db *gorm.DB) repo.Subject {
		return persistence.NewSubject(db)
	}
}

func (s *SubjectSuite) TestCreate() {
	type Sut struct {
		Sut    func(ctx context.Context, params repo.CreateSubjectParams) (*model.Subject, error)
		Ctx    context.Context
		Params repo.CreateSubjectParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		subject := fake.Subject()
		params := repo.CreateSubjectParams{
			Subject: subject,
		}

		sut := s.Subject(db).Create

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should create a new subject", func(db *gorm.DB) {
		sut := makeSut(db)

		teacherDeps := fixture.CreateTeacher(db, nil)

		sut.Params.Subject.TeacherID = teacherDeps.Teacher.ID

		model, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.NotZero(model.ID, "should have an ID")
	})

	s.Run("should return an error when the subject already exists", func(db *gorm.DB) {
		sut := makeSut(db)

		teacherDeps := fixture.CreateTeacher(db, nil)

		sut.Params.Subject.TeacherID = teacherDeps.Teacher.ID

		_, err := sut.Sut(sut.Ctx, sut.Params)
		s.NoError(err)

		_, err = sut.Sut(sut.Ctx, sut.Params)
		s.Error(err)
	})
}

func (s *SubjectSuite) TestDelete() {
	type Sut struct {
		Sut    func(ctx context.Context, params repo.DeleteSubjectParams) error
		Ctx    context.Context
		Params repo.DeleteSubjectParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.DeleteSubjectParams{}
		sut := s.Subject(db).Delete

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should delete a subject", func(db *gorm.DB) {
		sut := makeSut(db)

		subjectDeps := fixture.CreateSubject(db, nil)

		sut.Params.ID = subjectDeps.Subject.ID

		err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
	})

	s.Run("should delete nothing if the subject does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Params.ID = 404_404_404

		err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
	})
}

func (s *SubjectSuite) TestFindOne() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.FindOneSubjectParams,
			preload ...string,
		) (*model.Subject, error)
		Ctx    context.Context
		Params repo.FindOneSubjectParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.FindOneSubjectParams{
			ID: 0,
		}
		sut := s.Subject(db).FindOne

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should find a subject", func(db *gorm.DB) {
		sut := makeSut(db)

		subjectDeps := fixture.CreateSubject(db, nil)

		sut.Params.ID = subjectDeps.Subject.ID

		subject, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(subject.ID, subjectDeps.Subject.ID, "should have the same ID")
	})

	s.Run("should return an error if the subject does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Params.ID = 404_404_404

		_, err := sut.Sut(sut.Ctx, sut.Params)

		s.Error(err)
	})
}

func (s *SubjectSuite) TestFindAll() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.FindAllSubjectParams,
			preload ...string,
		) (*queryingPort.PaginationOutput[*model.Subject], error)
		Ctx    context.Context
		Params repo.FindAllSubjectParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.FindAllSubjectParams{
			Paginator: &querying.Pagination{},
			Sorter:    querying.Sort{},
			Filterer:  querying.Filter{},
		}
		sut := s.Subject(db).FindAll

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should find subjects", func(db *gorm.DB) {
		sut := makeSut(db)

		length := 3
		for i := 0; i < length; i++ {
			fixture.CreateSubject(db, nil)
		}

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.NotZero(result.Results[0].ID, "should have a valid id")
		s.Equal(length, result.Total, "should return %d total", length)
		s.Len(result.Results, length, "should return %d results", length)
	})

	s.Run("should return the correct subjects using filter", func(db *gorm.DB) {
		sut := makeSut(db)

		subjectDeps := fixture.CreateSubject(db, nil)
		length := 3
		for i := 0; i < length; i++ {
			fixture.CreateSubject(db, nil)
		}

		sut.Params.Filterer = sut.Params.Filterer.Add(
			"name",
			querying.EqOperator,
			subjectDeps.Subject.Name,
		)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(result.Results[0].ID, subjectDeps.Subject.ID, "should have the same ID")
		s.Equal(1, result.Total, "should return only one subject")
		s.Len(result.Results, 1, "should return only one subject")
	})

	s.Run("should return the correct subjects using sorter as desc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.CreateSubject(db, nil)
		}

		sut.Params.Sorter = sut.Params.Sorter.Add("id", true)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Greater(int(result.Results[1].ID), int(result.Results[2].ID), "should have the correct order")
	})

	s.Run("should return the correct subjects using sorter as asc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.CreateSubject(db, nil)
		}

		sut.Params.Sorter = sut.Params.Sorter.Add("id", false)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Greater(int(result.Results[2].ID), int(result.Results[1].ID), "should have the correct order")
	})

	s.Run("should return the correct subjects using pagination", func(db *gorm.DB) {
		sut := makeSut(db)

		subjects := make([]*model.Subject, 0)
		for i := 0; i < 3; i++ {
			subjectDeps := fixture.CreateSubject(db, nil)
			subjects = append(subjects, subjectDeps.Subject)
		}

		sut.Params.Paginator = sut.Params.Paginator.SetLimit(1).SetPage(1)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(3, result.Total, "should return the correct total")
		s.Len(result.Results, 1, "should return the correct number of subjects")
		s.Equal(int(subjects[0].ID), int(result.Results[0].ID), "should return the correct subject")
	})
}

func (s *SubjectSuite) TestUpdate() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.UpdateSubjectParams,
		) (*model.Subject, error)
		Ctx    context.Context
		Params repo.UpdateSubjectParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.UpdateSubjectParams{
			Subject: fake.Subject(),
		}
		sut := s.Subject(db).Update

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should update a subject", func(db *gorm.DB) {
		sut := makeSut(db)

		subjectDeps := fixture.CreateSubject(db, nil)

		*sut.Params.Subject = *subjectDeps.Subject
		sut.Params.Subject.Name = "new name"

		newSubject, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(subjectDeps.Subject.ID, newSubject.ID, "should return the correct ID")
		s.NotEqual(subjectDeps.Subject.Name, newSubject.Name, "should return the correct name")
	})

	s.Run("should return an error if tries to update a non existent subject", func(db *gorm.DB) {
		sut := makeSut(db)

		fixture.CreateSubject(db, nil)

		sut.Params.Subject.Name = "new name"
		sut.Params.Subject.ID = 404_404_404

		newSubject, err := sut.Sut(sut.Ctx, sut.Params)

		s.Error(err)
		s.Nil(newSubject, "should return nil")
	})
}
