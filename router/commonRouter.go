package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xinghe98/wlxy_server_go/controller"
)

func commonRouter(r *gin.Engine) *gin.Engine {
	commonHandler := controller.NewAdminHandler()
	r.GET("/ping", commonHandler.HelloHandler)
	return r
}
