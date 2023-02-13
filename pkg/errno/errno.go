// DO NOT EDIT

package errno

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrorCode int64
	ErrorMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("error code: %d, error msg: %s", e.ErrorCode, e.ErrorMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrorCode: code,
		ErrorMsg: msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrorMsg = msg
	return e
}

// ConvertErr convert error to ErrNo
// in Default user ServiceErrorCode
func ConvertErr(err error) ErrNo {
	errno := ErrNo{}
	if errors.As(err, &errno) {
		return errno
	}

	s := ServiceError
	s.ErrorMsg = err.Error()
	return s
}
