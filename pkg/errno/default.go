// This file is be designed to define any common error so that we can use it in any service simply.

package errno

var (
	Success                  = NewErrNo(SuccessCode, "Success")
	ServiceError             = NewErrNo(ServiceErrorCode, "Service is unable to start successfully")
	ParamError               = NewErrNo(ParamErrorCode, "Parameter error")
	AuthorizationFailedError = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
)
