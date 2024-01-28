package utils

// error msg
var (
	OK = 200

	TOKEN_ERROR = 1001
	LOGIN_ERROR = 1002
)

var codeMsg = map[int]string{
	OK: "OK",
	TOKEN_ERROR: "令牌错误",
	LOGIN_ERROR: "用户名或密码错误",
}


func GetErrMsg(code int) string {
	return codeMsg[code]
}