package templates

const ConfTpl = `
package conf

import (
	c "{{.Project}}/pkg/conf"
	"{{.Project}}/pkg/rest"
	"{{.Project}}/pkg/storage"
)

// Config ...
type Config struct {
	DB        *storage.DataBase
	Redis     *storage.RedisConf
	Http      *rest.HttpConf
	//TODO: 自定义的配置
	
}

var (
	// Conf default config
	Conf = Config{
		Redis: &storage.RedisConf{
			NetWork:     "tcp",
			HostAddr:    "localhost:6480",
			MaxIdle:     16,
			IdleTimeout: 240,
		},
		DB: &storage.DataBase{
			DriverName: "mysql",
			Host:       "127.0.0.1",
			Port:       3306,
			User:       "root",
			PassWord:   "123456",
			DBName:     "1024Helper",
		},
		Http: &rest.HttpConf{
			Listen:  ":8080",
			Actived: true,
		},
	}
)

// Init init config
func Init() error {
	return c.Load(&Conf)
}
`
