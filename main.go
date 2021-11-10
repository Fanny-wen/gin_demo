package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sirupsen/logrus"
	"io"
	"os"
	//"github.com/Fanny-wen/gin_demo/Router"
)



func init() {
	gin.ForceConsoleColor()
	//
	f, _ := os.OpenFile("Assets/logs/gin.log", os.O_WRONLY|os.O_APPEND|os.O_APPEND, 0777)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {

}
