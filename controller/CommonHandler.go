package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xinghe98/wlxy_server_go/types"
)

func NewCommonHandler(common types.Commoner) *CommonCo {
	return &CommonCo{common}
}

type CommonCo struct {
	common types.Commoner
}

func (a *CommonCo) HelloHandler(c *gin.Context) {
	a.common.Hello(c)
}
