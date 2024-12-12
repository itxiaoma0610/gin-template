package main

import (
	"fmt"
	"gin-api/config"
	"gin-api/internal/core"
	"gin-api/internal/global"
)

func main() {
	global.AIG_VP = core.Viper()
	mysql := config.Mysql{
		Host:     "123",
		Port:     123,
		User:     "123",
		Password: "123",
		DB:       "123",
	}
	fmt.Println(mysql.Dsn())
	core.RunServer()
}
