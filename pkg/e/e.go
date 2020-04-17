package e

const SUCCESS = 0
const FAIL = 1

const ParamsError = 1000

const ArticleCreateError = 2000

var errorMessage = map[int]string{
	SUCCESS: "请求成功",
	FAIL:    "请求失败",

	ParamsError: "参数错误",

	ArticleCreateError: "日志创建失败",
}

func GetErrMsg(code int) string {
	if message, ok := errorMessage[code]; ok {
		return message
	}
	return ""
}
