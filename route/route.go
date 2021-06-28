package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jjonline/go-mod-library/logger"
	"github.com/jjonline/golang-backend/conf"
)

var (
	router *gin.Engine
)

func iniRoute() {
	router = gin.New()
	router.AppEngine = true // 启用AppEngine模式; nginx反代通过`X-Appengine-Remote-Addr`头透传客户端真实IP

	// set base middleware
	router.Use(logger.GinLogger, logger.GinRecovery)
	if conf.Config.Server.Cors {
		router.Use(logger.GinCors)
	}
}

// Bootstrap 引导初始化路由route
func Bootstrap() *gin.Engine {
	iniRoute()
	return router
}
