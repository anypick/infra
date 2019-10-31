package middlewares

import (
	"github.com/anypick/infra-gin/helper"
	"github.com/gin-gonic/gin"
)

func Init() {
	helper.AddMiddleWare(helper.LogrusMiddle())
	//helper.AddMiddleWare(gin.Logger())
	helper.AddMiddleWare(gin.Recovery())
	helper.AddMiddleWare(SelfMiddleware())
	helper.AddMiddleWare()
}