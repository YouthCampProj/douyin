package model

import (
	"errors"
)

// User 用户信息
type User struct {
	Common
	UserID        int    `json:"user_id"`        // 用户id
	Name          string `json:"name"`           // 用户名称
	Password      string `json:"password"`       // 用户密码
	FollowCount   int    `json:"follow_count"`   // 关注总数
	FollowerCount int    `json:"follower_count"` // 粉丝总数
}

// GetUserByUsername 通过用户名与密码获取用户
func GetUserByUsername(username string) (*User, error) {
	user := &User{}
	DB.Find(user, "username=?", username)
	return nil, errors.New("用户名或密码错误")
}
