package main

import (
	"github.com/PittYao/lark/pkg/xgin"
	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := xgin.NewGinServer()
	ginServer.Engine.GET("hello", hello)
	ginServer.Run(8080)
}

func hello(context *gin.Context) {
	context.SecureJSON(200, "hello 访问成功")
}
