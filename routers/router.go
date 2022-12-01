package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-gin/docs"
	"go-gin/middleware/jwt"
	"go-gin/pkg/setting"
	"go-gin/routers/api"
	"go-gin/routers/api/v1"
)

func InitRouter() *gin.Engine {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Golang Gin API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "http://localhost:8000"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//新增auth接口
	r.GET("/auth", api.GetAuth)

	var apiv1 = r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签
		apiv1.GET("/tags", v1.GetTags)
		//新增标签
		apiv1.POST("/tags", v1.AddTag)
		//修改标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取单个文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//修改文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

	}

	return r
}
