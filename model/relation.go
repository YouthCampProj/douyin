package model

type Relation struct {
	Common
	UserID   int `json:"user_id"`   // 用户ID
	FollowID int `json:"follow_id"` // 被关注用户ID
}
