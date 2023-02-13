// This file is designed to define any error code
package errno

const (
	SuccessCode                = 10000 // 成功返回
	ServiceErrorCode           = 10001 // 未知微服务错误
	ParamErrorCode             = 10002 // 参数错误
	AuthorizationFailedErrCode = 10003 // 鉴权失败
)
