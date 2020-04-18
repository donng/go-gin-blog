package e

const SUCCESS = 0
const FAIL = 1

const ParamsError = 1000

const ArticleCreateError = 2000
const ArticleFindError = 2001
const ArticleGetError = 2002
const ArticleRemoveError = 2003
const ArticleModifyError = 2004

var errorMessage = map[int]string{
	SUCCESS: "请求成功",
	FAIL:    "请求失败",

	ParamsError: "参数错误",

	ArticleCreateError: "文章创建失败",
	ArticleFindError:   "文章获取失败",
	ArticleGetError:    "文章查询失败",
	ArticleRemoveError: "文章删除失败",
	ArticleModifyError: "文章修改失败",
}

func GetErrMsg(code int) string {
	if message, ok := errorMessage[code]; ok {
		return message
	}
	return ""
}
