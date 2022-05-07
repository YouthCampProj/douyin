package model

import "time"

// Video 视频信息
type Video struct {
	Common
	AuthorID      uint64 `json:"author_id"` // 视频作者
	PlayURL       string `json:"play_url"`  // 视频播放地址
	CoverURL      string `json:"cover_url"` // 视频封面地址
	FavoriteCount uint64 `json:"favorite_count"`
	CommentCount  uint64 `json:"comment_count"`
}

func GetVideoByTime(unixTime int64) ([]*Video, error) {
	var videos []*Video
	latest := time.Unix(unixTime, 0)
	err := DB.Limit(30).Where("created_at < ?", latest).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
