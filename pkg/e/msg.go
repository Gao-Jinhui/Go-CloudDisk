package e

var MsgFlags = map[int]string{
	Success:                    "ok",
	Error:                      "fail",
	InvalidParams:              "参数验证错误",
	InvalidRequest:             "请求格式错误",
	ErrorWrongPassword:         "密码错误",
	InsufficientParameters:     "参数数量不足",
	ErrorExistUser:             "用户已存在",
	ErrorNotExistUser:          "该用户不存在",
	ErrorNotExistArticle:       "该文章不存在",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
	ErrorTokenGenerate:         "Token生成失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
