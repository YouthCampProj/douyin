package model

import (
	"gorm.io/gorm"
	"time"
)

/*
    评论存储模型
	目前的想法是通过视频ID+用户ID索引出该用户在当前视频下的所有评论ID
	评论ID仅对当前视频+当前用户唯一以提高查询效率并避免爆int范围

	TODO: 客户端放出后， 需要检查客户端对删除评论按钮的具体实现 防止意外的评论删除操作
*/

type Comment struct {
	Common
	UserID      uint64 `json:"user_id"`      // 用户ID
	VideoID     uint64 `json:"video_id"`     // 视频ID
	CommentText string `json:"comment_text"` // 评论内容
}

type CommentUserBundle struct {
	ID         uint64    `json:"id"` // 评论ID
	User       *UserAPI  `json:"user" gorm:"embedded;embeddedPrefix:user_"`
	Content    string    `json:"content"`     // 评论内容
	CreateDate time.Time `json:"create_date"` // 评论时间
}

// AddComment 在数据库中添加评论
func AddComment(userID uint64, videoID uint64, text string) error {
	c := &Comment{
		UserID:      userID,
		VideoID:     videoID,
		CommentText: text,
	}
	err := c.Save()
	if err != nil {
		return err
	}
	// 更新video表评论计数
	return DB.Model(&Video{}).Where("id = ?", videoID).Update("comment_count", gorm.Expr("comment_count + 1")).Error
}

// DeleteComment 在数据库中删除评论
func DeleteComment(userID uint64, videoID uint64, CommentID uint64) error {
	err := DB.Delete(&Comment{}, "user_id = ? and video_id = ? and id = ?", userID, videoID, CommentID).Error
	if err != nil {
		return err
	}
	// 更新video表评论计数
	return DB.Model(&Video{}).Where("id = ?", videoID).Update("comment_count", gorm.Expr("comment_count - 1")).Error
}

func GetCommentUserBundle(userID uint64, videoID uint64) ([]*CommentUserBundle, error) {
	var comments []*CommentUserBundle
	err := DB.Raw("SELECT\n    c.id AS id,\n    u.id AS user_id,\n    u.name AS user_name,\n    u.follow_count AS follow_count,\n    u.follower_count AS follower_count,\n    IF(r.id IS NULL,false,true) AS is_follow,\n    c.comment_text AS content,\n    c.created_at AS create_date\nFROM comments c\nLEFT JOIN users u ON c.user_id=u.id\nLEFT JOIN relations r on u.id = r.follow_id AND r.user_id = ?\nWHERE c.video_id=?\nORDER BY c.updated_at;", userID, videoID).Scan(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *Comment) Save() error {
	return DB.Save(c).Error
}
