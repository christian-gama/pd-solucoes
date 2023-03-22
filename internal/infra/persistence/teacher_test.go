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

type TeacherSuite struct {
	suite.SuiteWithConn
	Teacher func(db *gorm.DB) repo.Teacher
}

func TestTeacherSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(TeacherSuite))
}

func (s *TeacherSuite) SetupTest() {
	s.Teacher = func(db *gorm.DB) repo.Teacher {
		return persistence.NewTeacher(db)
	}
}

func (s *TeacherSuite) TestCreate() {
	type Sut struct {
		Sut    func(ctx context.Context, params repo.CreateTeacherParams) (*model.Teacher, error)
		Ctx    context.Context
		Params repo.CreateTeacherParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		teacher := fake.Teacher()
		params := repo.CreateTeacherParams{
			Teacher: teacher,
		}

		sut := s.Teacher(db).Create

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should create a new teacher", func(db *gorm.DB) {
		sut := makeSut(db)

		teacher, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.NotZero(teacher.ID, "should have an ID")
	})

	s.Run("should return an error when the teacher already exists", func(db *gorm.DB) {
		sut := makeSut(db)

		_, err := sut.Sut(sut.Ctx, sut.Params)
		s.NoError(err)

		_, err = sut.Sut(sut.Ctx, sut.Params)
		s.Error(err)
	})
}

func (s *TeacherSuite) TestDelete() {
	type Sut struct {
		Sut    func(ctx context.Context, params repo.DeleteTeacherParams) error
		Ctx    context.Context
		Params repo.DeleteTeacherParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.DeleteTeacherParams{}
		sut := s.Teacher(db).Delete

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should delete a teacher", func(db *gorm.DB) {
		sut := makeSut(db)

		teacherDeps := fixture.CreateTeacher(db, nil)

		sut.Params.ID = teacherDeps.Teacher.ID

		err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
	})

	s.Run("should delete nothing if the teacher does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Params.ID = 404_404_404

		err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
	})
}

func (s *TeacherSuite) TestFindOne() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.FindOneTeacherParams,
		) (*model.Teacher, error)
		Ctx    context.Context
		Params repo.FindOneTeacherParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.FindOneTeacherParams{
			ID: 0,
		}
		sut := s.Teacher(db).FindOne

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should find a teacher", func(db *gorm.DB) {
		sut := makeSut(db)

		teacherDeps := fixture.CreateTeacher(db, nil)

		sut.Params.ID = teacherDeps.Teacher.ID

		teacher, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(teacher.ID, teacherDeps.Teacher.ID, "should have the same ID")
	})

	s.Run("should return an error if the teacher does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Params.ID = 404_404_404

		_, err := sut.Sut(sut.Ctx, sut.Params)

		s.Error(err)
	})
}

func (s *TeacherSuite) TestFindAll() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.FindAllTeacherParams,
		) (*queryingPort.PaginationOutput[*model.Teacher], error)
		Ctx    context.Context
		Params repo.FindAllTeacherParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.FindAllTeacherParams{
			Paginator: &querying.Pagination{},
			Sorter:    querying.Sort{},
			Filterer:  querying.Filter{},
		}
		sut := s.Teacher(db).FindAll

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should find teachers", func(db *gorm.DB) {
		sut := makeSut(db)

		length := 3
		for i := 0; i < length; i++ {
			fixture.CreateTeacher(db, nil)
		}

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.NotZero(result.Results[0].ID, "should have a valid id")
		s.Equal(length, result.Total, "should return %d total", length)
		s.Len(result.Results, length, "should return %d results", length)
	})

	s.Run("should return the correct teachers using filter", func(db *gorm.DB) {
		sut := makeSut(db)

		teacherDeps := fixture.CreateTeacher(db, nil)
		length := 3
		for i := 0; i < length; i++ {
			fixture.CreateTeacher(db, nil)
		}

		sut.Params.Filterer = sut.Params.Filterer.Add(
			"name",
			querying.EqOperator,
			teacherDeps.Teacher.Name,
		)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(result.Results[0].ID, teacherDeps.Teacher.ID, "should have the same ID")
		s.Equal(1, result.Total, "should return only one teacher")
		s.Len(result.Results, 1, "should return only one teacher")
	})

	s.Run("should return the correct teachers using sorter as desc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.CreateTeacher(db, nil)
		}

		sut.Params.Sorter = sut.Params.Sorter.Add("id", true)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Greater(int(result.Results[1].ID), int(result.Results[2].ID), "should have the correct order")
	})

	s.Run("should return the correct teachers using sorter as asc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.CreateTeacher(db, nil)
		}

		sut.Params.Sorter = sut.Params.Sorter.Add("id", false)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Greater(int(result.Results[2].ID), int(result.Results[1].ID), "should have the correct order")
	})

	s.Run("should return the correct teachers using pagination", func(db *gorm.DB) {
		sut := makeSut(db)

		teachers := make([]*model.Teacher, 0)
		for i := 0; i < 3; i++ {
			teacherDeps := fixture.CreateTeacher(db, nil)
			teachers = append(teachers, teacherDeps.Teacher)
		}

		sut.Params.Paginator = sut.Params.Paginator.SetLimit(1).SetPage(1)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(3, result.Total, "should return the correct total")
		s.Len(result.Results, 1, "should return the correct number of teachers")
		s.Equal(int(teachers[0].ID), int(result.Results[0].ID), "should return the correct teacher")
	})
}

func (s *TeacherSuite) TestUpdate() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.UpdateTeacherParams,
		) (*model.Teacher, error)
		Ctx    context.Context
		Params repo.UpdateTeacherParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.UpdateTeacherParams{
			Teacher: fake.Teacher(),
		}
		sut := s.Teacher(db).Update

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should update a teacher", func(db *gorm.DB) {
		sut := makeSut(db)

		teacherDeps := fixture.CreateTeacher(db, nil)

		*sut.Params.Teacher = *teacherDeps.Teacher
		sut.Params.Teacher.Name = "new name"

		newTeacher, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(teacherDeps.Teacher.ID, newTeacher.ID, "should return the correct ID")
		s.NotEqual(teacherDeps.Teacher.Name, newTeacher.Name, "should return the correct name")
	})

	s.Run("should return an error if tries to update a non existent teacher", func(db *gorm.DB) {
		sut := makeSut(db)

		fixture.CreateTeacher(db, nil)

		sut.Params.Teacher.Name = "new name"
		sut.Params.Teacher.ID = 404_404_404

		newTeacher, err := sut.Sut(sut.Ctx, sut.Params)

		s.Error(err)
		s.Nil(newTeacher, "should return nil")
	})
}
