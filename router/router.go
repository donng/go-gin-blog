package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-blog/api"
)

// 路由要完成的功能：路由初始化所有请求后，返回供 main.go 启动监听服务
// 不适用 init() 是因为 init 不能有返回值
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	addRoutes(r)

	return r
}

func addRoutes(r *gin.Engine) {
	router := r.Group("/api")

	router.Use()
	{
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		router.GET("/articles", api.GetArticles)
		router.POST("/article", api.CreateArticle)
	}

}
