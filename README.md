ginhelper是用于gin框架快速开发的辅助工具，支持monorepo方式，使用方法如下：
1. 安装ginhelper
```bash
>go install github.com/5bug/ginhelper@latest
>ginhelper -h
Usage:
  ginhelper [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  new         create an app template
  update      update app api and service

Flags:
  -h, --help      help for ginhelper
  -v, --version   version for ginhelper

Use "ginhelper [command] --help" for more information about a command.
```
2. 初始化工程
```bash
➜  ~/codes/projects/test ginhelper new
? What is project name ? github.com/5bug/test
? What is app name ? demo1
```
创建好的目录结构如下：
```bash
.
├── cmd
│   └── demo1
│       └── main.go
├── go.mod
├── internal
│   └── demo1
│       ├── api
│       │   ├── api.go
│       │   ├── handlers.go
│       │   └── router.go
│       ├── conf
│       │   └── conf.go
│       └── service
│           ├── api.go
│           └── service.go
└── pkg
    ├── conf
    │   └── conf.go
    ├── rest
    │   ├── jwt.go
    │   ├── page.go
    │   ├── rest.go
    │   └── server.go
    └── storage
        └── storage.go

11 directories, 14 files
```

3. 定义api文件
这里api定义使用纯go代码实现，通过注解来定义路由、请求方法、是否需要认证，具体注解说明如下：
- @Router: 路由定义 /xxx [请求方法(get post put delete head)]
- @Auth: 是否需要认证 true/false

例如在internal/demo1/api/目录下定义api文件user.go，内容如下：
```go
package api

// UserRequest user request
// @Router /user [get]
// @Auth false
type UserRequest struct {
	ID int `json:"id" binding:"required"`
}

// UserReply user reply
type UserReply struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Sex   int    `json:"sex"`
	Email string `json:"email"`
}
```
3. 根据API文件更新app服务
```bash
ginhelper update -a demo1
```
4. 在生成的internal/demo1/service/xxx.go里实现自己的业务逻辑即可

## RoadMap
完善中...
