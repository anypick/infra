package config

import "github.com/anypick/infra/base/props/container"

//这里不使用golang自带的init方法，因为不好控制init初始化顺序，所以该用自定义的

func Init() {
	container.Add(&GinApp{Prefix: DefaultPrefix})
}
