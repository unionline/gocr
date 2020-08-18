package server

import (
	"github.com/gin-gonic/gin"
	"gocr/routers/middleware"
)

// Initialize ...
func Initialize() *gin.Engine {

	var r = gin.Default()

	// 校验参数
	routerFile(r)

	r.Use(middleware.ValidPara)
	routerAPI(r)
	routerTest(r)

	return r
}
