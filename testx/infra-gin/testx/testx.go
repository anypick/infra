package testx

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra/base/props"
	"github.com/anypick/infra/base/props/container"
	"github.com/anypick/infra/testx/infra-gin"
	"github.com/anypick/infra/testx/infra-gin/config"
	"github.com/anypick/infra/testx/infra-gin/testx/middlewares"
)

func Init() {
	// 配置初始化
	config.Init()

	// 注册组件
	infra.Register(&container.YamlStarter{})

	middlewares.Init()
	infra.Register(&basegin.GinStarter{})

	infra.Register(&infra.BaseInitializerStarter{})

	// 读取配置文件
	source := props.NewYamlSource("./config.yml")
	app := infra.New(*source)
	app.Start()

}

// 用于测试
func GinX() {

}
