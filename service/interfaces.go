package service

import "github.com/gin-gonic/gin"

type Adminer interface {
	Hello(ctx *gin.Context)
}
