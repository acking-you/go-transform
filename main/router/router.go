package router

import (
	"github.com/gin-gonic/gin"
	"main/middleware"
	"net/http"
	"transform/http/api"
)

func Start() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.NotifyUDP) //每次业务执行时都将使得后台ydp轮询服务关闭
	r.StaticFS("/static", http.Dir("../static"))

	baseRouter := r.Group("/transform")

	baseRouter.POST("file", api.PostFiles)
	baseRouter.OPTIONS("file", func(context *gin.Context) {
		context.Header("Access-Control-Max-Age", "2592000")
		context.String(http.StatusOK, "don't talk")
	})

	baseRouter.GET("files", api.GetFiles)

	return r
}
