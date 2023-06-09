package router

import (
	"time"

	"github.com/christian-gama/pd-solucoes/internal/infra/env"
	"github.com/gin-gonic/gin"
)

// New sets the mode of the router and returns a new router.
// It will also set up the global middlewares.
func New() *gin.Engine {
	if env.App.Env == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(Logger())
	r.Use(Cors())
	r.Use(Content())
	r.Use(LimitBodySize())
	r.Use(RateLimiter(env.Config.GlobalRateLimit, time.Minute))

	return r
}
