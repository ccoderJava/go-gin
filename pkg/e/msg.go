package e

var MsgFlags = map[int]string{
	SUCCESS:        "OK",
	ERROR:          "FAIL",
	INVALID_PARAMS: "请求参数错误",

	ERROR_EXISTS_TAG:        "已存在该标签内容",
	ERROR_NOT_EXISTS_TAG:    "不存在该标签内容",
	ERROR_NOT_EXIST_ARTICLE: "该文章不存在",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token鉴权超时",
	ERROR_AUTH_TOEKN:               "token生成失败",
	ERROR_AUTH:                     "token错误",
}

// GetMsg 获取错误码描述
// code 错误码
func GetMsg(code int) string {
	if msg, ok := MsgFlags[code]; ok {
		return msg
	}
	return MsgFlags[ERROR]

}
