package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-blog/model"
	"go-gin-blog/pkg/app"
	"go-gin-blog/pkg/e"
)

type getArticles struct {
	Page int `form:"page"`
	Num  int `form:"num"`
}

type createArticle struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func GetArticles(c *gin.Context) {
	var params getArticles
	if err := c.ShouldBindQuery(&params); err != nil {
		app.Fail(c, e.ParamsError)
		return
	}
	articles := model.GetArticles(params.Page-1, params.Num)

	app.Success(c, articles)
}

func CreateArticle(c *gin.Context) {
	var id uint
	var err error
	var params createArticle
	if err := c.ShouldBindJSON(&params); err != nil {
		app.Fail(c, e.ParamsError)
		return
	}

	if id, err = model.CreateArticle(params.Title, params.Content); err != nil {
		app.Fail(c, e.ArticleCreateError)
	}

	app.Success(c, id)
}
