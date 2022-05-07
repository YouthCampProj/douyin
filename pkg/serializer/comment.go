package serializer

// Comment 评论信息
type Comment struct {
	ID         uint64 `json:"id"`          // 评论id
	User       User   `json:"user"`        // 评论用户
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论时间
}

// CommentActionResponse 评论操作响应
type CommentActionResponse struct {
	Response
}

// CommentListResponse 评论列表响应
type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}
