package e

// MsgFlags 错误信息
var MsgFlags = map[int]string{
	SUCCESS:      "成功",
	ERROR:        "错误",
	BAD_REQUEST:  "失败",
	UNAUTHORIZED: "未登录",
	FORBIDDEN:    "访问拒绝没有权限",
	// 业务编码从10000开始
	ERROR_EXIST_USER: "用户已存在",
	ERROR_USER_OR_PASSWORD: "错误的用户名或密码",
}

// GetMsg 获取错误信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
