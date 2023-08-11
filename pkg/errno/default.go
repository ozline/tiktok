// This file is be designed to define any common error so that we can use it in any service simply.

package errno

var (
	// Success
	Success = NewErrNo(SuccessCode, "Success")

	ServiceError             = NewErrNo(ServiceErrorCode, "Service is unable to start successfully")
	ServiceInternalError     = NewErrNo(ServiceErrorCode, "Service Internal Error")
	ParamError               = NewErrNo(ParamErrorCode, "Parameter error")
	AuthorizationFailedError = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")

	// User
	UserExistedError = NewErrNo(ParamErrorCode, "User existed")

	// Comment
	UnexpectedTypeError = NewErrNo(UnexpectedTypeErrorCode, "Unexpected type")
	NotImplementError   = NewErrNo(NotImplementErrorCode, "Not implement")
)
