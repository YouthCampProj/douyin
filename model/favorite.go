package model

type Favorite struct {
	Common
	UserID  int `json:"user_id"`  // 点赞的用户ID
	VideoID int `json:"video_id"` // 被点赞的视频ID
}
