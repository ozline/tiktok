package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode                = 10000
	ServiceErrorCode           = 10001
	ParamErrorCode             = 10002
	AuthorizationFailedErrCode = 10003
)

var (
	Success                  = NewErrNo(SuccessCode, "Success")
	ServiceError             = NewErrNo(ServiceErrorCode, "Service is unable to start successfully")
	ParamError               = NewErrNo(ParamErrorCode, "Parameter error")
	AuthorizationFailedError = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
)

type ErrNo struct {
	ErrorCode int64
	ErrorMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("error code: %d, error msg: %s", e.ErrorCode, e.ErrorMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrorMsg = msg
	return e
}

// ConvertErr convert error to ErrNo
func ConvertErr(err error) ErrNo {
	errno := ErrNo{}
	if errors.As(err, &errno) {
		return errno
	}

	s := ServiceError
	s.ErrorMsg = err.Error()
	return s
}
