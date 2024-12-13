package serializer

import "gin-api/model"

type LoginVo struct {
	UserID    uint   `json:"user_id"`
	NickName  string `json:"nick_name"`
	Avatar    string `json:"avatar"`
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
	ExpireAt  string `json:"expire_at"`
}

func BuildLogin(user *model.User, token, expireAt, tokenType string) *LoginVo {
	return &LoginVo{
		UserID:    user.ID,
		NickName:  user.NickName,
		Avatar:    user.HeaderImg,
		Token:     token,
		TokenType: tokenType,
		ExpireAt:  expireAt,
	}
}
