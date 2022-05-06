package middleware

import (
	"github.com/YouthCampProj/douyin/controller"
	"github.com/gin-gonic/gin"
)

// SendCommonResponse 发送仅包含状态码与状态信息的JSON响应
func SendCommonResponse(status int, message string, c *gin.Context) {
	res := &controller.Response{
		StatusCode: status,
		StatusMsg:  message,
	}
	c.JSON(200, res)
}
