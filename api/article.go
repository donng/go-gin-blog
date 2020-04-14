package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-blog/model"
	"net/http"
)

type getArticles struct {
	Page int `json:"page"`
	Num  int `json:"num"`
}

func GetArticles(c *gin.Context) {
	var params getArticles
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articles := model.GetArticles(params.Page-1, params.Num)

	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}
