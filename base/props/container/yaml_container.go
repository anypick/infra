package container

import (
	"github.com/anypick/infra"
	"reflect"
)

// 配置类的父类
type YamlConfig interface {
	// 配置添加
	ConfigAdd(map[interface{}]interface{})
}

var yamlContainer []YamlConfig = make([]YamlConfig, 0)

// 配置添加
func Add(yaml YamlConfig) {
	yamlContainer = append(yamlContainer, yaml)
}

// 装载配置
func Execute(ctx infra.StarterContext) {
	mapData := ctx.Yaml()
	for _, yamlConfig := range yamlContainer {
		prefix := reflect.ValueOf(yamlConfig).Elem().FieldByName("Prefix").String()
		config := mapData[prefix].(map[interface{}]interface{})
		ctx.Yaml()[prefix] = yamlConfig
		yamlConfig.ConfigAdd(config)
	}
}

type YamlStarter struct {
	infra.BaseStarter
}

func (y *YamlStarter) Init(ctx infra.StarterContext) {
	Execute(ctx)
}
