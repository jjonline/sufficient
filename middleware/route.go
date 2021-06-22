package middleware

import "github.com/gin-gonic/gin"

// router 全局的路由变量
var router *gin.Engine

func Bootstrap() *gin.Engine {
	router = gin.New()
	// 启用AppEngine模式，nginx反代通过`X-Appengine-Remote-Addr`头透传客户端真实IP
	router.AppEngine = true

	return router
}
