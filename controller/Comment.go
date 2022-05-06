package controller

import "github.com/gin-gonic/gin"

type CommentActionResponse struct {
	Response
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

// CommentAction 评论操作
// 登录用户对视频进行评论
// POST /douyin/comment/action
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902509
func CommentAction(c *gin.Context) {
	//userID := c.Query("user_id")           // 用户ID
	//token := c.Query("token")              // 用户token
	//videoID := c.Query("video_id")         // 视频ID
	//actionType := c.Query("action_type")   // 1-发布评论, 2-删除评论
	//commentText := c.Query("comment_text") // 评论内容,在action_type=1时使用
	//commentID := c.Query("comment_id")     // 要删除的评论ID,在action_type=2时使用
	// TODO 评论操作接口
}

// GetCommentList 评论列表
// 查看视频的所有评论，按发布时间倒序排列
// GET /douyin/comment/list
func GetCommentList(c *gin.Context) {
	//userID := c.Query("user_id") // 用户ID
	//token := c.Query("token")    // 用户token
	//videoID := c.Query("video_id") // 视频ID
	// TODO 评论列表接口
}
