package app

import (
	"github.com/gin-gonic/gin"
	"go-gin-blog/pkg/e"
	"net/http"
)

type responseData struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Response(c *gin.Context, code int, data interface{}, message string) {
	c.JSON(http.StatusOK, responseData{
		Code:    code,
		Data:    data,
		Message: message,
	})
}

func Success(c *gin.Context, data interface{}) {
	Response(c, e.SUCCESS, data, e.GetErrMsg(e.SUCCESS))
}

func Fail(c *gin.Context, code int) {
	Response(c, code, "", e.GetErrMsg(code))
}
