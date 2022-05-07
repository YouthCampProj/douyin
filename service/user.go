package service

import (
	"github.com/YouthCampProj/douyin/model"
	"github.com/YouthCampProj/douyin/pkg/config"
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"github.com/YouthCampProj/douyin/utils"
)

type UserLoginService struct {
	username string
	password string
}

// GetUserByBasicAuth 通过用户名与密码获取用户
func (u *UserLoginService) Login(username string, password string) *serializer.UserLoginResponse {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		return serializer.BuildUserLoginResponse(serializer.CodeUserNotFound, nil, "")
	}
	// 对密码进行匹配
	if !utils.StrMatch(user.Password, password, config.Conf.Salt) {
		return serializer.BuildUserLoginResponse(serializer.CodeUserLoginFailed, nil, "")
	}
	// 生成token

	return serializer.BuildUserLoginResponse(serializer.CodeSuccess, user, "")
}
