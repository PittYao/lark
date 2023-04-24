package xgin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GinServer struct {
	Engine *gin.Engine
}

func NewGinServer() *GinServer {
	var (
		engine *gin.Engine
	)
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	engine.Use(gin.Recovery())
	return &GinServer{Engine: engine}
}

func (g *GinServer) Run(port int) error {
	var (
		addr string
		err  error
	)
	addr = ":" + strconv.Itoa(port)
	err = g.Engine.Run(addr)
	if err != nil {
		fmt.Println("GinServer Start Failed:", err.Error())
	}
	return err
}

func (g *GinServer) Use(m ...gin.HandlerFunc) gin.IRoutes {
	return g.Engine.Use(m...)
}
