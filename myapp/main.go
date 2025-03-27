package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"myapp/pkg/constant"
	"myapp/pkg/logging"
	"myapp/src/routes"
	"myapp/src/utils/uip"
	"myapp/src/utils/uredis"
)

func main() {
	logLevel := os.Getenv(constant.LOG_LEVEL)

	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(lvl)
	}

	redis_url := os.Getenv(constant.APP_REDIS_URL)
	uredis.Open(redis_url)

	uip.Init()

	router := gin.Default()
	router.Use(logging.RequestIDMiddleware())

	// 初始化路由
	routes.InitRoutes(router)

	// 运行服务器
	router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
