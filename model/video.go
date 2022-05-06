package model

// Video 视频信息
type Video struct {
	Common
	VideoID       int    `json:"video_id"`  // 视频唯一标识
	AuthorID      int    `json:"author_id"` // 视频作者
	PlayURL       string `json:"play_url"`  // 视频播放地址
	CoverURL      string `json:"cover_url"` // 视频封面地址
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
}
