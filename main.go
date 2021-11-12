package main

import (
	"gin_demo/initialize"
	"gin_demo/router"
)

func init() {
	initialize.InitConfig(initialize.GlobalConfigFile, &initialize.GC)
	initialize.InitLog(initialize.GinLogPath, initialize.GinLogFile)
	initialize.DB, _ = initialize.InitMysqlConnect()
	router.InitRouter()
}

func main() {
	_ = router.Engine.Run(":8000")
}
