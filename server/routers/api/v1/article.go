package v1

import (
	"github.com/gin-gonic/gin"
	"jzblog/pkg/app"
	"jzblog/pkg/errcode"
)

type CountArticleRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State int8   `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State int8   `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Name     string `form:"name" binding:"max=100"`
	CreateBy string `form:"created_by" binding:"required,min=3,max=100"`
	State    int8   `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      int8   `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

// @Summary 获取文章
// @Produce  json
// @Param name query string false "文章名称" maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article/{id} [get]
func (t Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorRespone(errcode.ServerError)
}

// @Summary 获取多个文章
// @Produce  json
// @Param name query string false "文章名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (t Article) List(c *gin.Context) {

}

// @Summary 新增文章
// @Produce  json
// @Param name body string true "文章名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (t Article) Create(c *gin.Context) {

}

// @Summary 更新文章
// @Produce  json
// @Param id path int true "文章 ID"
// @Param name body string false "文章名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (t Article) Update(c *gin.Context) {

}

// @Summary 删除文章
// @Produce  json
// @Param id path int true "文章 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (t Article) Delete(c *gin.Context) {

}
