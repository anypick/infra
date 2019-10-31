package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func SelfMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("if exist this word, indicate this test is successful...")
		context.Next()
	}
}