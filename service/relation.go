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
func (s *RelationActionService) Execute() *serializer.RelationActionResponse {
	switch s.ActionType {
	case "1":
		return s.follow()
	case "2":
		return s.unfollow()
	}
	return serializer.BuildRelationActionResponse(serializer.CodeRelationActionTypeInvalid)
}

// follow 关注
func (s *RelationActionService) follow() *serializer.RelationActionResponse {
	switch model.AddFollow(s.UserID, s.ToUserID) {
	case 0:
		return serializer.BuildRelationActionResponse(serializer.CodeSuccess)
	case 1:
		return serializer.BuildRelationActionResponse(serializer.CodeRelationAlreadyFollow)
	case 2:
		return serializer.BuildRelationActionResponse(serializer.CodeRelationDBError)
	}
	return serializer.BuildRelationActionResponse(serializer.CodeRelationDBError)

}

// unfollow 取消关注
func (s *RelationActionService) unfollow() *serializer.RelationActionResponse {
	switch model.DeleteFollow(s.UserID, s.ToUserID) {
	case 0:
		return serializer.BuildRelationActionResponse(serializer.CodeSuccess)
	case 1:
		return serializer.BuildRelationActionResponse(serializer.CodeRelationNotFollow)
	case 2:
		return serializer.BuildRelationActionResponse(serializer.CodeRelationDBError)
	}
	return serializer.BuildRelationActionResponse(serializer.CodeRelationDBError)
}
