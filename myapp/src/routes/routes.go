package routes

import (
	"myapp/src/controllers"

	"github.com/gin-gonic/gin"
)

type RouterOptions func(router *gin.Engine)

var opts = []RouterOptions{
	controllers.SetupWelcomeRouters,
	controllers.SetupStatRouters,
}

// InitRoutes 初始化路由规则
func InitRoutes(router *gin.Engine) {
	for _, opt := range opts {
		opt(router)
	}
}
