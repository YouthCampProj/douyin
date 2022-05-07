package controller

import (
	"github.com/gin-gonic/gin"
)

// InitFeedRoute 初始化视频流相关路由
// /douyin/feed
func InitFeedRoute(r *gin.RouterGroup) {
	r.GET("/", GetFeed)
}

// GetFeed 视频流接口
// 无需登录，返回按投稿时间倒序的视频列表
// 视频数由服务端控制，单次最多30个
// GET /douyin/feed/
// https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145
func GetFeed(c *gin.Context) {
	//可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	//latestTime := c.DefaultQuery("latest_time", time.Now().String())
	//TODO: 视频流接口
}
