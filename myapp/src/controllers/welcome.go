package controllers

import (
	"fmt"
	"myapp/src/utils/udomain"
	"myapp/src/utils/utime"

	"github.com/gin-gonic/gin"
)

func SetupWelcomeRouters(e *gin.Engine) {
	e.GET("/", RenderIndex)
	e.GET("/ping", RenderPing)
}

// RenderIndex 渲染首页
func RenderIndex(c *gin.Context) {
	c.String(200, "404 not found")
}

func RenderPing(c *gin.Context) {
	udomain.CheckDomainStat()
	c.String(200, fmt.Sprintf("%v", utime.UnixSec()))
}
