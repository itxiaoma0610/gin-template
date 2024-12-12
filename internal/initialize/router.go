package initialize

import (
	"gin-api/internal/global"
	"gin-api/internal/middleware"
	"gin-api/internal/router"
	"github.com/gin-gonic/gin"
	"net/http"
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
	global.AIG_ROUTERS = Router.Routes()
	return Router
}
