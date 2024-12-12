package v1

import "gin-api/internal/service"

type ApiGroup struct {
	UserApi
}

var (
	UserService = service.UserService{}
)
