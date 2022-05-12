package serializer

const (
	CodeCommentTokenInvalid = 5000 + iota
	CodeCommentVideoIDInvalid
	CodeCommentActionInvalid
	CodeCommentIDNotFound
	CodeCommentTextInvalid
	CodeCommentActionDBFailed
)

var CodeCommentMessage = map[int]string{
	CodeCommentTokenInvalid:   "Token is invalid",
	CodeCommentVideoIDInvalid: "Video ID is invalid",
	CodeCommentActionInvalid:  "Action is invalid",
	CodeCommentIDNotFound:     "Comment ID not found",
	CodeCommentTextInvalid:    "Comment text is invalid",
	CodeCommentActionDBFailed: "DB failed",
}

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

func BuildCommentActionResponse(status int) *CommentActionResponse {
	res := &CommentActionResponse{}
	res.Response = NewResponse(status, CodeCommentMessage[status])
	return res
}
