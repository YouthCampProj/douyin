package main

import (
	"github.com/YouthCampProj/douyin/controller"
	"github.com/YouthCampProj/douyin/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	service.Init()
	controller.Init(r.Group("/douyin"))
	r.Run(":8080")
}
