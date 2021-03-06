package templates

const RouterTpl = `
// Code generated by generator. DO NOT EDIT.
package api

import (
	"{{.App.Project}}/pkg/rest"

	"github.com/gin-gonic/gin"
)

// RegisterRouter register router
func RegisterRouter(router *gin.Engine, svc {{.App.Service}}Service) {
	router.GET("/ping", handlerPing(svc))
	group := router.Group("v1")
	{
	{{- range .Requests}}
		{{- if .Auth }}
		{{- else }}
		group.{{.Method}}("{{.Route}}", handler{{.Name}}(svc))
		{{- end }}
	{{- end }}
	}

    auth := router.Group("/v1")
	auth.Use(rest.AuthCros())
	{
    {{- range .Requests}}
		{{- if .Auth }}
    	auth.{{.Method}}("{{.Route}}", handler{{.Name}}(svc))
		{{- end}}
    {{- end}}
	}
}
`
