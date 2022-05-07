package model

/*
    评论存储模型
	目前的想法是通过视频ID+用户ID索引出该用户在当前视频下的所有评论ID
	评论ID仅对当前视频+当前用户唯一以提高查询效率并避免爆int范围

	TODO: 客户端放出后， 需要检查客户端对删除评论按钮的具体实现 防止意外的评论删除操作
*/

type Comment struct {
	Common
	UserID      uint64 `json:"user_id"`      // 用户ID
	VideoID     uint64 `json:"video_id"`     // 视频ID
	CommentText string `json:"comment_text"` // 评论内容
}
