package router

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	noAuthUserRouter := Router.Group("/user")
	//userRouter := Router.Group("user").Use(middleware.JWT())
	//userRouterWithoutRecord := Router.Group("user")
	{
		noAuthUserRouter.POST("login-code", userApi.LoginWithCode)
		noAuthUserRouter.POST("verify", userApi.VerifyCode)
	}
}
