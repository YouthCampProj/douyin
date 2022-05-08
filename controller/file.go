package controller

import "github.com/gin-gonic/gin"

func InitFileRoute(r *gin.RouterGroup) {
	r.Static("/", "uploads")
}
