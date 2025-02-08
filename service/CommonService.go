package service

import "github.com/gin-gonic/gin"

func NewCommonService() *CommonService {
	return &CommonService{}
}

type CommonService struct{}

func (a CommonService) Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "common work",
	})
}
