package service

import "github.com/gin-gonic/gin"

func NewAdminService() *AdminService {
	return &AdminService{}
}

type AdminService struct{}

func (a AdminService) Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "i working",
	})
}
