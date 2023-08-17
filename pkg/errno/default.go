// This file is be designed to define any common error so that we can use it in any service simply.

package errno

var (
	// Success
	Success = NewErrNo(SuccessCode, "Success")

	ServiceError             = NewErrNo(ServiceErrorCode, "service is unable to start successfully")
	ServiceInternalError     = NewErrNo(ServiceErrorCode, "service internal error")
	ParamError               = NewErrNo(ParamErrorCode, "parameter error")
	AuthorizationFailedError = NewErrNo(AuthorizationFailedErrCode, "authorization failed")

	// User
	UserExistedError = NewErrNo(ParamErrorCode, "user existed")

	// Comment
	UnexpectedTypeError     = NewErrNo(UnexpectedTypeErrorCode, "unexpected type")
	NotImplementError       = NewErrNo(NotImplementErrorCode, "not implement")
	SensitiveWordsError     = NewErrNo(SensitiveWordsErrorCode, "existed sensitive words")
	SensitiveWordsHTTPError = NewErrNo(ServiceErrorCode, "sensiWords api error")
)
