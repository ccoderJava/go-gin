package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-gin-examples/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, err := com.StrTo(c.Query("page")).Int()
	if err != nil {
		return 0
	}
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
