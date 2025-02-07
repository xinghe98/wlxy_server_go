package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xinghe98/wlxy_server_go/router"
)

func main() {
	r := gin.Default()
	// 注册路由
	r = router.RegistRouters(r)

	r.Run(":3000")
}
