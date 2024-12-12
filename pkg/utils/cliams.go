package utils

import (
	"gin-api/internal/global"
	"gin-api/internal/model"
	systemReq "gin-api/internal/model/common/request"
	"github.com/gin-gonic/gin"
	"strings"
)

func LoginToken(user model.User) (token string, claims systemReq.CustomClaims, err error) {
	j := &JWT{SigningKey: []byte(global.AIG_CONFIG.JWT.SigningKey)} // 唯一签名
	claims = j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.GetUUID(),
		ID:          user.GetUserId(),
		NickName:    user.GetNickname(),
		Username:    user.GetUsername(),
		AuthorityId: user.GetAuthorityId(),
	})
	token, err = j.CreateToken(claims)
	if err != nil {
		return
	}
	return
}

func GetToken(c *gin.Context) string {
	// 默认Bearer Token
	tokenString := c.GetHeader("Authorization")
	token := strings.TrimPrefix(tokenString, "Bearer ")
	return token
}

func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.AIG_LOG.Error("从Gin的Context中获取从jwt解析信息失败")
	}
	return claims, err
}
