package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xinghe98/wlxy_server_go/service"
)

func NewAdminHandler(admin service.Adminer) *adminCo {
	return &adminCo{admin}
}

type adminCo struct {
	admin service.Adminer
}

func (a *adminCo) HelloHandler(c *gin.Context) {
	a.admin.Hello(c)
}
