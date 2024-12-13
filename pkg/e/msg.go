package e

var MsgFlags = map[int]string{
	Success:       "OK",
	Error:         "Fail",
	InvalidParams: "参数错误",

	// 用户相关错误信息
	CodeLimitPerDay:    "今日验证码次数达上限",
	ErrorSaveCode:      "缓存保存验证码失败",
	ErrorIcrCount:      "验证码次数加一失败",
	CodeExpired:        "验证码过期",
	CodeError:          "验证码错误",
	ErrorGetToken:      "获取Token失败",
	ErrorDuplicateCode: "验证码已发送，请检查手机",
	CurrentUserEmpty:   "当前用户不存在",
	ValidatorPhone:     "手机号格式错误",
	ErrorDelCode:       "删除验证码失败",

	// 其他模块相关错误信息
}

// GetMsg 获取状态码对应的信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
