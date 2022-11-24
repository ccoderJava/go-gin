package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-gin-examples/models"
	"go-gin-examples/pkg/e"
	"go-gin-examples/pkg/setting"
	"go-gin-examples/pkg/util"
	"net/http"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// AddTag 新增标签
func AddTag(c *gin.Context) {

}

// EditTag 修改标签
func EditTag(c *gin.Context) {

}

//DeleteTag 删除标签
func DeleteTag(c *gin.Context) {

}
