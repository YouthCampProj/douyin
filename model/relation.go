package model

import (
	"gorm.io/gorm"
	"log"
)

type Relation struct {
	Common
	UserID   uint64 `json:"user_id"`   // 用户ID
	FollowID uint64 `json:"follow_id"` // 被关注用户ID
}

func IsFollow(userID, followID uint64) bool {
	relation := &Relation{}
	return DB.First(relation, "user_id = ? and follow_id = ?", userID, followID).Error == nil
}

// AddFollow 添加关注 返回值 0-成功 1-已关注 2-数据库错误
func AddFollow(userID, followID uint64) int {
	if IsFollow(userID, followID) {
		return 1
	}
	relation := Relation{
		UserID:   userID,
		FollowID: followID,
	}
	// 添加关注
	if err := DB.Create(&relation).Error; err != nil {
		return 2
	}
	// 增加粉丝数
	if err := DB.Model(&User{}).Where("id = ?", followID).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
		return 2
	}
	return 0
}

// DeleteFollow 删除关注 返回值 0-成功 1-未关注 2-数据库错误
func DeleteFollow(userID, followID uint64) int {
	if !IsFollow(userID, followID) {
		return 1
	}
	// 删除关注
	if err := DB.Where("user_id=? AND follow_id=?", userID, followID).Delete(&Relation{}).Error; err != nil {
		log.Println(err)
		return 2
	}
	// 减少粉丝数
	if err := DB.Model(&User{}).Where("id = ?", followID).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
		log.Println(err)
		return 2
	}
	return 0
}

// GetFollowUserList 获取userID的关注列表
// requestFromID是发起请求的用户ID(用于构建关注关系)
func GetFollowUserList(userID, requestFromID uint64) ([]*UserAPI, error) {
	var userList []*UserAPI
	if err := DB.Raw("SELECT\n    u.id AS id,\n    u.name AS name,\n    u.follow_count AS follow_count,\n    u.follower_count AS follower_count,\n    IF(r.id IS NULL,false,true) AS is_follow\nFROM users u\nLEFT JOIN relations r ON u.id = r.follow_id AND r.user_id=?\nWHERE u.id IN (\n    SELECT r2.follow_id\n    FROM relations r2\n    WHERE r2.user_id=?\n    );", requestFromID, userID).Scan(&userList).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return userList, nil
}

// GetFollowerUserList 获取userID的粉丝列表
// requestFromID是发起请求的用户ID(用于构建关注关系)
func GetFollowerUserList(userID, requestFromID uint64) ([]*UserAPI, error) {
	var userList []*UserAPI
	if err := DB.Raw("SELECT\n    u.id AS id,\n    u.name AS name,\n    u.follow_count AS follow_count,\n    u.follower_count AS follower_count,\n    IF(r.id IS NULL,false,true) AS is_follow\nFROM users u\nLEFT JOIN relations r ON u.id = r.follow_id AND r.user_id=?\nWHERE u.id IN (\n    SELECT r2.user_id\n    FROM relations r2\n    WHERE r2.follow_id=?\n    );", requestFromID, userID).Scan(&userList).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return userList, nil
}
