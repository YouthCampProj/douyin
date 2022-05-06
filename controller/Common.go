package controller

// User 用户信息
type User struct {
	ID            int    `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户名称
	FollowCount   int    `json:"follow_count"`   // 关注总数
	FollowerCount int    `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}

// Response 常规响应信息
type Response struct {
	StatusCode int    `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
}

// Video 视频信息
type Video struct {
	ID            int    `json:"id"` // 视频唯一标识
	Author        User   `json:"author"`
	PlayURL       string `json:"play_url"`       // 视频播放地址
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int    `json:"favorite_count"` // 视频的点赞总数
	CommentCount  int    `json:"comment_count"`  // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞

}

// Comment 评论信息
type Comment struct {
	ID         int    `json:"id"`          // 评论id
	User       User   `json:"user"`        // 评论用户
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论时间
}
