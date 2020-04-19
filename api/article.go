package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-blog/model"
	"go-gin-blog/pkg/app"
	"go-gin-blog/pkg/e"
	"log"
)

func GetArticles(c *gin.Context) {
	var articles []model.Article
	var err error
	params := struct {
		Page  int `form:"page" binding:"required"`
		Num   int `form:"num" binding:"required"`
		TagID int `form:"tag_id"`
	}{}

	if err := c.ShouldBindQuery(&params); err != nil {
		log.Printf("params error: %s", err.Error())
		app.ShowError(c, e.ParamsError)
		return
	}

	log.Printf("request params: %+v", params)

	if articles, err = model.GetArticles(params.Page-1, params.Num, params.TagID); err != nil {
		log.Printf("get error: %s", err.Error())
		app.ShowError(c, e.ArticleGetError)
		return
	}

	app.ShowData(c, articles)
}

func CreateArticle(c *gin.Context) {
	var id uint
	var err error
	params := struct {
		Title   string `json:"title" binding:"required"`
		Desc    string `json:"desc"`
		Content string `json:"content" binding:"required"`
		TagID   int    `json:"tag_id"`
	}{}
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("params error: %s", err.Error())
		app.ShowError(c, e.ParamsError)
		return
	}

	log.Printf("request params: %+v", params)

	if id, err = model.CreateArticle(params.Title, params.Desc, params.Content, params.TagID); err != nil {
		app.ShowError(c, e.ArticleCreateError)
	}

	app.ShowData(c, gin.H{
		"id": id,
	})
}

func ShowArticle(c *gin.Context) {
	var err error
	var article model.Article
	var params = struct {
		ID int `json:"id" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("params error: %s", err.Error())
		app.ShowError(c, e.ParamsError)
		return
	}

	if article, err = model.FindArticle(params.ID); err != nil {
		log.Printf("find error: %s", err.Error())
		app.ShowError(c, e.ArticleFindError)
		return
	}

	app.ShowData(c, article)
}

func RemoveArticle(c *gin.Context) {
	var params = struct {
		ID int `json:"id" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("params error: %s", err.Error())
		app.ShowError(c, e.ParamsError)
		return
	}

	log.Printf("request params: %+v", params)

	if err := model.RemoveArticle(params.ID); err != nil {
		log.Printf("article remove error: %s", err.Error())
		app.ShowError(c, e.ArticleRemoveError)
		return
	}

	app.ShowData(c, "")
}

func ModifyArticle(c *gin.Context) {
	var params = struct {
		ID      int    `json:"id" binding:"required"`
		Title   string `json:"title" binding:"required"`
		Desc    string `json:"desc"`
		Content string `json:"content" binding:"required"`
		TagID   int    `json:"tag_id"`
	}{}

	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("params error: %s", err.Error())
		app.ShowError(c, e.ParamsError)
		return
	}

	log.Printf("request params: %+v", params)

	if err := model.ModifyArticle(params.ID, params.Title, params.Desc, params.Content, params.TagID); err != nil {
		log.Printf("article modify error: %s", err.Error())
		app.ShowError(c, e.ArticleModifyError)
		return
	}

	app.ShowData(c, "")
}
