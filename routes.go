package main

import (
	"ginVue/controller"
	_ "ginVue/docs"
	"ginVue/middleWare"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleWare.CORSMiddleware())

	r.GET("/docs/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.POST("/api/v1/auth/register", controller.Register)
	r.POST("/api/v1/auth/login", controller.Login)
	r.GET("/api/v1/auth/info", middleWare.AuthMiddleware(), controller.Info) //使用中间件保护用户信息接口
	//r.POST("api/v1/coal/getOneBalance",middleWare.AuthMiddleware(),controller.ViewOneCoalBalance)

	return r
}
