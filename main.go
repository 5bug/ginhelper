package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// go install github.com/go-bindata/go-bindata/...
// go-bindata pkg/*/**

func readProject() (project, rootDir string) {
	modFile, rootDir, ok := findModFile()
	if !ok {
		return
	}
	if modFile != "" {
		data, err := ioutil.ReadFile(modFile)
		if err != nil {
			logrus.Fatal(err)
			return
		}
		str := "module\\s(.*?)\n"
		r := regexp.MustCompile(str)
		if r.Match(data) {
			arr := r.FindStringSubmatch(string(data))
			project = string(arr[1])
			return
		}
	}
	return
}

func findModFile() (modFile, rootDir string, ok bool) {
	dir, _ := os.Getwd()
	for {
		files, _ := ioutil.ReadDir(dir)
		for _, v := range files {
			if v.Name() == "go.mod" {
				modFile = path.Join(dir, v.Name())
				rootDir = dir
				ok = true
				return
			}
		}
		dir = filepath.Dir(dir)
		if dir == "/" {
			break
		}
	}
	return
}

func main() {
	appName := ""
	project, rootDir := readProject()
	cmdCreate := &cobra.Command{
		Use:   "new",
		Short: "create an app template",
		Run: func(cmd *cobra.Command, args []string) {
			rootDir, _ = os.Getwd()
			if strings.TrimSpace(project) == "" {
				files, _ := ioutil.ReadDir(rootDir)
				if len(files) != 0 {
					fmt.Println("ginhelper: 请在空目录下执行此命令")
					return
				}
				prompt := &survey.Input{
					Message: "What is project name ?",
					Help:    "Created project name.",
				}
				if err := survey.AskOne(prompt, &project); err != nil {
					logrus.Fatal(err)
					return
				}
				if project == "" {
					return
				}
				if err := InitProject(project); err != nil {
					logrus.Errorf("初始化go mod失败, 错误：%v", err)
					return
				}
			}
			if strings.TrimSpace(appName) == "" {
				prompt := &survey.Input{
					Message: "What is app name ?",
					Help:    "Created app name.",
				}
				if err := survey.AskOne(prompt, &appName); err != nil {
					logrus.Fatal(err)
					return
				}
				if appName == "" {
					return
				}
				if AppExist(rootDir, appName) {
					return
				}
			}
			if err := CreateApp(rootDir, project, appName); err != nil {
				logrus.Fatal(err)
			}
		},
	}
	cmdCreate.Flags().StringVarP(&project, "project", "p", project, "project name")
	cmdCreate.Flags().StringVarP(&appName, "appName", "a", "", "app name")

	cmdUpdate := &cobra.Command{
		Use:   "update",
		Short: "update app api and service",
		Run: func(cmd *cobra.Command, args []string) {
			if project == "" {
				fmt.Println("ginhelper: 请在工程目录下运行")
				return
			}
			if strings.TrimSpace(appName) == "" {
				prompt := &survey.Input{
					Message: "What is app name ?",
					Help:    "Created app name.",
				}
				if err := survey.AskOne(prompt, &appName); err != nil {
					logrus.Fatal(err)
					return
				}
				if project == "" {
					return
				}
			}
			if err := UpdateAPI(rootDir, project, appName); err != nil {
				logrus.Fatal(err)
			}
		},
	}
	cmdUpdate.Flags().StringVarP(&appName, "appName", "a", "", "app name")

	rootCmd := &cobra.Command{
		Use:     "ginhelper",
		Version: "0.1"}
	rootCmd.AddCommand(cmdCreate, cmdUpdate)
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
