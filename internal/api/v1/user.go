package v1

import (
	systemReq "gin-api/internal/model/common/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type UserApi struct{}

func (u *UserApi) LoginWithCode(c *gin.Context) {
	var l *systemReq.LoginRequest
	if err := c.ShouldBind(&l); err == nil {
		res := UserService.LoginWithCode(c, l)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
func (u *UserApi) VerifyCode(c *gin.Context) {
	var l *systemReq.VerificationMobile
	if len(strings.TrimSpace(l.Mobile)) != 6 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "手机号不足六位"})
		return
	}
	if err := c.ShouldBind(&l); err == nil {
		res := UserService.VerificationCode(c, l.Mobile)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
