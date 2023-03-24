package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/courseSubject"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocksCSubjectService "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/courseSubject"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CreateCourseSubjectSuite struct {
	suite.Suite
}

func TestCreateCourseSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CreateCourseSubjectSuite))
}

func (s *CreateCourseSubjectSuite) TestHandle() {
	type Sut struct {
		Sut                         service.CreateCourseSubject
		CourseSubjectRepo           *mocks.CourseSubject
		Input                       *service.CreateCourseSubjectInput
		CourseSubject               *model.CourseSubject
		FindOneCourseSubjectService *mocksCSubjectService.FindOneCourseSubject
	}

	makeSut := func() *Sut {
		courseSubjectRepo := mocks.NewCourseSubject(s.T())
		findOneCourseSubjectService := mocksCSubjectService.NewFindOneCourseSubject(
			s.T(),
		)
		sut := service.NewCreateCourseSubject(
			courseSubjectRepo,
			findOneCourseSubjectService,
		)

		return &Sut{
			Sut:                         sut,
			CourseSubjectRepo:           courseSubjectRepo,
			Input:                       fake.CreateCourseSubjectInput(),
			CourseSubject:               fakeModel.CourseSubject(),
			FindOneCourseSubjectService: findOneCourseSubjectService,
		}
	}

	s.Run("should add a new courseSubject", func() {
		sut := makeSut()

		sut.CourseSubjectRepo.
			On("Create", mock.Anything, mock.Anything).
			Return(sut.CourseSubject, nil)

		sut.FindOneCourseSubjectService.
			On("Handle", mock.Anything, mock.Anything).
			Return(copy.MustCopy(&service.Output{}, sut.CourseSubject), nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.CourseSubject.ID, result.ID)
		s.Equal(sut.CourseSubject.SubjectID, result.Subject.ID)
	})

	s.Run("courseSubjectRepo.Create returns an error", func() {
		sut := makeSut()

		sut.CourseSubjectRepo.
			On("Create", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})

	s.Run("findOneCourseSubject.Handle returns an error", func() {
		sut := makeSut()

		sut.CourseSubjectRepo.
			On("Create", mock.Anything, mock.Anything).
			Return(sut.CourseSubject, nil)

		sut.FindOneCourseSubjectService.
			On("Handle", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})

	s.Run("should return an error if domain validation fails", func() {
		sut := makeSut()

		sut.Input.SubjectID = 0

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.Error(err)
		s.Nil(result)
	})
}
