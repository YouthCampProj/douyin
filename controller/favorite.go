package controller

import "github.com/gin-gonic/gin"

type FavoriteActionResponse struct {
	Response
}

// FavoriteAction 赞操作
// 登录用户对视频的点赞和取消点赞操作
// POST /douyin/favorite/action/
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902441
func FavoriteAction(c *gin.Context) {
	//userID := c.Query("user_id")         // 用户ID
	//token := c.Query("token")            // 用户token
	//videoID := c.Query("video_id")       // 视频ID
	//actionType := c.Query("action_type") // 1-点赞, 2-取消点赞
	// TODO 赞操作接口
}

type FavoriteListResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
}

// GetFavoriteList 获取点赞列表
// 登录用户的所有点赞视频列表
// GET /douyin/favorite/list/
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902464
func GetFavoriteList(c *gin.Context) {
	//userID := c.Query("user_id")         // 用户ID
	//token := c.Query("token")            // 用户token
	// TODO 点赞列表接口
}
