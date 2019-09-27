package Constant

const (
	SUCCESS = 200 // 成功
	TOAST = 400 // 前端toast提示信息
	ERROR = 500 // 系统错误

	// 以下是业务错误码
	SystemUnavailable = 10000
	InvalidParams     = 10001
	IllegalRequest     = 10002
)

var MsgFlags = map[int]string {
	SUCCESS :          "成功",
	ERROR :            "失败",
	TOAST :            "提示",
	InvalidParams:     "请求参数错误",
	SystemUnavailable: "服务不可用",
	IllegalRequest 	: "非法请求",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}