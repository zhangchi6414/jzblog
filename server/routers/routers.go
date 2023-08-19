package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "jzblog/docs"
	"jzblog/server/middleware"
	v1 "jzblog/server/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//国际语言
	r.Use(middleware.Translations())
	article := v1.NewArticle()
	tag := v1.NewTag()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
