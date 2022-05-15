package service

import (
	"github.com/YouthCampProj/douyin/model"
	"github.com/YouthCampProj/douyin/pkg/serializer"
)

type RelationActionService struct {
	UserID     uint64
	ToUserID   uint64
	ActionType string
}

// Execute 执行关注操作
func (s *RelationActionService) Execute() *serializer.FavoriteActionResponse {
	switch s.ActionType {
	case "1":
		return s.follow()
	case "2":
		return s.unfollow()
	}
	return serializer.BuildFavoriteActionResponse(serializer.CodeRelationActionTypeInvalid)
}

// follow 关注
func (s *RelationActionService) follow() *serializer.FavoriteActionResponse {
	switch model.AddFollow(s.UserID, s.ToUserID) {
	case 0:
		return serializer.BuildFavoriteActionResponse(serializer.CodeSuccess)
	case 1:
		return serializer.BuildFavoriteActionResponse(serializer.CodeRelationAlreadyFollow)
	case 2:
		return serializer.BuildFavoriteActionResponse(serializer.CodeRelationDBError)
	}
	return serializer.BuildFavoriteActionResponse(serializer.CodeRelationDBError)

}

// unfollow 取消关注
func (s *RelationActionService) unfollow() *serializer.FavoriteActionResponse {
	switch model.DeleteFollow(s.UserID, s.ToUserID) {
	case 0:
		return serializer.BuildFavoriteActionResponse(serializer.CodeSuccess)
	case 1:
		return serializer.BuildFavoriteActionResponse(serializer.CodeRelationNotFollow)
	case 2:
		return serializer.BuildFavoriteActionResponse(serializer.CodeRelationDBError)
	}
	return serializer.BuildFavoriteActionResponse(serializer.CodeRelationDBError)
}
