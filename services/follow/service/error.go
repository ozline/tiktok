package service

const (
	Success             = 0
	Error               = 10000
	ErrorDBRecord       = 10001
	InvalidParams       = 20000
	InvalidRelationType = 20001
)

var MsgFlags = map[int64]string{
	Success:             "success",
	Error:               "fail",
	ErrorDBRecord:       "数据库返回错误",
	InvalidParams:       "参数错误",
	InvalidRelationType: "无此类型关系操作",
}

func GetMsg(code int64) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
