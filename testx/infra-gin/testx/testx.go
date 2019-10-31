package testx

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra-gin"
	"github.com/anypick/infra-gin/testx/middlewares"
	"github.com/anypick/infra-logrus"
	"github.com/anypick/infra/base/props"
	"github.com/anypick/infra/base/props/container"
)

func Init() {
	// 注册组件
	infra.Register(&container.YamlStarter{})

	infra.Register(&baselog.LogrusStarter{})

	middlewares.Init()
	infra.Register(&basegin.GinStarter{})

	// 读取配置文件
	source := props.NewYamlSource("./config.yml")
	app := infra.New(*source)
	app.Start()

}

// 用于测试
func GinX() {

}
