package router

import (
	"gin-api/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	noAuthUserRouter := Router.Group("/user")
	userAuthRouter := Router.Group("/user").Use(middleware.JWTAuth())
	//userRouterWithoutRecord := Router.Group("user")
	{
		noAuthUserRouter.POST("login-code", userApi.LoginWithCode)
		noAuthUserRouter.POST("verify", userApi.VerifyCode)
	}
	{
		userAuthRouter.GET("userinfo", userApi.UserInfo)
	}
}
