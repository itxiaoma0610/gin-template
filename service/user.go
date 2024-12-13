package service

import (
	"context"
	"errors"
	"fmt"
	"gin-api/global"
	"gin-api/model"
	"gin-api/model/common/request"
	"gin-api/model/common/serializer"
	"gin-api/pkg/e"
	"gin-api/pkg/utils"
	"gin-api/repository"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"time"
)

const (
	prefixVerificationCount = "biz#verification#count#"
	prefixVerification      = "biz#verification#"
	verificationLimitPerDay = 10
	expireActivation        = 10 * time.Minute
	expireActivationCount   = 24 * time.Hour
)

type UserService struct{}

var UserServiceApp = new(UserService)

// LoginWithCode @Description: 短信验证码登录
// @receiver u
// @param ctx
// @param username
// @return serializer.Response
func (l *UserService) LoginWithCode(ctx context.Context, req *request.LoginRequest) serializer.Response {
	code := e.Success
	if err := l.checkVerificationCode(ctx, req.Mobile, req.Code); err != 0 {
		code = err
		return serializer.Response{
			Status: code,
			Error:  e.GetMsg(code),
		}
	}
	mobile := utils.Encrypt.AesEncoding(req.Mobile)
	user, err := repository.UserRepo.GetUserByMobile(mobile)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.Response{
			Status: e.CurrentUserEmpty,
			Error:  e.GetMsg(e.CurrentUserEmpty),
		}
	}
	err = l.delVerificationCode(ctx, req.Mobile)
	if err != nil {
		code = e.ErrorDelCode
		return serializer.Response{
			Status: code,
			Error:  e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   *user,
		Msg:    e.GetMsg(code),
	}

}

// TokenNext @Description: 登录成功后签发Token
// @param user
// @return 登录成功签发Token信息
func (l *UserService) TokenNext(ctx context.Context, user *model.User) serializer.Response {
	token, claims, err := utils.LoginToken(user)
	if err != nil {
		global.LOG.Error("获取Token失败", zap.Error(err))
		return serializer.Response{
			Status: e.ErrorGetToken,
			Error:  e.GetMsg(e.ErrorGetToken),
		}
	}
	return serializer.Response{
		Status: e.Success,
		Data:   serializer.BuildLogin(user, token, claims.ExpiresAt.String(), "Bearer"),
		Msg:    e.GetMsg(e.Success),
	}
}

// VerificationCode @Description: 处理发送短信逻辑
// @param mobile 电话号码
// @return serializer.Response
func (l *UserService) VerificationCode(ctx context.Context, mobile string) serializer.Response {
	code := e.Success
	// 验证验证码发送次数，不超过五次
	count, _ := l.getVerificationCount(ctx, mobile)
	if count >= verificationLimitPerDay {
		code = e.CodeLimitPerDay
		return serializer.Response{
			Status: code,
			Error:  e.GetMsg(code),
		}
	}
	verifycode, err := l.getActivationCache(mobile)
	if errors.Is(err, redis.Nil) {
		global.LOG.Info("getActivationCache mobile error :", zap.Error(err))
	}
	if len(verifycode) == 0 {
		verifycode = utils.RandomNumeric(6)
		fmt.Println(verifycode)
	} else {
		code = e.ErrorDuplicateCode
		return serializer.Response{
			Status: code,
			Error:  e.GetMsg(code),
		}
	}
	// todo 发送验证码

	//验证码保存到缓存中
	err = l.saveActivationCache(mobile, verifycode)
	if err != nil {
		code = e.ErrorSaveCode
		return serializer.Response{
			Status: code,
			Error:  e.GetMsg(code),
		}
	}
	// 验证码次数加一
	err = l.icrActivationCount(mobile)
	if err != nil {
		code = e.ErrorIcrCount
		return serializer.Response{
			Status: code,
			Error:  e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
func (l *UserService) getVerificationCount(ctx context.Context, mobile string) (int, error) {
	key := prefixVerificationCount + mobile
	val, err := global.REDIS.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return 0, nil
	}
	if len(val) == 0 {
		return 0, nil
	}
	return strconv.Atoi(val)
}
func (l *UserService) getActivationCache(mobile string) (string, error) {
	key := prefixVerification + mobile
	return global.REDIS.Get(context.Background(), key).Result()
}
func (l *UserService) saveActivationCache(mobile, code string) error {
	key := prefixVerification + mobile
	return global.REDIS.Set(context.Background(), key, code, expireActivation).Err()
}
func (l *UserService) icrActivationCount(mobile string) error {
	key := prefixVerificationCount + mobile
	if err := global.REDIS.Incr(context.Background(), key).Err(); err != nil {
		return err
	}
	return global.REDIS.Expire(context.Background(), key, expireActivationCount).Err()
}
func (l *UserService) checkVerificationCode(ctx context.Context, mobile, code string) int {
	// 检查验证码
	cacheCode, err := l.getActivationCache(mobile)
	if errors.Is(err, redis.Nil) {
		return e.CodeExpired
	}
	if cacheCode != code {
		return e.CodeError
	}
	return 0
}

func (l *UserService) delVerificationCode(ctx context.Context, mobile string) error {
	// 登录成功后删除验证码
	key := prefixVerification + mobile
	return global.REDIS.Del(ctx, key).Err()
}
