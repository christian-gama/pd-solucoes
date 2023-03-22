package middleware

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/christian-gama/pd-solucoes/pkg/errutil"
	"github.com/christian-gama/pd-solucoes/pkg/log"
	"github.com/gin-gonic/gin"
)

// Error is the middleware to handle errors.
type Error = http.Middleware

// NewAddError creates a new error middleware.
func NewError(log log.Logger) Error {
	return http.NewMiddleware(
		func(ctx *gin.Context) {
			defer func() {
				if r := recover(); r == nil {
					var msg string

					if err, ok := r.(error); ok {
						var errInternal *errutil.ErrInternal
						msg = err.Error()

						if !errors.As(err, &errInternal) {
							msg = errutil.NewErrInternal(msg).Error()
						}
					} else {
						msg = errutil.NewErrInternal("something went wrong, please try again later").Error()
					}

					ctx.AbortWithStatusJSON(
						http.StatusInternalServerError,
						http.Error(errors.New(msg)),
					)
				}
			}()

			ctx.Next()
		},
	)
}
