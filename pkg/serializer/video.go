package serializer

import "github.com/YouthCampProj/douyin/model"

// Video 视频信息
type Video struct {
	ID            uint64 `json:"id"` // 视频唯一标识
	Author        *User  `json:"author"`
	PlayURL       string `json:"play_url"`       // 视频播放地址
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount uint64 `json:"favorite_count"` // 视频的点赞总数
	CommentCount  uint64 `json:"comment_count"`  // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
}

func BuildVideoResponse(video *model.Video, author *User, isFavorite bool) *Video {
	return &Video{
		ID:            video.ID,
		Author:        author,
		PlayURL:       video.PlayURL,
		CoverURL:      video.CoverURL,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    isFavorite,
	}
}
