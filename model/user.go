package model

import (
	"gin-api/global"
	"github.com/gofrs/uuid/v5"
)

type Login interface {
	GetUsername() string
	GetNickname() string
	GetUUID() uuid.UUID
	GetUserId() uint
	GetAuthorityId() uint
	GetUserInfo() any
}

// 确保User实现了login的接口
var _ Login = new(User)

type User struct {
	global.AIC_MODEL
	UUID        uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`                                                     // 用户UUID
	Username    string    `json:"userName" gorm:"index;comment:用户登录名"`                                                  // 用户登录名
	Password    string    `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName    string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	HeaderImg   string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	AuthorityId uint      `json:"authorityId" gorm:"default:0;comment:用户角色ID"`                                          // 用户角色ID
	Phone       string    `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	Email       string    `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
}

func (User) TableName() string {
	return "sys_users"
}

func (s *User) GetUsername() string {
	return s.Username
}

func (s *User) GetNickname() string {
	return s.NickName
}

func (s *User) GetUUID() uuid.UUID {
	return s.UUID
}

func (s *User) GetUserId() uint {
	return s.ID
}

func (s *User) GetAuthorityId() uint {
	return s.AuthorityId
}

func (s *User) GetUserInfo() any {
	return *s
}
