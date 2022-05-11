package model

import (
	"errors"
	"time"
)

func GetFeedListByTime(unixTime time.Time) (time.Time, []*VideoAuthorBundle, error) {
	var videoAuthorBundles []*VideoAuthorBundle
	err := DB.Raw("SELECT\n    v.ID AS id,\n    u.id AS author_id,\n    u.name AS author_name,\n    u.follow_count AS author_follow_count,\n    u.follower_count AS author_follower_count,\n    false AS author_is_follow,\n    v.play_url AS play_url,\n    v.cover_url AS cover_url,\n    v.favorite_count AS favorite_count,\n    v.comment_count AS comment_count,\n    false AS is_favorite\nFROM videos v\nLEFT JOIN users u ON v.author_id=u.id\nWHERE v.created_at < ?\nLIMIT 30;", unixTime).Scan(&videoAuthorBundles).Error
	if err != nil {
		return unixTime, nil, err
	}
	if len(videoAuthorBundles) == 0 {
		return unixTime, nil, errors.New("no video found")
	}
	lastVideoID := videoAuthorBundles[len(videoAuthorBundles)-1].ID
	type CreatedAt struct {
		CreatedAt time.Time `json:"created_at"`
	}
	lastVideoCreatedAt := &CreatedAt{}
	err = DB.Raw("SELECT created_at FROM videos WHERE id = ?", lastVideoID).Scan(lastVideoCreatedAt).Error
	if err != nil {
		return unixTime, nil, err
	}
	return lastVideoCreatedAt.CreatedAt, videoAuthorBundles, nil
}
