package utils

import (
	"fmt"
	"gin-api/internal/model"
	"github.com/gofrs/uuid/v5"
	"testing"
)

func TestJwt(t *testing.T) {
	uid, _ := uuid.NewV4()
	user := model.User{
		UUID:        uid,
		Username:    "19129211344",
		Password:    "123456789",
		NickName:    "小马",
		HeaderImg:   "https://www..baidu..com",
		AuthorityId: 0,
		Phone:       "19129211344",
		Email:       "2804179081@qq.com",
	}
	token, _, _ := LoginToken(user)

	GetUserId()
	fmt.Println(token)
}
