package service

import (
	"github.com/YouthCampProj/douyin/model"
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"log"
)

type CommentActionService struct {
	UserID      uint64
	VideoID     uint64
	ActionType  string
	CommentID   uint64
	CommentText string
}

func (s *CommentActionService) Execute() *serializer.CommentActionResponse {
	switch s.ActionType {
	case "1":
		return s.Publish()
	case "2":
		return s.Delete()
	}
	return serializer.BuildCommentActionResponse(serializer.CodeCommentActionInvalid)
}

func (s *CommentActionService) Publish() *serializer.CommentActionResponse {
	// TODO: 对传入的评论内容进行安全校验

	if err := model.AddComment(s.VideoID, s.UserID, s.CommentText); err != nil {
		log.Println(err)
		return serializer.BuildCommentActionResponse(serializer.CodeCommentActionDBFailed)
	}
	return serializer.BuildCommentActionResponse(serializer.CodeSuccess)
}

func (s *CommentActionService) Delete() *serializer.CommentActionResponse {
	if err := model.DeleteComment(s.UserID, s.VideoID, s.CommentID); err != nil {
		log.Println(err)
		return serializer.BuildCommentActionResponse(serializer.CodeCommentActionDBFailed)
	}
	return serializer.BuildCommentActionResponse(serializer.CodeSuccess)
}
