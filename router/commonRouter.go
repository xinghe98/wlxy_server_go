package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xinghe98/wlxy_server_go/controller"
	"github.com/xinghe98/wlxy_server_go/service"
)

func commonRouter(r *gin.Engine) *gin.Engine {
	// 统一做依赖注入
	commonHandler := controller.NewCommonHandler(service.NewCommonService())

	r.GET("/ping", commonHandler.HelloHandler)
	return r
}
