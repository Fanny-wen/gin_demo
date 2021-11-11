package main

import (
	"gin_demo/Router"
	_ "gin_demo/Services/logrus"
	_ "gin_demo/Services/viper"
)

func main() {
	_ = Router.Engine.Run(":8000")
}
