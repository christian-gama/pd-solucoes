package middleware_test

import (
	gohttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/christian-gama/pd-solucoes/internal/presentation/middleware"
	"github.com/christian-gama/pd-solucoes/pkg/errutil"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/log"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/gin-gonic/gin"
)

type AddErrorSuite struct {
	suite.Suite
}

func TestAddErrorSuite(t *testing.T) {
	suite.RunUnitTest(t, new(AddErrorSuite))
}

func (s *AddErrorSuite) TestHandle() {
	const (
		method = http.MethodGet
		path   = "/test"
	)

	type Sut struct {
		Sut    middleware.Error
		Logger *mocks.Logger
	}

	makeSut := func() *Sut {
		logger := mocks.NewLogger(s.T())
		sut := middleware.NewError(logger)

		return &Sut{Sut: sut, Logger: logger}
	}

	makeCtxWithReq := func(sut http.Middleware) (*gin.Context, *gin.Engine, *gohttp.Request) {
		w := httptest.NewRecorder()
		ctx, r := gin.CreateTestContext(w)
		r.Use(sut.Handle)
		req := httptest.NewRequest(method, path, nil)

		return ctx, r, req
	}

	s.Run("returns 200 when no error is thrown", func() {
		sut := makeSut()

		ctx, router, req := makeCtxWithReq(sut.Sut)

		router.Handle(method, path, func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		router.ServeHTTP(ctx.Writer, req)

		s.Equal(http.StatusOK, ctx.Writer.Status())
	})

	s.Run("fails and return 500 when an internal error is thrown", func() {
		sut := makeSut()

		ctx, router, req := makeCtxWithReq(sut.Sut)

		router.Handle(method, path, func(c *gin.Context) {
			panic(errutil.NewErrInternal("something went wrong, please try again later"))
		})

		router.ServeHTTP(ctx.Writer, req)

		s.Equal(http.StatusInternalServerError, ctx.Writer.Status())
	})

	s.Run("fails and returns 500 when an errs.Error different from internal is thrown", func() {
		sut := makeSut()

		ctx, router, req := makeCtxWithReq(sut.Sut)

		router.Handle(method, path, func(c *gin.Context) {
			panic(errutil.NewErrInvalid("test", "test"))
		})

		router.ServeHTTP(ctx.Writer, req)

		s.Equal(http.StatusInternalServerError, ctx.Writer.Status())
	})

	s.Run("fails and returns 500 when something different from error is thrown", func() {
		sut := makeSut()

		ctx, router, req := makeCtxWithReq(sut.Sut)

		router.Handle(method, path, func(c *gin.Context) {
			panic("test")
		})

		router.ServeHTTP(ctx.Writer, req)

		s.Equal(http.StatusInternalServerError, ctx.Writer.Status())
	})
}
