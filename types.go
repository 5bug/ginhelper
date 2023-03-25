package main

import (
	"strings"

	"github.com/5bug/ginhelper/templates"
)

// APIRequest api request
type APIRequest struct {
	Group   string
	Name    string
	Comment string
	Route   string
	Method  string
	Auth    bool
}

// RequestMap define RequestMap
type RequestMap map[string][]APIRequest

// Requests request list
func (m RequestMap) Requests() (list []APIRequest) {
	for _, requests := range m {
		list = append(list, requests...)
	}
	return
}

// App ...
type App struct {
	Single  bool   //是否单App
	Project string //工程名称
	Name    string //app名
	Service string //服务名(首字母大写)
	RootDir string //项目根目录
	APIDir  string //api目录
}

// RenderValue render data
type RenderValue struct {
	Single   bool
	App      App
	Requests []APIRequest
}

func (app *App) appfiles() map[string]string {
	if app.Single {
		return map[string]string{
			"main.go":                     templates.MainTpl,
			"internal/conf/conf.go":       templates.ConfTpl,
			"internal/api/api.go":         strings.ReplaceAll(templates.APITpl, "[*]", "`"),
			"internal/service/service.go": templates.ServiceTpl,
		}
	} else {
		return map[string]string{
			"cmd/%s/main.go":                 templates.MainTpl,
			"internal/%s/conf/conf.go":       templates.ConfTpl,
			"internal/%s/api/api.go":         strings.ReplaceAll(templates.APITpl, "[*]", "`"),
			"internal/%s/service/service.go": templates.ServiceTpl,
		}
	}
}
