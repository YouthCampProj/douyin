package model

type Favorite struct {
	Common
	UserID  uint64 `json:"user_id"`  // 点赞的用户ID
	VideoID uint64 `json:"video_id"` // 被点赞的视频ID
}
