package conf

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

// Load load config
func Load(c interface{}) (err error) {
	defer func() {
		data, _ := json.Marshal(c)
		logrus.Infof("config data:%s", string(data))
	}()
	var fileName string
	flag.StringVar(&fileName, "config", "", "config fileName")
	flag.Parse()
	if fileName == "" {
		logrus.Info("use default config")
		return
	}
	logrus.Infof("config file[%s]", fileName)
	data, err := ioutil.ReadFile(fileName)
	if err == nil {
		if err = yaml.Unmarshal(data, c); err != nil {
			logrus.Error("config Unmarshal err:", err)
			return
		}
	}
	return
}
