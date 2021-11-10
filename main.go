package main

import (
	"gin_demo/Middlewares"
	"gin_demo/Router"
	"github.com/gin-gonic/gin"
	_ "github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {
	gin.ForceConsoleColor()
	Router.Engine.MaxMultipartMemory = 8 << 20 // 8 MiB

	Router.Engine.Use(gin.Recovery())
	Router.Engine.Use(Middlewares.LogMiddleware())

	f, _ := os.OpenFile("Assets/logs/gin.log", os.O_WRONLY|os.O_APPEND|os.O_APPEND, 0777)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	_ = Router.Engine.Run(":8000")
}
