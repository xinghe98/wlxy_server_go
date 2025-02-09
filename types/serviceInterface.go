package types

import "github.com/gin-gonic/gin"

type Commoner interface {
	Hello(ctx *gin.Context)
}
