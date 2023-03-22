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

type StudentSuite struct {
	suite.SuiteWithConn
	Student func(db *gorm.DB) repo.Student
}

func TestStudentSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(StudentSuite))
}

func (s *StudentSuite) SetupTest() {
	s.Student = func(db *gorm.DB) repo.Student {
		return persistence.NewStudent(db)
	}
}

func (s *StudentSuite) TestCreate() {
	type Sut struct {
		Sut    func(ctx context.Context, params repo.CreateStudentParams) (*model.Student, error)
		Ctx    context.Context
		Params repo.CreateStudentParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		student := fake.Student()
		params := repo.CreateStudentParams{
			Student: student,
		}

		sut := s.Student(db).Create

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should create a new student", func(db *gorm.DB) {
		sut := makeSut(db)

		student, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.NotZero(student.ID, "should have an ID")
	})

	s.Run("should return an error when the student already exists", func(db *gorm.DB) {
		sut := makeSut(db)

		_, err := sut.Sut(sut.Ctx, sut.Params)
		s.NoError(err)

		_, err = sut.Sut(sut.Ctx, sut.Params)
		s.Error(err)
	})
}

func (s *StudentSuite) TestDelete() {
	type Sut struct {
		Sut    func(ctx context.Context, params repo.DeleteStudentParams) error
		Ctx    context.Context
		Params repo.DeleteStudentParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.DeleteStudentParams{}
		sut := s.Student(db).Delete

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should delete a student", func(db *gorm.DB) {
		sut := makeSut(db)

		studentDeps := fixture.CreateStudent(db, nil)

		sut.Params.ID = studentDeps.Student.ID

		err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
	})

	s.Run("should delete nothing if the student does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Params.ID = 404_404_404

		err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
	})
}

func (s *StudentSuite) TestFindOne() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.FindOneStudentParams,
		) (*model.Student, error)
		Ctx    context.Context
		Params repo.FindOneStudentParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.FindOneStudentParams{
			Filterer: querying.Filter{},
			ID:       0,
		}
		sut := s.Student(db).FindOne

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should find a student", func(db *gorm.DB) {
		sut := makeSut(db)

		studentDeps := fixture.CreateStudent(db, nil)

		sut.Params.ID = studentDeps.Student.ID

		student, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(student.ID, studentDeps.Student.ID, "should have the same ID")
	})

	s.Run("should return an error if the student does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Params.ID = 404_404_404

		_, err := sut.Sut(sut.Ctx, sut.Params)

		s.Error(err)
	})

	s.Run("should return a student with filtered fields", func(db *gorm.DB) {
		sut := makeSut(db)

		studentDeps := fixture.CreateStudent(db, nil)

		sut.Params.ID = studentDeps.Student.ID
		sut.Params.Filterer = querying.AddFilter("name", querying.EqOperator, studentDeps.Student.Name)

		student, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(student.ID, studentDeps.Student.ID, "should have the same ID")
		s.Equal(student.Name, studentDeps.Student.Name, "should have the same name")
	})

	s.Run("should not return a student if filter does not met the condition", func(db *gorm.DB) {
		sut := makeSut(db)

		studentDeps := fixture.CreateStudent(db, nil)

		sut.Params.ID = studentDeps.Student.ID
		sut.Params.Filterer = querying.AddFilter("name", querying.EqOperator, "wrong name")

		student, err := sut.Sut(sut.Ctx, sut.Params)

		s.Error(err)
		s.Nil(student, "should not return a student")
	})
}

func (s *StudentSuite) TestFindAll() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.FindAllStudentParams,
		) (*queryingPort.PaginationOutput[*model.Student], error)
		Ctx    context.Context
		Params repo.FindAllStudentParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.FindAllStudentParams{
			Paginator: &querying.Pagination{},
			Sorter:    querying.Sort{},
			Filterer:  querying.Filter{},
		}
		sut := s.Student(db).FindAll

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should find students", func(db *gorm.DB) {
		sut := makeSut(db)

		length := 3
		for i := 0; i < length; i++ {
			fixture.CreateStudent(db, nil)
		}

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.NotZero(result.Results[0].ID, "should have a valid id")
		s.Equal(length, result.Total, "should return %d total", length)
		s.Len(result.Results, length, "should return %d results", length)
	})

	s.Run("should return the correct students using filter", func(db *gorm.DB) {
		sut := makeSut(db)

		studentDeps := fixture.CreateStudent(db, nil)
		length := 3
		for i := 0; i < length; i++ {
			fixture.CreateStudent(db, nil)
		}

		sut.Params.Filterer = sut.Params.Filterer.Add(
			"name",
			querying.EqOperator,
			studentDeps.Student.Name,
		)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(result.Results[0].ID, studentDeps.Student.ID, "should have the same ID")
		s.Equal(1, result.Total, "should return only one student")
		s.Len(result.Results, 1, "should return only one student")
	})

	s.Run("should return the correct students using sorter as desc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.CreateStudent(db, nil)
		}

		sut.Params.Sorter = sut.Params.Sorter.Add("id", true)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Greater(int(result.Results[1].ID), int(result.Results[2].ID), "should have the correct order")
	})

	s.Run("should return the correct students using sorter as asc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.CreateStudent(db, nil)
		}

		sut.Params.Sorter = sut.Params.Sorter.Add("id", false)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Greater(int(result.Results[2].ID), int(result.Results[1].ID), "should have the correct order")
	})

	s.Run("should return the correct students using pagination", func(db *gorm.DB) {
		sut := makeSut(db)

		students := make([]*model.Student, 0)
		for i := 0; i < 3; i++ {
			studentDeps := fixture.CreateStudent(db, nil)
			students = append(students, studentDeps.Student)
		}

		sut.Params.Paginator = sut.Params.Paginator.SetLimit(1).SetPage(1)

		result, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(3, result.Total, "should return the correct total")
		s.Len(result.Results, 1, "should return the correct number of students")
		s.Equal(int(students[0].ID), int(result.Results[0].ID), "should return the correct student")
	})
}

func (s *StudentSuite) TestUpdate() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			params repo.UpdateStudentParams,
		) (*model.Student, error)
		Ctx    context.Context
		Params repo.UpdateStudentParams
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		params := repo.UpdateStudentParams{
			Student: fake.Student(),
		}
		sut := s.Student(db).Update

		return Sut{
			Sut:    sut,
			Ctx:    ctx,
			Params: params,
		}
	}

	s.Run("should update a student", func(db *gorm.DB) {
		sut := makeSut(db)

		studentDeps := fixture.CreateStudent(db, nil)

		*sut.Params.Student = *studentDeps.Student
		sut.Params.Student.Name = "new name"

		newStudent, err := sut.Sut(sut.Ctx, sut.Params)

		s.NoError(err)
		s.Equal(studentDeps.Student.ID, newStudent.ID, "should return the correct ID")
		s.NotEqual(studentDeps.Student.Name, newStudent.Name, "should return the correct name")
	})

	s.Run("should return an error if tries to update a non existent student", func(db *gorm.DB) {
		sut := makeSut(db)

		fixture.CreateStudent(db, nil)

		sut.Params.Student.Name = "new name"
		sut.Params.Student.ID = 404_404_404

		newStudent, err := sut.Sut(sut.Ctx, sut.Params)

		s.Error(err)
		s.Nil(newStudent, "should return nil")
	})
}
