package utils

import (
	"fmt"
	"gin-api/internal/model"
	"github.com/google/uuid"
	"testing"
)

func TestJwt(t *testing.T) {
	user := model.User{
		UUID:        uuid.New(),
		Username:    "19129211344",
		Password:    "123456789",
		NickName:    "小马",
		HeaderImg:   "https://www..baidu..com",
		AuthorityId: 0,
		Phone:       "19129211344",
		Email:       "2804179081@qq.com",
	}
	token, _, _ := LoginToken(user)

	fmt.Println(token)
}
