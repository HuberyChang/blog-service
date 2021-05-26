/*
* @Author: HuberyChang
* @Date: 2021/5/8 22:03
 */

package routes

import (
	"net/http"
	"time"

	"github.com/HuberyChang/blog-service/pkg/limiter"

	"github.com/HuberyChang/blog-service/global"
	"github.com/HuberyChang/blog-service/internal/middleware"
	"github.com/HuberyChang/blog-service/internal/routes/api"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/HuberyChang/blog-service/docs"
	v1 "github.com/HuberyChang/blog-service/internal/routes/api/v1"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.POST("/auth", api.GetAuth)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.Use(middleware.Tracing())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r

}
