package model

type Favorite struct {
	Common
	UserID  uint64 `json:"user_id"`  // 点赞的用户ID
	VideoID uint64 `json:"video_id"` // 被点赞的视频ID
}

type FavoriteVideo struct {
	ID     uint64 `json:"id"`
	Author struct {
		ID            uint64 `json:"id"`
		Name          string `json:"name"`           // 用户名称
		FollowCount   uint64 `json:"follow_count"`   // 关注总数
		FollowerCount uint64 `json:"follower_count"` // 粉丝总数
		IsFollow      bool   `json:"is_follow"`      // 是否关注
	} `json:"author"`
	PlayURL       string `json:"play_url"`  // 视频播放地址
	CoverURL      string `json:"cover_url"` // 视频封面地址
	FavoriteCount uint64 `json:"favorite_count"`
	CommentCount  uint64 `json:"comment_count"`
}

// AddFavorite 添加点赞信息 返回值 0-点赞成功 1-已经点赞 2-数据库错误
func AddFavorite(userID uint64, videoID uint64) int {
	if IsFavorite(userID, videoID) {
		return 1
	}
	favorite := &Favorite{
		UserID:  userID,
		VideoID: videoID,
	}
	if DB.Create(favorite).Error != nil {
		return 2
	}
	return 0
}

// DeleteFavorite 删除点赞信息 返回值 0-删除成功 1-没有点赞 2-数据库错误
func DeleteFavorite(userID uint64, videoID uint64) int {
	if !IsFavorite(userID, videoID) {
		return 1
	}
	if DB.Where("user_id = ? AND video_id = ?", userID, videoID).Delete(&Favorite{}).Error != nil {
		return 2
	}
	return 0
}

// IsFavorite 判断是否已经点赞
func IsFavorite(userID uint64, videoID uint64) bool {
	return DB.First(&Favorite{}, "user_id = ? and video_id = ?", userID, videoID).Error == nil
}

//// GetFavoriteVideoList 获取用户的点赞视频
//func GetFavoriteVideoList(userID uint64) ([]*FavoriteVideo, error) {
//	var favorites []*FavoriteVideo
//	DB.
//}
