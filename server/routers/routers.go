package routers

import (
	"github.com/gin-gonic/gin"
	v1 "jzblog/server/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	article := v1.NewArticle()
	tag := v1.NewTag()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)
		apiv1.PUT("/tags/:id", tag.Get)

		apiv1.GET("/articles", article.List)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.POST("/articles", article.Create)

	}

	return r
}