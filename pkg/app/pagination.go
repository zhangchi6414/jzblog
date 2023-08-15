package app

import (
	"github.com/gin-gonic/gin"
	"jzblog/global"
	"jzblog/pkg/convert"
)

// 分页处理
func GetPage(c *gin.Context) int {

	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > 0 {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}
func GetPageOffset(pag, pageSize int) int {
	result := 0
	if pag > 0 {
		result = (pag - 1) * pageSize
	}
	return result
}
