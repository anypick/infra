package infra

import (
	"github.com/anypick/infra/base/props"
	"github.com/sirupsen/logrus"
	"reflect"
)

// 负责Starter各个阶段方法的调用
type BootApplication struct {
	conf           props.YamlSource
	starterContext StarterContext
}

var (
	yamlProps props.YamlSource
)

func GetYamlProps() props.YamlSource {
	return yamlProps
}

func New(conf props.YamlSource) *BootApplication {
	application := &BootApplication{conf, StarterContext{}}
	application.starterContext[defaultProps] = conf
	yamlProps = conf
	return application
}

func (b *BootApplication) Start() {
	//1. 初始化starter
	b.init()
	//2. 安装starter
	b.setup()
	//3. 启动starter
	b.start()
}

func (b *BootApplication) init() {
	starters := StarterRegister.AllStarters()
	for _, starter := range starters {
		starter.Init(b.starterContext)
	}
	logrus.Info("Application init finished...")
}

func (b *BootApplication) setup() {
	starters := StarterRegister.AllStarters()
	for _, starter := range starters {
		starter.Setup(b.starterContext)
	}
	logrus.Info("Application setup finished...")
}

func (b *BootApplication) start() {
	starters := StarterRegister.AllStarters()
	for index, starter := range starters {
		typ := reflect.TypeOf(starter)
		logrus.Infof("Starting:%s", typ.String())
		if starter.StartBlocking() {
			if index+1 == len(StarterRegister.AllStarters()) {
				starter.Start(b.starterContext)
			} else {
				go starter.Start(b.starterContext)
			}
		} else {
			starter.Start(b.starterContext)
		}
	}
}
