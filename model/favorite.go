package model

import "gorm.io/gorm"

type Favorite struct {
	Common
	UserID  uint64 `json:"user_id"`  // 点赞的用户ID
	VideoID uint64 `json:"video_id"` // 被点赞的视频ID
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
	// 在favorite表中添加该记录
	if DB.Create(favorite).Error != nil {
		return 2
	}
	// 更新video表中的点赞数
	if DB.Model(&Video{}).Where("id = ?", videoID).Update("favorite_count", gorm.Expr("favorite_count + 1")).Error != nil {
		return 2
	}
	return 0
}

// DeleteFavorite 删除点赞信息 返回值 0-删除成功 1-没有点赞 2-数据库错误
func DeleteFavorite(userID uint64, videoID uint64) int {
	if !IsFavorite(userID, videoID) {
		return 1
	}
	// 在favorite表中删除该记录
	if DB.Where("user_id = ? AND video_id = ?", userID, videoID).Delete(&Favorite{}).Error != nil {
		return 2
	}
	// 更新video表中的点赞数
	if DB.Model(&Video{}).Where("id = ?", videoID).Update("favorite_count", gorm.Expr("favorite_count - 1")).Error != nil {
		return 2
	}
	return 0
}

// IsFavorite 判断是否已经点赞
func IsFavorite(userID uint64, videoID uint64) bool {
	return DB.First(&Favorite{}, "user_id = ? and video_id = ?", userID, videoID).Error == nil
}

// GetFavoriteVideoList 获取用户的点赞视频
func GetFavoriteVideoList(userID uint64, reqID ...uint64) ([]*VideoAuthorBundle, error) {
	var favorites []*VideoAuthorBundle
	var err error
	if len(reqID) == 0 {
		err = DB.Raw("SELECT\n    v.ID AS id,\n    u.id AS author_id,\n    u.name AS author_name,\n    u.follow_count AS author_follow_count,\n    u.follower_count AS author_follower_count,\n    FALSE AS author_is_follow,\n    v.play_url AS play_url,\n    v.cover_url AS cover_url,\n    v.favorite_count AS favorite_count,\n    v.comment_count AS comment_count,\n    true AS is_favorite,\n    v.title AS title\nFROM videos v\nLEFT JOIN users u ON v.author_id=u.id\nWHERE v.id IN(\n    SELECT video_id\n    FROM favorites\n    WHERE user_id = ?\n    )\nORDER BY v.created_at;", userID).Scan(&favorites).Error
	} else {
		err = DB.Raw("SELECT\n    v.ID AS id,\n    u.id AS author_id,\n    u.name AS author_name,\n    u.follow_count AS author_follow_count,\n    u.follower_count AS author_follower_count,\n    IF(r.id IS NULL,false,true) AS author_is_follow,\n    v.play_url AS play_url,\n    v.cover_url AS cover_url,\n    v.favorite_count AS favorite_count,\n    v.comment_count AS comment_count,\n    true AS is_favorite,\n    v.title AS title\nFROM videos v\nLEFT JOIN users u ON v.author_id=u.id\nLEFT JOIN relations r on u.id = r.follow_id AND r.user_id = ?\nWHERE v.id IN(\n    SELECT video_id\n    FROM favorites\n    WHERE user_id = ?\n    )\nORDER BY v.created_at;", reqID[0], userID).Scan(&favorites).Error

	}
	if err != nil {
		return nil, err
	}
	return favorites, nil
}
