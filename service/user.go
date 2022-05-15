package service

import (
	"github.com/YouthCampProj/douyin/model"
	"github.com/YouthCampProj/douyin/pkg/auth"
	"github.com/YouthCampProj/douyin/pkg/config"
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"github.com/YouthCampProj/douyin/utils"
)

type UserLoginService struct {
	Username string
	Password string
}

type UserRegisterService struct {
	Username string
	Password string
}

type UserInfoService struct {
	UserID       uint64 // 发起请求的用户ID
	WantedUserID uint64 // 希望获取的用户ID
}

// Login 通过用户名与密码获取用户
func (u *UserLoginService) Login() *serializer.UserLoginResponse {
	user, err := model.GetUserByUsername(u.Username)
	// 下述判断失败均返回相同的错误 防止恶意试探用户名
	if err != nil {
		return serializer.BuildUserLoginResponse(serializer.CodeUserLoginFailed, nil, "")
	}
	// 对密码进行匹配
	if !utils.StrMatch(user.Password, u.Password, config.Conf.Salt) {
		return serializer.BuildUserLoginResponse(serializer.CodeUserLoginFailed, nil, "")
	}

	// 生成token
	token := auth.NewToken(user)
	return serializer.BuildUserLoginResponse(serializer.CodeSuccess, user, token)
}

// Register 用户注册
func (u *UserRegisterService) Register() *serializer.UserRegisterResponse {
	// 判断用户名是否存在
	user, err := model.GetUserByUsername(u.Username)
	if err == nil {
		return serializer.BuildUserRegisterResponse(serializer.CodeUserAlreadyExists, nil, "")
	}
	// 创建用户
	user = model.NewUser()
	user.Name = u.Username
	user.Password = utils.StrEncrypt(u.Password, config.Conf.Salt)
	err = model.DB.Create(&user).Error
	if err != nil {
		return serializer.BuildUserRegisterResponse(serializer.CodeUserRegisterFailed, nil, "")
	}
	// 生成token
	token := auth.NewToken(user)
	return serializer.BuildUserRegisterResponse(serializer.CodeSuccess, user, token)

}

// GetUserInfo 获取用户信息
func (u *UserInfoService) GetUserInfo() *serializer.UserInfoResponse {
	user, err := model.GetUserByID(u.WantedUserID)
	if err != nil {
		return serializer.BuildUserInfoResponse(serializer.CodeUserNotFound, nil)
	}
	// 获取关注状态
	user.IsFollow = model.IsFollow(u.UserID, u.WantedUserID)
	return serializer.BuildUserInfoResponse(serializer.CodeSuccess, user)
}
