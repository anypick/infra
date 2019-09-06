package props

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var mapData = make(map[string]interface{})

func GetMapData() map[string]interface{} {
	return mapData
}

// 将yaml文件映射成结构体
type YamlSource struct {
	Application `yaml:"application"`
	OtherConfig map[string]interface{}
}

func NewYamlSource(filePathName string) *YamlSource {
	var (
		yamlSource = new(YamlSource)
		data       []byte
		e          error
	)
	if data, e = ioutil.ReadFile(filePathName); e != nil {
		log.Fatal(e)
		return nil
	}
	if e = yaml.Unmarshal(data, &mapData); e != nil {
		panic(e)
	}
	if e = yaml.Unmarshal(data, yamlSource); e != nil {
		panic(e)
	}
	yamlSource.OtherConfig = make(map[string]interface{})
	return yamlSource
}

// 应用服务配置信息
type Application struct {
	Port string `yaml:"server.port"`
	Name string `yaml:"name"`
}
