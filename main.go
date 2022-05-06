package main

import (
	"github.com/YouthCampProj/douyin/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.Init(r.Group("/douyin"))
	r.Run(":8080")
}
