package routing

import (
	"errors"
	"fmt"
	"strings"

	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/christian-gama/pd-solucoes/pkg/slice"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Middlewares []http.Middleware
	Controller  http.Controller
}

type Routing struct {
	Group       string
	Routes      []*Route
	Middlewares []http.Middleware
}

func (r *Routing) validate(route *Route) {
	if route.Controller == nil {
		panic(errors.New("controller is nil"))
	}

	if route.Controller.Method() == "" {
		panic(fmt.Errorf("method is empty for controller %v", route.Controller))
	}

	if route.Controller.Path() == "" {
		panic(fmt.Errorf("path is empty for controller %v", route.Controller))
	}
}

func (r *Routing) Register(router *gin.RouterGroup) {
	var group *gin.RouterGroup
	if r.Group != "" {
		group = router.Group(r.Group)
	}

	for _, middleware := range r.Middlewares {
		if group != nil {
			group.Use(middleware.Handle)
		} else {
			router.Use(middleware.Handle)
		}
	}

	for _, route := range r.Routes {
		r.validate(route)
		route.Middlewares = append(route.Middlewares, route.Controller)
		handlers := slice.
			Map(route.Middlewares, func(middleware http.Middleware) gin.HandlerFunc {
				return middleware.Handle
			}).
			Build()

		path := route.Controller.Path()
		if len(route.Controller.Params()) > 0 {
			path = fmt.Sprintf("%s/:%s", path, strings.Join(route.Controller.Params(), "/:"))
		}

		if group != nil {
			group.Handle(route.Controller.Method(), path, handlers...)
		} else {
			router.Handle(route.Controller.Method(), path, handlers...)
		}
	}
}
