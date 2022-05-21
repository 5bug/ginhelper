package templates

const MainTpl = `
package main

import (
	"context"

	"{{.Project}}/internal/{{.Name}}/api"
	"{{.Project}}/internal/{{.Name}}/conf"
	"{{.Project}}/internal/{{.Name}}/service"
	"{{.Project}}/pkg/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}
	svc := service.New()
	if err := svc.Init(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	server := rest.NewHTTPServer(conf.Conf.Http)
	server.Start(&ctx, func(gin *gin.Engine) {
		api.RegisterRouter(gin, svc)
	})
	select {}
}

`
