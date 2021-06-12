package route

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init()  {
	router = gin.New()
	//router.Use(logger, recovery)
	//router.NoRoute(notRoute)
}

// Bootstrap 引导初始化路由route
func Bootstrap() *gin.Engine {
	return router
}
