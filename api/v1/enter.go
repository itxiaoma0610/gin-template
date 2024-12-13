package v1

import "gin-api/service"

type ApiGroup struct {
	UserApi
}

var (
	UserService = service.UserService{}
)
