package basegin

import (
	"fmt"
	"github.com/anypick/infra"
	"github.com/anypick/infra/testx/infra-gin/config"
	"github.com/sirupsen/logrus"

	//"github.com/anypick/infra-gin/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ginEngine *gin.Engine

// 对外暴露
func Gin() *gin.Engine {
	return ginEngine
}

type GinStarter struct {
	infra.BaseStarter
}

func (g *GinStarter) Init(ctx infra.StarterContext) {
	ginEngine = initGinApp()
	ginEngine.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})
}

func (g *GinStarter) Start(ctx infra.StarterContext) {
	conf := ctx.Yaml()[config.DefaultPrefix].(*config.GinApp)
	var (
		engine *gin.Engine
		e      error
	)
	engine = Gin()
	routes := engine.Routes()
	for _, info := range routes {
		logrus.Infof("API: %s %s %s", info.Method, info.Path, info.Handler)
	}
	logrus.Infof("gin start with port %d", conf.Port)
	if e = engine.Run(fmt.Sprintf(":%d", conf.Port)); e != nil {
		panic(e)
	}

}

// web服务是阻塞的
func (g *GinStarter) StartBlocking() bool {
	return true
}

// 初始化gin
func initGinApp() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	return app
}
