package router

import (
	"errors"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/christian-gama/pd-solucoes/pkg/bench"
	"github.com/christian-gama/pd-solucoes/pkg/unit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Logger returns a gin middleware that logs the request.
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		duration := bench.Duration(ctx.Next)

		logLevel(ctx.Writer.Status())(
			"%-6s | %-5s | %4dms | %s",
			ctx.Request.Method,
			statusColor(ctx.Writer.Status()),
			duration.Milliseconds(),
			ctx.Request.URL.Path,
		)
	}
}

// Cors returns a gin middleware that enables CORS.
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowFiles:      true,
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods:    []string{"GET", "POST", "DELETE", "PATCH"},
	})
}

// LimitBodySize returns a gin middleware that limits the request body size.
func LimitBodySize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const maxBodySize = 3 * unit.Megabyte

		if ctx.Request.ContentLength > maxBodySize {
			http.BadRequest(ctx, errors.New("request body too large"))
			return
		}

		ctx.Next()
	}
}

// Content returns a gin middleware that sets the content type.
func Content() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		ctx.Next()
	}
}

func RateLimiter(limit uint, duration time.Duration) gin.HandlerFunc {
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  duration,
		Limit: limit,
	})

	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: func(ctx *gin.Context, info ratelimit.Info) {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, http.Error(
				errors.New("too many requests"),
			))
		},

		KeyFunc: func(ctx *gin.Context) string {
			return ctx.ClientIP()
		},
	})

	return mw
}
