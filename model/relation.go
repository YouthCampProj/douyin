package model

type Relation struct {
	Common
	UserID   uint64 `json:"user_id"`   // 用户ID
	FollowID uint64 `json:"follow_id"` // 被关注用户ID
}

func IsFollow(userID, followID uint64) bool {
	relation := Relation{}
	return DB.First(relation, "user_id = ? and follow_id = ?", userID, followID).Error == nil
}
