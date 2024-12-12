package service

import (
	"context"
	"errors"
	"gin-api/internal/global"
	"gin-api/internal/model/common/request"
	"gin-api/internal/model/common/serializer"
	"gin-api/pkg/e"
	"gin-api/pkg/utils"
	"github.com/redis/go-redis/v9"
	"strconv"
)

const (
	prefixVerificationCount = "biz#verification#count#%s"
	prefixVerification      = "biz#verification#%s"
	verificationLimitPerDay = 10
	expireActivation        = 60 * 30
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
	err := l.checkVerificationCode(ctx, req.Mobile, req.Code)
	if err != nil {
		code = e.ErrorCode
		return serializer.Response{
			Status: 0,
			Msg:    err.Error(),
		}
	}
	//mmobile := util.Encrypt.AesEncoding(req.Mobile)
	return serializer.Response{
		Status: code,
		Data:   nil,
		Msg:    "",
		Error:  "",
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
		global.AIG_LOG.Info("getActivationCache mobile: %s error :%v")
	}
	if len(verifycode) == 0 {
		verifycode = utils.RandomNumeric(6)
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
	val, err := global.AIG_REDIS.Get(ctx, key).Result()
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
	return global.AIG_REDIS.Get(context.Background(), key).Result()
}
func (l *UserService) saveActivationCache(mobile, code string) error {
	key := prefixVerification + mobile
	return global.AIG_REDIS.Set(context.Background(), key, code, expireActivation).Err()
}
func (l *UserService) icrActivationCount(mobile string) error {
	key := prefixVerificationCount + mobile
	if err := global.AIG_REDIS.Incr(context.Background(), key).Err(); err != nil {
		return err
	}
	return global.AIG_REDIS.Expire(context.Background(), key, expireActivation).Err()
}
func (l *UserService) checkVerificationCode(ctx context.Context, mobile, code string) error {
	// 检查验证码
	cacheCode, err := l.getActivationCache(mobile)
	if err != nil {
		return err
	}
	if cacheCode == "" {
		return errors.New("verification code expired")
	}
	if cacheCode != code {
		return errors.New("verification code failed")
	}
	return nil
}
