package controller

import (
	"github.com/YouthCampProj/douyin/pkg/auth"
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"github.com/YouthCampProj/douyin/service"
	"github.com/YouthCampProj/douyin/utils"
	"github.com/gin-gonic/gin"
)

// InitRelationRoutes 初始化互动相关接口
// /douyin/relation
func InitRelationRoutes(r *gin.RouterGroup) {
	r.POST("action/", RelationAction)
	r.GET("follow/list/", GetRelationFollowList)
	r.GET("follower/list/", GetRelationFollowerList)
}

// RelationAction 关注操作
// POST /douyin/relation/action/
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902556
func RelationAction(c *gin.Context) {
	token := c.Query("token") // 用户token
	if !auth.CheckToken(token) {
		c.JSON(200, serializer.BuildFavoriteActionResponse(serializer.CodeRelationTokenInvalid))
	}
	toUserIDstr := c.Query("to_user_id") // 对方用户ID
	actionType := c.Query("action_type") // 1-关注 2-取消关注
	// TODO 关注操作接口

	user, err := auth.ParseToken(token)
	if err != nil {
		c.JSON(200, serializer.BuildFavoriteActionResponse(serializer.CodeRelationTokenInvalid))
	}
	relationActionService := &service.RelationActionService{
		UserID:     user.ID,
		ToUserID:   utils.Str2uint64(toUserIDstr),
		ActionType: actionType,
	}
	c.JSON(200, relationActionService.Execute())
}

// GetRelationFollowList 获取关注列表
// GET /douyin/relation/follow/list/
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902568
func GetRelationFollowList(c *gin.Context) {
	userIDstr := c.Query("user_id") // 用户ID
	token := c.Query("token")       // 用户token
	if !auth.CheckToken(token) {
		c.JSON(200, serializer.BuildRelationFollowListResponse(serializer.CodeRelationTokenInvalid, nil))
		return
	}
	//TODO 关注列表接口

	userID := utils.Str2uint64(userIDstr)
	requestFromUser, err := auth.ParseToken(token)
	if err != nil {
		c.JSON(200, serializer.BuildRelationFollowListResponse(serializer.CodeRelationTokenInvalid, nil))
		return
	}
	relationFollowListService := &service.RelationFollowListService{
		UserID:        userID,
		RequestFromID: requestFromUser.ID,
	}
	c.JSON(200, relationFollowListService.GetFollowList())

}

// GetRelationFollowerList 获取粉丝列表
// GET /douyin/relation/follower/list/
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902563
func GetRelationFollowerList(c *gin.Context) {
	userIDstr := c.Query("user_id") // 用户ID
	token := c.Query("token")       // 用户token
	// TODO 粉丝列表接口

	userID := utils.Str2uint64(userIDstr)
	requestFromUser, err := auth.ParseToken(token)
	if err != nil {
		c.JSON(200, serializer.BuildRelationFollowListResponse(serializer.CodeRelationTokenInvalid, nil))
		return
	}
	relationFollowerListService := &service.RelationFollowerListService{
		UserID:        userID,
		RequestFromID: requestFromUser.ID,
	}
	c.JSON(200, relationFollowerListService.GetFollowerList())
}
