package model

// User 用户信息
type User struct {
	Common
	UserID        int    `json:"user_id"`        // 用户id
	Name          string `json:"name"`           // 用户名称
	Password      string `json:"password"`       // 用户密码
	FollowCount   int    `json:"follow_count"`   // 关注总数
	FollowerCount int    `json:"follower_count"` // 粉丝总数
}
