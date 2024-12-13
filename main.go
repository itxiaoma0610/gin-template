package main

import (
	"gin-api/core"
	"gin-api/global"
	"gin-api/initialize"
)

func main() {
	global.VP = core.Viper()
	global.LOG = core.Zap()
	global.DB = initialize.InitGorm()
	initialize.Redis()
	initialize.OtherInit()
	core.RunServer()
}
