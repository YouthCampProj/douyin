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
	latest := time.UnixMilli(unixTime)
	err := DB.Limit(30).Where("created_at < ?", latest).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func GetVideoByAuthorID(authorID uint64) ([]*Video, error) {
	var videos []*Video
	err := DB.Where("author_id = ?", authorID).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

// GetVideoListByID 通过视频ID获取视频列表
func GetVideoListByID(ids []uint64) ([]*Video, error) {
	var videos []*Video
	err := DB.Where("id IN (?)", ids).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func NewVideo() *Video {
	return &Video{
		FavoriteCount: 0,
		CommentCount:  0,
	}
}

func (v *Video) Save() error {
	return DB.Save(v).Error
}
