package v1

import (
	"gin-api/model"
	systemReq "gin-api/model/common/request"
	"gin-api/pkg/e"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (u *UserApi) LoginWithCode(c *gin.Context) {
	var l *systemReq.LoginRequest
	if err := c.ShouldBind(&l); err == nil {
		res := UserService.LoginWithCode(c, l)
		if res.Status != 200 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": res.Status,
				"error":  e.GetMsg(res.Status),
			})
			return
		}
		user := res.Data.(model.User) // 签发TokenNext,多种登录方式复用
		res = UserService.TokenNext(c, &user)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
func (u *UserApi) VerifyCode(c *gin.Context) {
	var l systemReq.VerificationMobile
	if err := c.ShouldBind(&l); err == nil {
		if len(strings.TrimSpace(l.Mobile)) != 11 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": e.ValidatorPhone,
				"error":  e.GetMsg(e.ValidatorPhone)})
			return
		}
		res := UserService.VerificationCode(c, l.Mobile)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func (u *UserApi) UserInfo(c *gin.Context) {
	var l systemReq.VerificationMobile
	if err := c.ShouldBind(&l); err == nil {
		if len(strings.TrimSpace(l.Mobile)) != 11 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": e.ValidatorPhone,
				"error":  e.GetMsg(e.ValidatorPhone)})
			return
		}
		res := UserService.VerificationCode(c, l.Mobile)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
