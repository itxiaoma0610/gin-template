package e

// 通用错误码
const (
	Success       = 200
	Error         = 500
	InvalidParams = 400
)

// 用户相关错误码
const (
	CodeLimitPerDay    = 10001
	ErrorSaveCode      = 10002
	ErrorIcrCount      = 10003
	CodeExpired        = 10004
	ErrorGetToken      = 10005
	ErrorDuplicateCode = 10006
	CurrentUserEmpty   = 10007
	ValidatorPhone     = 10008
	ErrorDelCode       = 10009
	CodeError          = 10010
)
