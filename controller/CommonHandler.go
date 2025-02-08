package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xinghe98/wlxy_server_go/service"
)

func NewCommonHandler(common service.Adminer) *CommonCo {
	return &CommonCo{common}
}

type CommonCo struct {
	common service.Adminer
}

func (a *CommonCo) HelloHandler(c *gin.Context) {
	a.common.Hello(c)
}
