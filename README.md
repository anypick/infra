# 1. infra功能介绍

> infra是一个Golang实现的极其简单的依赖注入框架，使用Yaml作为配置文件。支持`组件配置`,`组件扩展`
>
> 

# 2. infra文件介绍

**【starter.go】**

> 1. 定义接口`typ Starter interface{...}`，所有的自定义的Starter都要实现这个接口，这里已经有一个自定义实现了`type BaseStarter struct{}`,当你定义一个属于自己的Starter的时候可以直接组合该结构体，像下面这样,这样你就不需要去实现所有的方法了
>
> ```go
> type MyStarter struct {
>   infra.BaseStarter
> }
> ```
>
> 2. 定义结构体`type starterRegister struct{...}`,用于存储你自定义的`Starter`,这些Starter有的是阻塞的（如`Gin`,`Grpc`）,阻塞启动优先级要比非阻塞优先级低，当有多个阻塞Starter，只有最后一个可以阻塞整个程序，其他的应该使用协程去运行。

**【boot.go】**

> 1. 定义结构体`type BootApplication struct{...}`，对于`Starter`的启动，需要经历`Init`,` Start`,` Setup`三个阶段,所以这里使用了`模板设计模式`

**【initializer.go】**

> 1. 定义接口`type Initializer interface {...}`,对于业务代码中实例（例如Controller层），不希望使用`typ Starter interface{...}`这个接口，我们可以实现这个接口。这里还有`type InitializerRegister struct{}`结构体，用于存储业务代码中实例。
> 2. `type BaseInitializerStarter struct{}` 这是一个`Starter`,负责统一调用业务实例

**【utils/】**:定义一些工具类。

**【base/】**:定义`YamlStarter`

> 对于这个YamlStarter使用`装饰者设计模式`支持用户扩展自定义配置。

# 3. infra使用

```shell
$ go get github.com/anypick/infra
```

## 1. 自定义Starter

具体代码请参照github源码：[infra-gin](https://github.com/anypick/infra-gin)

> 这里以整合Gin框架为例，关键点：
>
> **【gin.go】**
>
> 1. 组合`BaseStarter`
>
> ```go
> type GinStarter struct {
> 	infra.BaseStarter
> }
> ```
>
> 2. 对外暴露gin实例（单例）
>
> ```go
> var ginEngine *gin.Engine
> 
> // 对外暴露
> func Gin() *gin.Engine {
> 	return ginEngine
> }
> ```

## 2. 自定义Starter

具体代码参照github源码：[infra-logrus](https://github.com/anypick/infra-logrus)

这个自定义Stater整合logrus框架为例

**【logrus.go】**

> 和`gin.go`一样，目的都是组合`infra.BaseStarter`实现`Starter接口`

这里主要关注点：**【logrus_config.go】**

定义一个配置log的配置结构体,`Prefix`，为必须字段，主要作用是用来获取改结构体。这个结构体实现了infra项目中`type YamlConfig interface{}`接口（在目录base/props/container/下）。

```go
// 日志配置
type LogConfig struct {
	Prefix       string
	Level        string `yaml:"level"`
	LogFileName  string `yaml:"logFileName"`
	FilePath     string `yaml:"filePath"`
	MaxAge       int    `yaml:"maxAge"`
	RotationTime int    `yaml:"rotationTime"`
}

func (l *LogConfig) ConfigAdd(config map[interface{}]interface{}) {
	l.Level = fmt.Sprintf("%v", config["level"])
	l.LogFileName = fmt.Sprintf("%v", config["logFileName"])
	l.FilePath = fmt.Sprintf("%v", config["filePath"])
	l.MaxAge = config["maxAge"].(int)
	l.RotationTime = config["rotationTime"].(int)
}
```

使用init方法，将该结构体放入`YamlContainer`中

```go
func init() {
	container.RegisterYamContainer(&LogConfig{Prefix: YamlPrefix})
}
```

# 4.整合

项目地址：[infra-example](https://github.com/anypick/infra-example)

```shell
$ go get github.com/anypick/infra
$ go get github.com/anypick/infra-logrus
$ go get github.com/anypick/infra-gin
```

  前面我们已经定义好了两个starter,现在来新建一个项目`infra-example`,对于项目的目录我喜欢这样：

```
├── README.md
├── app.go
├── brun
│   └── main.go
├── go.mod
├── go.sum
├── resources
│   └── application.yml
└── src
    └── GinExample.go
```

`application.yml`:为配置文件目录

`main.go`:程序启动类

`app.go`:负责注册需要的实例

`src`：定义业务代码

**【app.go】**

```go
package example

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra-gin"
	"github.com/anypick/infra-logrus"
	"github.com/anypick/infra/base/props/container"
)

func init() {
  // YamlStarter,是必须要注册的Starter
	infra.Register(&container.YamlStarter{})
	infra.Register(&baselog.LogrusStarter{})
	infra.Register(&basegin.GinStarter{})

  // BaseInitializerStarter也是必须要注册的Starter
	infra.Register(&infra.BaseInitializerStarter{})
}
```

**【GinExample.go】**

> 这是一个Controller,需要实现接口
>
> ```go
> // 用于业务的代码的注入，例如Dao层，Service层，Controller层
> type Initializer interface {
> 	Init()
> }
> ```

**【main.go】**

> GinExample.go中有一个init函数，需要在这里引入

```go
import (
  // 引入app.go init函数
	_ "example"
  // 引入GinExample.go init函数
	_ "example/src"
	"flag"
	"fmt"
	"github.com/anypick/infra"
	"github.com/anypick/infra/base/props"
	"github.com/anypick/infra/utils/common"
)

func main() {
	// 生成网站：http://patorjk.com/software/taag
	banner := `...`
	fmt.Println(banner)
	profile := flag.String("profile", "", "环境信息")
	flag.Parse()
	resource := ""
	if common.StrIsBlank(*profile) {
		resource = "resources/application.yml"
	} else {
		resource = fmt.Sprintf("resources/application-%s.yml", *profile)
	}
	yamlConf := props.NewYamlSource(resource)
	application := infra.New(*yamlConf)
	application.Start()
}
```

