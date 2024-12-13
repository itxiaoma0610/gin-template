package initialize

import (
	"gin-api/global"
	"gin-api/middleware"
	"gin-api/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	gin.SetMode(gin.DebugMode)
	Router.Use(middleware.Cors())

	systemRouter := router.RouterGroup{}
	AigRouterGroup := Router.Group("/v1")
	{
		//健康监测
		Router.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitUserRouter(AigRouterGroup)
	}
	global.ROUTERS = Router.Routes()
	return Router
}
