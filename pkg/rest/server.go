package rest

import (
	"context"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type (
	httpRegisterFunc func(*gin.Engine)
)

// HTTPConf config define
type HttpConf struct {
	Listen  string
	Actived bool
	PProf   bool
}

// HTTPServer ...
type HTTPServer struct {
	conf *HttpConf
}

// NewHTTPServer new HTTPServer
func NewHTTPServer(conf *HttpConf) *HTTPServer {
	return &HTTPServer{
		conf: conf,
	}
}

// Start start grpc service
func (s *HTTPServer) Start(ctx *context.Context, register httpRegisterFunc) error {
	router := gin.Default()
	router.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"http://", "https://"},
			AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"},
			AllowHeaders:     []string{"Authorization", "Content-Length", "Content-Type"},
			ExposeHeaders:    []string{"Authorization"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return true
			},
			MaxAge: 12 * time.Hour,
		}),
	)
	router.Use(gin.Recovery())
	if s.conf.PProf {
		pprof.Register(router)
	}
	register(router)
	return router.Run(s.conf.Listen)
}
