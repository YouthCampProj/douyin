package controller

import (
	"github.com/YouthCampProj/douyin/pkg/auth"
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"github.com/YouthCampProj/douyin/service"
	"github.com/YouthCampProj/douyin/utils"
	"github.com/gin-gonic/gin"
)

// InitCommentRoute 初始化评论相关路由
// /douyin/comment
func InitCommentRoute(r *gin.RouterGroup) {
	r.POST("/action/", CommentAction)
	r.GET("/list", GetCommentList)
}

// CommentAction 评论操作
// 登录用户对视频进行评论
// POST /douyin/comment/action
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902509
func CommentAction(c *gin.Context) {
	// TODO 需要对参数进行验证

	token := c.Query("token")              // 用户token
	videoIDstr := c.Query("video_id")      // 视频ID
	actionType := c.Query("action_type")   // 1-发布评论, 2-删除评论
	commentText := c.Query("comment_text") // 评论内容,在action_type=1时使用
	commentIDstr := c.Query("comment_id")  // 要删除的评论ID,在action_type=2时使用

	videoID := utils.Str2uint64(videoIDstr)
	commentID := utils.Str2uint64(commentIDstr)

	user, err := auth.ParseToken(token)
	if err != nil {
		c.JSON(200, serializer.BuildCommentActionResponse(serializer.CodeCommentTokenInvalid))
	}
	commentActionService := &service.CommentActionService{
		UserID:      user.ID,
		VideoID:     videoID,
		ActionType:  actionType,
		CommentID:   commentID,
		CommentText: commentText,
	}
	c.JSON(200, commentActionService.Execute())
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
