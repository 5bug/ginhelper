package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
	"text/template"

	"github.com/5bug/ginhelper/templates"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// CreateApp create an app
func CreateApp(rootDir, project, appName string, single bool) (err error) {
	dir := fmt.Sprintf("internal/%s/api", strings.ToLower(appName))
	if single {
		dir = "internal/api"
	}
	app := App{
		Single:  single,
		Project: project,
		Name:    strings.ToLower(appName),
		Service: cases.Title(language.Und, cases.NoLower).String(appName),
		RootDir: rootDir,
		APIDir:  path.Join(rootDir, dir),
	}
	for k, v := range app.appfiles() {
		var destFile string
		if single {
			destFile = path.Join(app.RootDir, k)
		} else {
			destFile = path.Join(app.RootDir, fmt.Sprintf(k, app.Name))
		}
		if ok, _ := PathExists(destFile); !ok {
			if err = writeOneFile(app, v, destFile); err != nil {
				return
			}
		} else {
			logrus.Warnf("file[%s] is exists", destFile)
		}
	}
	if ok, _ := PathExists("pkg"); !ok {
		if err = RestoreAssets(app.RootDir, "pkg"); err != nil {
			return
		}
	}
	if err = UpdateAPI(rootDir, project, appName, single); err != nil {
		return
	}
	return
}

// AppExist app exist
func AppExist(rootDir, appName string) (ok bool) {
	addDir := path.Join(rootDir, fmt.Sprintf("cmd/%s", strings.ToLower(appName)))
	ok, _ = PathExists(addDir)
	return
}

// UpdateAPI  update api
func UpdateAPI(rootDir, project, appName string, single bool) (err error) {
	dir := fmt.Sprintf("internal/%s", strings.ToLower(appName))
	// 判断是否符合single模式,若存在internal/{app}目录则认为是多app模式
	if ok, _ := PathExists(dir); single == ok {
		err = errors.New("App模式不相符，请确认-s参数是否正确")
		return
	}
	if single {
		dir = "internal"
	}
	app := App{
		Single:  single,
		Project: project,
		Name:    strings.ToLower(appName),
		Service: strings.ReplaceAll(cases.Title(language.Und, cases.NoLower).String(appName), "-", ""),
		RootDir: rootDir,
		APIDir:  path.Join(rootDir, dir+"/api"),
	}
	requestMap, err := readRequests(app.APIDir)
	if err != nil {
		return
	}
	value := RenderValue{
		App:      app,
		Single:   single,
		Requests: requestMap.Requests(),
	}
	// 创建handlers.go
	if err = writeOneFile(value, templates.HandlersTpl, path.Join(app.APIDir, "handlers.go")); err != nil {
		return
	}
	// 创建router.go
	if err = writeOneFile(value, templates.RouterTpl, path.Join(app.APIDir, "router.go")); err != nil {
		return
	}
	// 创建service文件
	for group, requests := range requestMap {
		value = RenderValue{
			Single:   single,
			App:      app,
			Requests: requests,
		}
		destFile := path.Join(app.RootDir, fmt.Sprintf("/%s/service/%s.go", dir, group))
		if ok, _ := PathExists(destFile); !ok {
			if err = writeOneFile(value, templates.SvcFullTpl, destFile); err != nil {
				return
			}
		} else {
			if err = patchService(&value, templates.SvcFuncTpl, destFile); err != nil {
				return
			}
		}
	}
	return nil
}

// InitProject init project
func InitProject(project string) error {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("go mod init %s", project))
	_, err := cmd.CombinedOutput()
	return err
}

func readRequests(apiPath string) (requestMap RequestMap, err error) {
	var data []byte
	files, err := ioutil.ReadDir(apiPath)
	if err != nil {
		return
	}
	requestMap = make(RequestMap)
	for _, file := range files {
		if strings.EqualFold(file.Name(), "router.go") || strings.EqualFold(file.Name(), "handlers.go") {
			continue
		}
		apiFile := path.Join(apiPath, file.Name())
		if path.Ext(apiFile) != ".go" {
			continue
		}
		group := strings.ReplaceAll(file.Name(), ".go", "")
		data, err = ioutil.ReadFile(apiFile)
		if err != nil {
			logrus.Infof("file[%s]read file error:%v", file.Name(), err)
			continue
		}
		str := "//\\s(.*?)Request\\s(.*?)\n//\\s@Router\\s(.*?)\\s\\[(.*?)\\]\n//\\s@Auth\\s(.*?)\n"
		r := regexp.MustCompile(str)
		params := r.FindAllStringSubmatch(string(data), -1)
		if len(params) <= 0 {
			logrus.Infof("file[%s]cannot match method", file.Name())
			continue
		}
		requests, ok := requestMap[group]
		if !ok {
			requests = []APIRequest{}
		}
		for _, p := range params {
			requests = append(requests, APIRequest{
				Group:   group,
				Name:    p[1],
				Comment: p[2],
				Route:   p[3],
				Method:  strings.ToUpper(p[4]),
				Auth:    ParseBool(p[5]),
			})
		}
		requestMap[group] = requests
	}
	return
}

func render(value interface{}, tplData string) ([]byte, error) {
	tmpl, err := template.New("value").Parse(tplData)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, "value", value)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func writeOneFile(value interface{}, tplData string, destFile string) error {
	data, err := render(value, tplData)
	if err != nil {
		return errors.Wrap(err, destFile)
	}
	if data, err = format.Source(data); err != nil {
		return err
	}
	if err = os.MkdirAll(path.Dir(destFile), os.ModePerm); err != nil {
		return err
	}
	if err = ioutil.WriteFile(destFile, data, 0666); err != nil {
		return err
	}
	return nil
}

func patchService(value *RenderValue, tplData string, destFile string) error {
	oldData, err := ioutil.ReadFile(destFile)
	if err != nil {
		return err
	}
	bs := [][]byte{oldData}
	for _, request := range value.Requests {
		str := fmt.Sprintf("func \\(s \\*Service\\) %s\\(ctx \\*rest\\.Context, request \\*api\\.%sRequest\\) \\(reply \\*api.%sReply, err error\\)",
			request.Name, request.Name, request.Name)
		r := regexp.MustCompile(str)
		if !r.Match(oldData) {
			data, err := render(request, tplData)
			if err != nil {
				return err
			}
			bs = append(bs, data)
		}
	}
	newData := bytes.Join(bs, []byte("\n\n"))
	if newData, err = format.Source(newData); err != nil {
		return err
	}
	if err = ioutil.WriteFile(destFile, newData, 0666); err != nil {
		return err
	}
	return nil
}
