package model

func GetPublishListByAuthorID(authorID uint64) ([]*VideoAuthorBundle, error) {
	userID := authorID // 这里的UserID是发起请求的用户ID
	var videoAuthorBundles []*VideoAuthorBundle
	// 下述查询预留了发起请求的用户ID与视频作者ID不一致的情况
	// 如果发起请求的用户ID与视频作者ID不一致，需要传入userID
	err := DB.Raw("SELECT\n    v.ID AS id,\n    u.id AS author_id,\n    u.name AS author_name,\n    u.follow_count AS author_follow_count,\n    u.follower_count AS author_follower_count,\n    IF(r.id IS NULL,false,true) AS author_is_follow,\n    v.play_url AS play_url,\n    v.cover_url AS cover_url,\n    v.favorite_count AS favorite_count,\n    v.comment_count AS comment_count,\n    IF(f.id IS NULL,false,true) AS is_favorite,\n    v.title AS title\nFROM videos v\nLEFT JOIN users u ON v.author_id=u.id\nLEFT JOIN relations r on u.id = r.follow_id AND r.user_id = ?\nLEFT JOIN favorites f on u.id = f.user_id AND v.id = f.video_id\nWHERE author_id = ?\nORDER BY v.created_at", userID, authorID).Scan(&videoAuthorBundles).Error
	if err != nil {
		return nil, err
	}
	return videoAuthorBundles, nil
}
