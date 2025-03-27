package controllers

import (
	"myapp/pkg/types"
	"myapp/src/utils/udomain"
	"myapp/src/utils/uip"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupStatRouters(e *gin.Engine) {
	g := e.Group("/stat")
	f := &StatDomainFunc{}
	g.POST("/domain", f.StatDomain)
	g.GET("/check_timer", f.CheckTimer)
}

type StatDomainFunc struct {
}

func (s *StatDomainFunc) StatDomain(c *gin.Context) {

	defer c.String(200, "done")

	var req types.StatDomainReq

	err := c.BindJSON(&req)
	if err != nil {
		logrus.WithContext(c).Debugf("StatDomain, bindjson: %v", err)
		return
	}

	logrus.
		WithContext(c).
		WithField("client_ip", req.IP).
		WithField("remote_ip", c.ClientIP()).
		WithField("stats", req.Stats).
		Debugln("StatDomain")

	if !uip.CheckIP(req.IP) {
		return
	}

	for _, s := range req.Stats {
		udomain.StatDomain(s.Domain, s.Stats)
	}
}

func (s *StatDomainFunc) CheckTimer(c *gin.Context) {
	udomain.CheckDomainStat()
	c.String(200, "done")
}
