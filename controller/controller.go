package controller

import "github.com/gin-gonic/gin"

// Init 初始化controller
func Init(r *gin.Engine) {
	InitFileRoute(r.Group("uploads"))
	InitDouyinRoute(r.Group("douyin"))
}

func InitDouyinRoute(r *gin.RouterGroup) {
	InitCommentRoute(r.Group("comment"))
	InitFavoriteRoute(r.Group("favorite"))
	InitFeedRoute(r.Group("feed"))
	InitPublishRoute(r.Group("publish"))
	InitRelationRoutes(r.Group("relation"))
	InitUserRoute(r.Group("user"))

}
