package controller

import (
	"github.com/YouthCampProj/douyin/pkg/auth"
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"github.com/YouthCampProj/douyin/service"
	"github.com/YouthCampProj/douyin/utils"
	"github.com/gin-gonic/gin"
)

// InitUserRoute 初始化用户相关路由
// /douyin/user
func InitUserRoute(r *gin.RouterGroup) {
	r.POST("register/", UserRegister)
	r.POST("login/", UserLogin)
	r.GET("/", GetUserInfo)
}

// UserRegister 用户注册
// 新用户注册时提供用户名，密码，昵称即可,用户名需要保证唯一。
// 创建成功后返回用户 id 和权限token
// POST /douyin/user/register
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18899952
func UserRegister(c *gin.Context) {
	username := c.Query("username")
	if !utils.UsernameTest(username) {
		c.JSON(200, serializer.BuildUserRegisterResponse(serializer.CodeUserNameInvalid, nil, ""))
		return
	}
	password := c.Query("password")
	if !utils.PasswordTest(password) {
		c.JSON(200, serializer.BuildUserRegisterResponse(serializer.CodeUserPasswordInvalid, nil, ""))
		return
	}

	registerService := service.UserRegisterService{
		Username: username,
		Password: password,
	}
	c.JSON(200, registerService.Register())
	return
}

// UserLogin 用户登录
// 通过用户名和密码进行登录
// 登录成功后返回用户 id 和权限 token
// POST /douyin/user/login
//https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18900033
func UserLogin(c *gin.Context) {
	username := c.Query("username")
	if !utils.UsernameTest(username) {
		c.JSON(200, serializer.BuildUserLoginResponse(serializer.CodeUserNameInvalid, nil, ""))
		return
	}
	password := c.Query("password")
	if !utils.PasswordTest(password) {
		c.JSON(200, serializer.BuildUserLoginResponse(serializer.CodeUserPasswordInvalid, nil, ""))
		return
	}

	loginService := service.UserLoginService{
		Username: username,
		Password: password,
	}
	c.JSON(200, loginService.Login())
	return
}

// GetUserInfo 获取用户信息
// 获取登录用户的 id、昵称
// 如果实现社交部分的功能，还会返回关注数和粉丝数
// GET /douyin/user
func GetUserInfo(c *gin.Context) {
	userID := c.Query("user_id")
	if !utils.UserIDTest(userID) {
		c.JSON(200, serializer.BuildUserInfoResponse(serializer.CodeUserIDInvalid, nil))
		return
	}
	token := c.Query("token")
	user, err := auth.ParseToken(token)
	if err != nil {
		c.JSON(200, serializer.BuildUserInfoResponse(serializer.CodeUserTokenInvalid, nil))
		return
	}

	userInfoService := service.UserInfoService{
		UserID:       user.ID,
		WantedUserID: utils.Str2uint64(userID),
	}
	c.JSON(200, userInfoService.GetUserInfo())

}
