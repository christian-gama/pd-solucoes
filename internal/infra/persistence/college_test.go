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

type CollegeSuite struct {
	suite.SuiteWithConn
	College func(db *gorm.DB) repo.College
}

func TestCollegeSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(CollegeSuite))
}

func (s *CollegeSuite) SetupTest() {
	s.College = func(db *gorm.DB) repo.College {
		return persistence.NewCollege(db)
	}
}

func (s *CollegeSuite) TestCreate() {
	type Sut struct {
		Sut    func(ctx context.Context, params repo.CreateCollegeParams) (*model.College, error)
		Ctx    context.Context
		Params repo.CreateCollegeParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		college := fake.College()
		params := repo.CreateCollegeParams{
			College: college,
		}

		sut := s.College(db).Create

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should create a new college", func(db *gorm.DB) {
		sut := makeSut(db)

		college, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.NotZero(college.ID, "should have an ID")
	})

	s.Run("should return an error when the college already exists", func(db *gorm.DB) {
		sut := makeSut(db)

		_, err := sut.Sut(sut.Ctx, sut.Params)
		s.NoError(err)

		_, err = sut.Sut(sut.Ctx, sut.Params)
		s.Error(err)
	})
}

func (s *CollegeSuite) TestDelete() {
	type Sut struct {
		Sut    func(ctx context.Context, params repo.DeleteCollegeParams) error
		Ctx    context.Context
		Params repo.DeleteCollegeParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.DeleteCollegeParams{}
		sut := s.College(db).Delete

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should delete a college", func(db *gorm.DB) {
		sut := makeSut(db)

		collegeDeps := fixture.CreateCollege(db, nil)

		sut.Params.ID = collegeDeps.College.ID

		err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
	})

	s.Run("should delete nothing if the college does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Params.ID = 404_404_404

		err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
	})
}

func (s *CollegeSuite) TestFindOne() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.FindOneCollegeParams,
			preload ...string,
		) (*model.College, error)
		Ctx    context.Context
		Params repo.FindOneCollegeParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.FindOneCollegeParams{
			ID: 0,
		}
		sut := s.College(db).FindOne

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should find a college", func(db *gorm.DB) {
		sut := makeSut(db)

		collegeDeps := fixture.CreateCollege(db, nil)

		sut.Params.ID = collegeDeps.College.ID

		college, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(college.ID, collegeDeps.College.ID, "should have the same ID")
	})

	s.Run("should return an error if the college does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Params.ID = 404_404_404

		_, err := sut.Sut(sut.Ctx, sut.Params)

		s.Error(err)
	})
}

func (s *CollegeSuite) TestFindAll() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.FindAllCollegeParams,
			preload ...string,
		) (*queryingPort.PaginationOutput[*model.College], error)
		Ctx    context.Context
		Params repo.FindAllCollegeParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.FindAllCollegeParams{
			Paginator: &querying.Pagination{},
			Sorter:    querying.Sort{},
			Filterer:  querying.Filter{},
		}
		sut := s.College(db).FindAll

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should find colleges", func(db *gorm.DB) {
		sut := makeSut(db)

		length := 3
		for i := 0; i < length; i++ {
			fixture.CreateCollege(db, nil)
		}

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.NotZero(result.Results[0].ID, "should have a valid id")
		s.Equal(length, result.Total, "should return %d total", length)
		s.Len(result.Results, length, "should return %d results", length)
	})

	s.Run("should return the correct colleges using filter", func(db *gorm.DB) {
		sut := makeSut(db)

		collegeDeps := fixture.CreateCollege(db, nil)
		length := 3
		for i := 0; i < length; i++ {
			fixture.CreateCollege(db, nil)
		}

		sut.Params.Filterer = sut.Params.Filterer.Add(
			"name",
			querying.EqOperator,
			collegeDeps.College.Name,
		)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(result.Results[0].ID, collegeDeps.College.ID, "should have the same ID")
		s.Equal(1, result.Total, "should return only one college")
		s.Len(result.Results, 1, "should return only one college")
	})

	s.Run("should return the correct colleges using sorter as desc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.CreateCollege(db, nil)
		}

		sut.Params.Sorter = sut.Params.Sorter.Add("id", true)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Greater(int(result.Results[1].ID), int(result.Results[2].ID), "should have the correct order")
	})

	s.Run("should return the correct colleges using sorter as asc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.CreateCollege(db, nil)
		}

		sut.Params.Sorter = sut.Params.Sorter.Add("id", false)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Greater(int(result.Results[2].ID), int(result.Results[1].ID), "should have the correct order")
	})

	s.Run("should return the correct colleges using pagination", func(db *gorm.DB) {
		sut := makeSut(db)

		colleges := make([]*model.College, 0)
		for i := 0; i < 3; i++ {
			collegeDeps := fixture.CreateCollege(db, nil)
			colleges = append(colleges, collegeDeps.College)
		}

		sut.Params.Paginator = sut.Params.Paginator.SetLimit(1).SetPage(1)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(3, result.Total, "should return the correct total")
		s.Len(result.Results, 1, "should return the correct number of colleges")
		s.Equal(int(colleges[0].ID), int(result.Results[0].ID), "should return the correct college")
	})
}

func (s *CollegeSuite) TestUpdate() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.UpdateCollegeParams,
		) (*model.College, error)
		Ctx    context.Context
		Params repo.UpdateCollegeParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.UpdateCollegeParams{
			College: fake.College(),
		}
		sut := s.College(db).Update

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should update a college", func(db *gorm.DB) {
		sut := makeSut(db)

		collegeDeps := fixture.CreateCollege(db, nil)

		*sut.Params.College = *collegeDeps.College
		sut.Params.College.Name = "new name"

		newCollege, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(collegeDeps.College.ID, newCollege.ID, "should return the correct ID")
		s.NotEqual(collegeDeps.College.Name, newCollege.Name, "should return the correct name")
	})

	s.Run("should return an error if tries to update a non existent college", func(db *gorm.DB) {
		sut := makeSut(db)

		fixture.CreateCollege(db, nil)

		sut.Params.College.Name = "new name"
		sut.Params.College.ID = 404_404_404

		newCollege, err := sut.Sut(sut.Ctx, sut.Params)

		s.Error(err)
		s.Nil(newCollege, "should return nil")
	})
}
