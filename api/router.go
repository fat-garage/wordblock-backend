package api

import (
	"github.com/fat-garage/wordblock-backend/api/handler"
	"github.com/fat-garage/wordblock-backend/api/middleware"
	"github.com/fat-garage/wordblock-backend/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(router *gin.Engine) {
	router.Use(middleware.Cors())

	docs.SwaggerInfo_swagger.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1Router := router.Group("/api/v1")
	initNoAuthRouter(v1Router)
	//initNeedAuthRouter(v1Router)
}

func initNoAuthRouter(r *gin.RouterGroup) {
	r.GET("ipfs/list/:cid", handler.GetDIDCidList)
	r.POST("ipfs/upload", handler.Upload)
}

func initNeedAuthRouter(r *gin.RouterGroup) {
	r.Use(middleware.JWTAuth())
}
