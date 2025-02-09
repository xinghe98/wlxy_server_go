package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xinghe98/wlxy_server_go/dao"
)

func NewAdminService(db *dao.DBService) *AdminService {
	return &AdminService{db}
}

type AdminService struct {
	db *dao.DBService
}

func (a AdminService) Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "i working",
	})
}
