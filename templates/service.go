package templates

const ServiceTpl = `
package service

// Service servive define
type Service struct {
	// TODO

}

// New new service
func New() (s *Service) {
	s = &Service{
		// TODO
	}
	return
}

// Init init service
func (s *Service) Init() (err error) {
	// TODO

	return
}

// Ping health check
func (s *Service) Ping() (err error) {
	// TODO

	return
}
`

//全量模板
const SvcFullTpl = `
package service

import (
	"{{.App.Project}}/internal/{{.App.Name}}/api"
	"{{.App.Project}}/pkg/rest"
)

{{- range .Requests}}
// {{.Name}} {{.Name}} service
func (s *Service) {{.Name}}(ctx *rest.Context, request *api.{{.Name}}Request) (reply *api.{{.Name}}Reply, err error) {
	// todo 

	return
}
{{- end}}
`

// 增量添加方法
const SvcFuncTpl = `
// {{.Name}} {{.Name}} service
func (s *Service) {{.Name}}(ctx *rest.Context, request *api.{{.Name}}Request) (reply *api.{{.Name}}Reply, err error) {
	// todo 

	return
}
`
