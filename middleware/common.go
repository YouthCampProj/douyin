package middleware

import (
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"github.com/gin-gonic/gin"
)

// SendCommonResponse 发送仅包含状态码与状态信息的JSON响应
func SendCommonResponse(code int, message string, c *gin.Context) {
	c.JSON(200, serializer.NewResponse(code, message))
}
