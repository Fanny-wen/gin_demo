package main

import (
	"fmt"
	"gin_demo/initialize"
	"gin_demo/router"
)

func main() {
	initialize.InitConfig(initialize.GlobalConfigFile, &initialize.GC)
	initialize.InitLog(initialize.GinLogPath, initialize.GinLogFile)
	initialize.DB, _ = initialize.InitMysqlConnect()
	r := router.NewRouter()
	_ = r.Run(fmt.Sprintf("%s:%d", initialize.GC.App.Addr, initialize.GC.App.Port))
}
