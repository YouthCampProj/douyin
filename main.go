package main

import (
	"github.com/YouthCampProj/douyin/controller"
	"github.com/YouthCampProj/douyin/pkg/config"
	"github.com/YouthCampProj/douyin/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	service.Init()
	controller.Init(r)
	if config.Conf.Debug {
		gin.SetMode(gin.DebugMode)
	}
	listenIP := config.Conf.Site.IP
	listenPort := strconv.Itoa(config.Conf.Site.Port)
	r.Run(listenIP + ":" + listenPort)
}
