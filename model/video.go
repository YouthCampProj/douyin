package model

// Video 视频信息
type Video struct {
	Common
	AuthorID      uint64 `json:"author_id"` // 视频作者
	PlayURL       string `json:"play_url"`  // 视频播放地址
	CoverURL      string `json:"cover_url"` // 视频封面地址
	FavoriteCount uint64 `json:"favorite_count"`
	CommentCount  uint64 `json:"comment_count"`
}
