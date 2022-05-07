package controller

import (
	"github.com/YouthCampProj/douyin/pkg/auth"
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"github.com/YouthCampProj/douyin/service"
	"github.com/YouthCampProj/douyin/utils"
	"github.com/gin-gonic/gin"
)

// InitPublishRoute 初始化投稿相关路由
// /douyin/publish
func InitPublishRoute(r *gin.RouterGroup) {
	r.POST("/action", PublishAction)
	r.GET("/list", GetPublishList)
}

// PublishAction
// 登录用户选择视频上传
// POST /douyin/publish/action/
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18875092
func PublishAction(c *gin.Context) {
	token := c.PostForm("token")
	user, err := auth.ParseToken(token)
	if err != nil {
		c.JSON(200, serializer.BuildPublishActionResponse(serializer.CodePublishTokenInvalid))
		return
	}
	data, err := c.FormFile("data")
	code := utils.PublishDataTest(data)
	if code != 0 {
		c.JSON(200, serializer.BuildPublishActionResponse(code))
		return
	}

	PublishActionService := &service.PublishActionService{
		User:       user,
		FileHeader: data,
	}
	c.JSON(200, PublishActionService.Publish())
}

// GetPublishList
// 登录用户的视频发布列表
// 直接列出用户所有投稿过的视频
// GET /douyin/publish/list/
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18901444
func GetPublishList(c *gin.Context) {
	//token := c.Query("token")
	// TODO 发布列表接口
}
