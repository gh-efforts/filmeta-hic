package httpservice

import (
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type RegisterRoutes func(gin *gin.Engine)

func GetHttpServer(registerFn RegisterRoutes, middlewares []middleware.Middleware,
	options ...http.ServerOption) *http.Server {
	e := gin.New()
	// heath check
	e.GET("/heath", func(c *gin.Context) {
		c.String(200, "success")
	})
	middlewares = append([]middleware.Middleware{
		recovery.Recovery(),
	}, middlewares...)

	e.Use(kgin.Middlewares(middlewares...))
	registerFn(e)
	httpSrv := http.NewServer(options...)
	httpSrv.HandlePrefix("/", e)
	return httpSrv
}
