package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xinghe98/wlxy_server_go/service"
)

func NewAdminHandler() *adminCo {
	// 直接创建service实例，注入依赖
	return &adminCo{service.NewAdminService()}
}

type adminCo struct {
	service service.Adminer
}

func (a *adminCo) HelloHandler(c *gin.Context) {
	a.service.Hello(c)
}
