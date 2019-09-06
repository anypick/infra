package main

import (
	"fmt"
	"github.com/anypick/infra"
	"github.com/anypick/infra/base/props"
	"github.com/anypick/infra/base/props/container"
	"github.com/anypick/infra/utils/base"
	_ "github.com/anypick/infra/utils/base"
)

func main() {

	infra.Register(&container.YamlStarter{})

	yaml := props.NewYamlSource("testx/application.yml")

	application := infra.New(*yaml)
	application.Start()
	fmt.Println(yaml.OtherConfig["testConfig"].(*base.TestConfig))
	fmt.Println(yaml.Application)
}