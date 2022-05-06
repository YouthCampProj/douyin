package controller

import "github.com/gin-gonic/gin"

type RelationActionResponse struct {
	Response
}

// RelationAction 关注操作
// POST /douyin/relation/action/
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902556
func RelationAction(c *gin.Context) {
	//userID := c.Query("user_id")         // 用户ID
	//token := c.Query("token")            // 用户token
	//toUserID := c.Query("to_user_id")    // 对方用户ID
	//actionType := c.Query("action_type") // 1-关注 2-取消关注
	// TODO 关注操作接口
}

// RelationFollowListResponse 关注列表的响应体
type RelationFollowListResponse struct {
	Response
	UserList []User `json:"user_list,omitempty"` // 用户信息列表
}

// GetRelationFollowList 获取关注列表
// GET /douyin/relation/follow/list/
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902568
func GetRelationFollowList(c *gin.Context) {
	//userID := c.Query("user_id") // 用户ID
	//token := c.Query("token")    // 用户token
	//TODO 关注列表接口
}

type RelationFollowerListResponse struct {
	Response
	UserList []User `json:"user_list,omitempty"` // 用户列表
}

// GetRelationFollowerList 获取粉丝列表
// GET /douyin/relation/follower/list/
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902563
func GetRelationFollowerList(c *gin.Context) {
	//userID := c.Query("user_id") // 用户ID
	//token := c.Query("token")    // 用户token
	// TODO 粉丝列表接口
}
