package controller

import (
	"github.com/YouthCampProj/douyin/pkg/auth"
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"github.com/YouthCampProj/douyin/service"
	"github.com/YouthCampProj/douyin/utils"
	"github.com/gin-gonic/gin"
)

// InitFavoriteRoute 初始化点赞相关路由
// /douyin/favorite
func InitFavoriteRoute(r *gin.RouterGroup) {
	r.POST("action/", FavoriteAction)
	r.GET("list/", GetFavoriteList)
}

// FavoriteAction 赞操作
// 登录用户对视频的点赞和取消点赞操作
// POST /douyin/favorite/action/
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902441
func FavoriteAction(c *gin.Context) {
	if !utils.UserIDTest(c.Query("user_id")) {
		c.JSON(200, serializer.BuildPublishActionResponse(serializer.CodeUserIDInvalid))
	}
	userID := utils.Str2uint64(c.Query("user_id"))
	if !utils.VideoIDTest(c.Query("video_id")) {
		c.JSON(200, serializer.BuildPublishActionResponse(serializer.CodeFavoriteVideoIDInvalid))
	}
	videoID := utils.Str2uint64(c.Query("video_id"))
	if !utils.FavoriteActionTypeTest(c.Query("action_type")) {
		c.JSON(200, serializer.BuildPublishActionResponse(serializer.CodeFavoriteActionTypeInvalid))
	}
	actionType := utils.Str2int32(c.Query("action_type"))
	if !auth.CheckToken(c.Query("token")) {
		c.JSON(200, serializer.BuildFavoriteActionResponse(serializer.CodeFavoriteTokenInvalid))
	}

	favoriteActionService := &service.FavoriteActionService{
		UserID:  userID,
		VideoID: videoID,
		Action:  actionType,
	}
	c.JSON(200, favoriteActionService.Execute())
}

// GetFavoriteList 获取点赞列表
// 登录用户的所有点赞视频列表
// GET /douyin/favorite/list/
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902464
func GetFavoriteList(c *gin.Context) {
	if !utils.UserIDTest(c.Query("user_id")) {
		c.JSON(200, serializer.BuildFavoriteListResponse(serializer.CodeUserIDInvalid, nil))
	}
	//token := c.Query("token") // 用户token
	// TODO 点赞列表接口
}
