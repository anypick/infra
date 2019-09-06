package base

import (
	"fmt"
	"github.com/anypick/infra/base/props/container"
)

func init() {
	container.RegisterYamContainer(&TestConfig{Prefix: "testConfig"})
}

type TestConfig struct {
	Prefix   string
	Ip       string
	Username string
	Password string
}

func (t *TestConfig) ConfigAdd(config map[interface{}]interface{}) {
	t.Ip = config["ip"].(string)
	t.Username = config["username"].(string)
	t.Password = fmt.Sprintf("%v", config["password"])
}
