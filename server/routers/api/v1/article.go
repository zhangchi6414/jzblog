package v1

import (
	"github.com/gin-gonic/gin"
	"jzblog/pkg/app"
	"jzblog/pkg/errcode"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (t Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorRespone(errcode.ServerError)
}

func (t Article) List(c *gin.Context) {

}

func (t Article) Create(c *gin.Context) {

}
func (t Article) Update(c *gin.Context) {

}
func (t Article) Delete(c *gin.Context) {

}
