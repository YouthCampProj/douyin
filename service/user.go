package service

import (
	"errors"
	"github.com/YouthCampProj/douyin/model"
	"github.com/YouthCampProj/douyin/utils"
)

// GetUserByBasicAuth 通过用户名与密码获取用户
func GetUserByBasicAuth(username string, password string) (*model.User, error) {
	user := &model.User{}
	// 完整加载用户到内存 避免二次查询
	DB.Find(user, "username=?", username)
	// 对密码进行匹配
	if utils.StrMatch(user.Password, password, Conf.Salt) {
		return user, nil
	}
	return nil, errors.New("用户名或密码错误")
}
