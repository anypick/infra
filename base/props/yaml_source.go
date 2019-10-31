package props

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type YamlSource map[string]interface{}

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
	if e = yaml.Unmarshal(data, yamlSource); e != nil {
		panic(e)
	}
	return yamlSource
}
