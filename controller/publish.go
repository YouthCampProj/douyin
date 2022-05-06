package controller

import "github.com/gin-gonic/gin"

type PublishActionResponse struct {
	Response
}

type PublishListResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"` // 用户发布的视频列表
}

// PublishAction
// 登录用户选择视频上传
// POST /douyin/publish/action/
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18875092
func PublishAction(c *gin.Context) {
	//userid := c.Query("user_id")
	//token := c.Query("token")
	//data, err := c.FormFile("data")
	// TODO 投稿接口
}

// GetPublishList
// 登录用户的视频发布列表
// 直接列出用户所有投稿过的视频
// GET /douyin/publish/list/
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18901444
func GetPublishList(c *gin.Context) {
	//userid := c.Query("user_id")
	//token := c.Query("token")
	// TODO 发布列表接口
}
