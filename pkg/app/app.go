package app

import (
	"github.com/gin-gonic/gin"
	"jzblog/pkg/errcode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}
type Pager struct {
	page      int `json:"page,omitempty"`
	pageSize  int `json:"page_size,omitempty"`
	TotalRows int `json:"total_rows,omitempty"`
}

func NewResponse(c *gin.Context) *Response {
	return &Response{Ctx: c}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			page:      GetPage(r.Ctx),
			pageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}
func (r *Response) ToErrorRespone(err *errcode.Error) {
	reponse := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		reponse["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), reponse)
}
