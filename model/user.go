package model

// User 用户信息
type User struct {
	Common
	Name          string `json:"name"`           // 用户名称
	Password      string `json:"password"`       // 用户密码
	FollowCount   uint64 `json:"follow_count"`   // 关注总数
	FollowerCount uint64 `json:"follower_count"` // 粉丝总数
}

type UserAPI struct {
	ID            uint64 `json:"id"`
	Name          string `json:"name"`           // 用户名称
	FollowCount   uint64 `json:"follow_count"`   // 关注总数
	FollowerCount uint64 `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // 是否关注
}

// GetUserByUsername 通过用户名获取用户
func GetUserByUsername(username string) (*User, error) {
	user := &User{}
	err := DB.First(user, "Name=?", username).Error
	return user, err
}

// GetUserByID 通过ID获取用户
func GetUserByID(id uint64) (*UserAPI, error) {
	user := &UserAPI{}
	err := DB.First(user, "ID=?", id).Error
	return user, err
}

// NewUser 返回一个新的用户
func NewUser() *User {
	return &User{
		FollowCount:   0,
		FollowerCount: 0,
	}
}

// Save 在数据库中保存该用户
func (user *User) Save() error {
	return DB.Save(user).Error
}
