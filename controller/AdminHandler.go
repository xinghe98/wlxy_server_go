package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xinghe98/wlxy_server_go/types"
)

func NewAdminHandler(admin types.Commoner) *adminCo {
	return &adminCo{admin}
}

type adminCo struct {
	admin types.Commoner
}

func (a *adminCo) HelloHandler(c *gin.Context) {
	a.admin.Hello(c)
}
