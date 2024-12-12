package e

var MsgFlags = map[int]string{
	Success:         "OK",
	Error:           "Fail",
	InvalidParams:   "参数错误",
	CodeLimitPerDay: "今日验证码次数达上限",
	ErrorSaveCode:   "缓存保存验证码失败",
	ErrorIcrCount:   "验证码次数加一失败",
	ErrorCode:       "验证码错误",
}

// GetMsg 获取状态码对应的信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
