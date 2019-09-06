package container

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra/base/props"
	"reflect"
)

type YamlConfig interface {
	// 配置添加
	ConfigAdd(map[interface{}]interface{})
}


type YamlContainer struct {
	YamlConfigs []YamlConfig
}

func (y *YamlContainer) Add(yaml YamlConfig) {
	y.YamlConfigs = append(y.YamlConfigs, yaml)
}

var yamlContainers = new(YamlContainer)

func RegisterYamContainer(yaml YamlConfig) {
	yamlContainers.Add(yaml)
}

func GetYamlConfigs() []YamlConfig {
	return yamlContainers.YamlConfigs
}

type YamlStarter struct {
	infra.BaseStarter
}

func (y *YamlStarter) Init(ctx infra.StarterContext) {
	mapData := props.GetMapData()
	for _, register := range GetYamlConfigs() {
		prefix := reflect.ValueOf(register).Elem().FieldByName("Prefix").String()
		config := mapData[prefix].(map[interface{}]interface{})
		ctx.Yaml().OtherConfig[prefix] = register
		register.ConfigAdd(config)
	}
}





