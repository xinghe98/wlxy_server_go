package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xinghe98/wlxy_server_go/controller"
	"github.com/xinghe98/wlxy_server_go/dao"
	"github.com/xinghe98/wlxy_server_go/service"
)

func commonRouter(r *gin.Engine) *gin.Engine {
	// 统一做依赖注入
	dao, _ := dao.NewDBService("mongodb://admin:pwd@43.139.45.187:27019", "admin")
	commonHandler := controller.NewCommonHandler(service.NewAdminService(dao))

	r.GET("/ping", commonHandler.HelloHandler)
	return r
}
