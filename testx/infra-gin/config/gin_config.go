package config

const (
	DefaultPrefix = "application"
)

type GinApp struct {
	Prefix string
	Port   int    `yaml:"server.port"`
	Name   string `yaml:"name"`
}

func (g *GinApp) ConfigAdd(conf map[interface{}]interface{}) {
	g.Port = conf["server.port"].(int)
	g.Name = conf["name"].(string)
}
