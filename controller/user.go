package controller

import "github.com/gin-gonic/gin"

// InitUserRoute 初始化用户相关路由
// /douyin/user
func InitUserRoute(r *gin.RouterGroup) {
	r.POST("/register", UserRegister)
	r.POST("/login", UserLogin)
	r.GET("/", GetUser)
}

// UserRegister 用户注册
// 新用户注册时提供用户名，密码，昵称即可,用户名需要保证唯一。
// 创建成功后返回用户 id 和权限token
// POST /douyin/user/register
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18899952
func UserRegister(c *gin.Context) {
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	// TODO 用户注册接口
}

// UserLogin 用户登录
// 通过用户名和密码进行登录
// 登录成功后返回用户 id 和权限 token
// POST /douyin/user/login
//https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18900033
func UserLogin(c *gin.Context) {
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	// TODO 用户登录接口
}

// GetUser 获取用户信息
// 获取登录用户的 id、昵称
// 如果实现社交部分的功能，还会返回关注数和粉丝数
// GET /douyin/user
func GetUser(c *gin.Context) {
	//c.Query("user_id")
	//c.Query("token")
	// TODO 获取用户信息接口
}
