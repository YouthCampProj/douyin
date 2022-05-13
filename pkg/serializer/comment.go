package serializer

import "github.com/YouthCampProj/douyin/model"

const (
	CodeCommentTokenInvalid = 5000 + iota
	CodeCommentVideoIDInvalid
	CodeCommentActionInvalid
	CodeCommentIDNotFound
	CodeCommentTextInvalid
	CodeCommentDBFailed
)

var CodeCommentMessage = map[int]string{
	CodeCommentTokenInvalid:   "Token is invalid",
	CodeCommentVideoIDInvalid: "Video ID is invalid",
	CodeCommentActionInvalid:  "Action is invalid",
	CodeCommentIDNotFound:     "Comment ID not found",
	CodeCommentTextInvalid:    "Comment text is invalid",
	CodeCommentDBFailed:       "DB failed",
}

// Comment 评论信息
type Comment struct {
	ID         uint64 `json:"id"`          // 评论id
	User       *User  `json:"user"`        // 评论用户
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

// BuildCommentActionResponse 构建评论操作响应
func BuildCommentActionResponse(status int) *CommentActionResponse {
	res := &CommentActionResponse{}
	res.Response = NewResponse(status, CodeCommentMessage[status])
	return res
}

// BuildCommentListResponse 构造评论列表响应
func BuildCommentListResponse(status int, commentList []*model.CommentUserBundle) *CommentListResponse {
	res := &CommentListResponse{}
	res.Response = NewResponse(status, CodeCommentMessage[status])
	if status != CodeSuccess {
		return res
	}
	res.CommentList = make([]Comment, len(commentList))
	for i, v := range commentList {
		res.CommentList[i] = Comment{
			ID:         v.ID,
			User:       (*User)(v.User), // 强转 以后需要进行优化
			Content:    v.Content,
			CreateDate: v.CreateDate.Format("01-02"),
		}
	}
	return res
}
