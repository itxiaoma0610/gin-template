package router

import (
	api "gin-api/internal/api/v1"
)

type RouterGroup struct {
	UserRouter
}

var (
	userApi = api.UserApi{}
)
